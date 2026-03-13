package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// ProjectFields marshals v to JSON, then filters to only the specified dot-path fields.
// fields is a comma-separated list like "title,document.url,metadata.datasource".
// If fields is empty, the full value is written.
func ProjectFields(w io.Writer, v any, fields string) error {
	if fields == "" {
		return WriteJSON(w, v)
	}

	// Marshal to generic map first
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}

	var raw any
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	fieldList := strings.Split(fields, ",")
	for i := range fieldList {
		fieldList[i] = strings.TrimSpace(fieldList[i])
	}

	result := extractFields(raw, fieldList)

	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("json marshal result: %w", err)
	}
	_, err = fmt.Fprintln(w, string(out))
	return err
}

// extractFields builds a filtered version of value containing only the requested fields.
func extractFields(value any, fields []string) any {
	switch v := value.(type) {
	case map[string]any:
		result := map[string]any{}
		for _, field := range fields {
			parts := strings.SplitN(field, ".", 2)
			key := parts[0]
			if val, ok := v[key]; ok {
				if len(parts) == 1 {
					result[key] = val
				} else {
					// Recurse into nested path
					sub := extractFields(val, []string{parts[1]})
					if existing, ok := result[key]; ok {
						// Merge with existing sub-object
						if em, ok := existing.(map[string]any); ok {
							if sm, ok := sub.(map[string]any); ok {
								for sk, sv := range sm {
									em[sk] = sv
								}
								result[key] = em
								continue
							}
						}
					}
					result[key] = sub
				}
			}
		}
		return result

	case []any:
		// Apply projection to each element
		out := make([]any, len(v))
		for i, elem := range v {
			out[i] = extractFields(elem, fields)
		}
		return out

	default:
		return value
	}
}

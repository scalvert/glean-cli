package output

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CleanseSearchResponse strips UI-specific fields from a search response,
// keeping only the fields relevant to programmatic consumers.
//
// This is a stopgap until POST /api/search ships (see RFC: Search Data
// Retrieval API). Delete this file and its call sites once the new API
// is available.
func CleanseSearchResponse(resp any) (any, error) {
	data, err := json.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf("cleanse marshal: %w", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("cleanse unmarshal: %w", err)
	}

	result := filterMap(raw, responseAllowlist)

	if results, ok := result["results"].([]any); ok {
		result["results"] = filterEmptyResults(results)
		for _, r := range result["results"].([]any) {
			if m, ok := r.(map[string]any); ok {
				if snippets, ok := m["snippets"].([]any); ok {
					m["snippets"] = filterEmptySnippets(snippets)
					if len(m["snippets"].([]any)) == 0 {
						delete(m, "snippets")
					}
				}
			}
		}
	}

	return result, nil
}

type allowlist map[string]allowlist

var responseAllowlist = allowlist{
	"results":        resultAllowlist,
	"cursor":         nil,
	"hasMoreResults": nil,
	"requestID":      nil,
}

var resultAllowlist = allowlist{
	"title":    nil,
	"url":      nil,
	"snippets": snippetAllowlist,
	"document": documentAllowlist,
}

var documentAllowlist = allowlist{
	"title":      nil,
	"url":        nil,
	"datasource": nil,
	"docType":    nil,
	"metadata":   metadataAllowlist,
}

var metadataAllowlist = allowlist{
	"datasource":   nil,
	"objectType":   nil,
	"author":       personAllowlist,
	"owner":        personAllowlist,
	"assignedTo":   personAllowlist,
	"updatedBy":    personAllowlist,
	"updateTime":   nil,
	"createTime":   nil,
	"status":       nil,
	"priority":     nil,
	"container":    nil,
	"datasourceId": nil,
}

var personAllowlist = allowlist{
	"name": nil,
}

var snippetAllowlist = allowlist{
	"text":     nil,
	"mimeType": nil,
}

// WarnStrippedFields checks whether any of the requested --fields paths
// were removed by cleansing. Returns a list of field paths that don't
// exist in the allowlist.
//
// Stopgap — delete with cleanse.go when POST /api/search ships.
func WarnStrippedFields(fields string) []string {
	if fields == "" {
		return nil
	}
	var stripped []string
	for _, f := range strings.Split(fields, ",") {
		f = strings.TrimSpace(f)
		if f == "" {
			continue
		}
		if !isAllowedPath(f, responseAllowlist) {
			stripped = append(stripped, f)
		}
	}
	return stripped
}

// isAllowedPath checks whether a dot-separated field path exists in the allowlist tree.
// A path is allowed if every segment resolves in the allowlist. A nil allowlist value
// at any point means "keep everything below here", so all deeper paths are allowed.
func isAllowedPath(path string, al allowlist) bool {
	parts := strings.SplitN(path, ".", 2)
	key := parts[0]

	childAL, ok := al[key]
	if !ok {
		return false
	}
	if len(parts) == 1 {
		return true
	}
	// nil allowlist = keep everything below → any sub-path is valid
	if childAL == nil {
		return true
	}
	return isAllowedPath(parts[1], childAL)
}

// filterMap recursively keeps only keys present in the allowlist.
// If the allowlist value for a key is nil, the entire value is kept as-is.
// If the allowlist value is a nested allowlist, the value is filtered recursively.
func filterMap(m map[string]any, al allowlist) map[string]any {
	out := make(map[string]any, len(al))
	for key, childAL := range al {
		val, ok := m[key]
		if !ok {
			continue
		}
		if childAL == nil {
			out[key] = val
			continue
		}
		switch v := val.(type) {
		case map[string]any:
			out[key] = filterMap(v, childAL)
		case []any:
			out[key] = filterSlice(v, childAL)
		default:
			out[key] = val
		}
	}
	return out
}

// filterEmptyResults removes results that have no meaningful content after cleansing.
// Structured results from the SDK (e.g. knowledge cards) cleanse down to just {"url": ""}
// since they have no document, title, or snippets.
func filterEmptyResults(results []any) []any {
	out := make([]any, 0, len(results))
	for _, r := range results {
		m, ok := r.(map[string]any)
		if !ok {
			out = append(out, r)
			continue
		}
		if _, hasDoc := m["document"]; hasDoc {
			out = append(out, r)
			continue
		}
		if title, _ := m["title"].(*string); title != nil {
			out = append(out, r)
			continue
		}
		if title, ok := m["title"].(string); ok && title != "" {
			out = append(out, r)
			continue
		}
		// No document and no title — skip this empty result
	}
	return out
}

// filterEmptySnippets removes snippets where the text field is empty or missing.
func filterEmptySnippets(snippets []any) []any {
	out := make([]any, 0, len(snippets))
	for _, s := range snippets {
		m, ok := s.(map[string]any)
		if !ok {
			continue
		}
		if text, ok := m["text"].(string); ok && text != "" {
			out = append(out, s)
		}
	}
	return out
}

// filterSlice applies the allowlist to each element in a slice.
func filterSlice(s []any, al allowlist) []any {
	out := make([]any, len(s))
	for i, elem := range s {
		if m, ok := elem.(map[string]any); ok {
			out[i] = filterMap(m, al)
		} else {
			out[i] = elem
		}
	}
	return out
}

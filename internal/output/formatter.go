// Package output provides structured output formatting for glean-cli commands.
package output

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

// Format constants for --output flag values.
const (
	OutputJSON   = "json"
	OutputNDJSON = "ndjson"
	OutputText   = "text"
)

// WriteJSON marshals v as indented JSON to w.
func WriteJSON(w io.Writer, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}
	_, err = fmt.Fprintln(w, string(data))
	return err
}

// WriteNDJSON marshals each element of a slice as a separate JSON line to w.
// If v is not a slice, it writes the whole value as a single line.
func WriteNDJSON(w io.Writer, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Slice {
		// Not a slice — write as single-line JSON
		data, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("json marshal: %w", err)
		}
		_, err = fmt.Fprintln(w, string(data))
		return err
	}

	for i := range rv.Len() {
		data, err := json.Marshal(rv.Index(i).Interface())
		if err != nil {
			return fmt.Errorf("json marshal element %d: %w", i, err)
		}
		_, err = fmt.Fprintln(w, string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteFormatted writes v in the format specified by outputFormat.
// textFn is called for OutputText; it should write human-readable output to w.
func WriteFormatted(w io.Writer, v any, outputFormat string, textFn func(io.Writer, any) error) error {
	switch outputFormat {
	case OutputNDJSON:
		return WriteNDJSON(w, v)
	case OutputText:
		if textFn != nil {
			return textFn(w, v)
		}
		return WriteJSON(w, v) // fallback
	default: // json
		return WriteJSON(w, v)
	}
}

package debug

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// resetState clears global state between tests.
func resetState(t *testing.T) {
	t.Helper()
	patterns = nil
	globalEnabled.Store(false)
	lastTime.Store(0)
	maxNameLen.Store(0)
}

func TestMatchesPatterns_Wildcard(t *testing.T) {
	resetState(t)
	patterns = []pattern{{glob: true, prefix: ""}} // "*"

	if !matchesPatterns("auth:token") {
		t.Error("wildcard should match auth:token")
	}
	if !matchesPatterns("http:request") {
		t.Error("wildcard should match http:request")
	}
}

func TestMatchesPatterns_PrefixGlob(t *testing.T) {
	resetState(t)
	patterns = []pattern{{glob: true, prefix: "auth:"}} // "auth:*"

	if !matchesPatterns("auth:token") {
		t.Error("auth:* should match auth:token")
	}
	if !matchesPatterns("auth:discovery") {
		t.Error("auth:* should match auth:discovery")
	}
	if matchesPatterns("http:request") {
		t.Error("auth:* should not match http:request")
	}
}

func TestMatchesPatterns_Exact(t *testing.T) {
	resetState(t)
	patterns = []pattern{{prefix: "auth:token"}} // "auth:token"

	if !matchesPatterns("auth:token") {
		t.Error("exact pattern should match auth:token")
	}
	if matchesPatterns("auth:discovery") {
		t.Error("exact pattern should not match auth:discovery")
	}
}

func TestMatchesPatterns_Negation(t *testing.T) {
	resetState(t)
	// "*,-http:*" — everything except http namespaces
	patterns = []pattern{
		{glob: true, prefix: ""},
		{glob: true, prefix: "http:", negate: true},
	}

	if !matchesPatterns("auth:token") {
		t.Error("should match auth:token")
	}
	if matchesPatterns("http:request") {
		t.Error("should not match http:request (negated)")
	}
}

func TestMatchesPatterns_CommaSeparated(t *testing.T) {
	resetState(t)
	// "auth:token,auth:discovery"
	patterns = []pattern{
		{prefix: "auth:token"},
		{prefix: "auth:discovery"},
	}

	if !matchesPatterns("auth:token") {
		t.Error("should match auth:token")
	}
	if !matchesPatterns("auth:discovery") {
		t.Error("should match auth:discovery")
	}
	if matchesPatterns("auth:dcr") {
		t.Error("should not match auth:dcr")
	}
}

func TestMatchesPatterns_Empty(t *testing.T) {
	resetState(t)
	patterns = nil

	if matchesPatterns("auth:token") {
		t.Error("empty patterns should match nothing")
	}
}

func TestEnable_RetroactiveActivation(t *testing.T) {
	resetState(t)
	// No patterns set — logger starts disabled
	l := New("auth:token")
	if l.Enabled() {
		t.Error("logger should be disabled before Enable()")
	}

	Enable()
	if !l.Enabled() {
		t.Error("logger should be enabled after Enable()")
	}
}

func TestLog_DisabledProducesNoOutput(t *testing.T) {
	resetState(t)

	// Capture stderr
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	l := New("auth:token")
	l.Log("should not appear")

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stderr = old

	if buf.Len() > 0 {
		t.Errorf("disabled logger should produce no output, got: %s", buf.String())
	}
}

func TestLog_EnabledProducesOutput(t *testing.T) {
	resetState(t)
	patterns = []pattern{{glob: true, prefix: ""}} // "*"

	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	l := New("auth:token")
	l.Log("hello %s", "world")

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stderr = old

	output := buf.String()
	if !strings.Contains(output, "auth:token") {
		t.Errorf("output should contain namespace, got: %s", output)
	}
	if !strings.Contains(output, "hello world") {
		t.Errorf("output should contain message, got: %s", output)
	}
	if !strings.Contains(output, "+") {
		t.Errorf("output should contain time delta, got: %s", output)
	}
}

func TestLog_GlobalEnableProducesOutput(t *testing.T) {
	resetState(t)
	// No patterns, but Enable() is called

	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	l := New("http:request")
	Enable()
	l.Log("GET https://example.com")

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)
	os.Stderr = old

	output := buf.String()
	if !strings.Contains(output, "http:request") {
		t.Errorf("output should contain namespace, got: %s", output)
	}
	if !strings.Contains(output, "GET https://example.com") {
		t.Errorf("output should contain message, got: %s", output)
	}
}

func TestFormatDelta(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"0ms", "+0ms"},
		{"500µs", "+0ms"},
		{"150ms", "+150ms"},
		{"1.5s", "+1.5s"},
		{"90s", "+1m30s"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			d, err := parseDuration(tt.input)
			if err != nil {
				t.Fatalf("bad test input: %v", err)
			}
			got := formatDelta(d)
			if got != tt.expected {
				t.Errorf("formatDelta(%v) = %q, want %q", d, got, tt.expected)
			}
		})
	}
}

// parseDuration is a test helper that wraps time.ParseDuration with support
// for the "ms" suffix that time.ParseDuration already handles.
func parseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}

func TestPickColor_Deterministic(t *testing.T) {
	c1 := pickColor("auth:token")
	c2 := pickColor("auth:token")
	if fmt.Sprintf("%p", c1) != fmt.Sprintf("%p", c2) {
		t.Error("pickColor should return the same color for the same namespace")
	}
}

func TestLog_ConcurrentSafety(t *testing.T) {
	resetState(t)
	patterns = []pattern{{glob: true, prefix: ""}}

	old := os.Stderr
	_, w, _ := os.Pipe()
	os.Stderr = w

	l := New("test:concurrent")
	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			l.Log("message %d", n)
		}(i)
	}
	wg.Wait()

	w.Close()
	os.Stderr = old
	// Test passes if no data races (run with -race)
}

func TestNew_TracksMaxNameLen(t *testing.T) {
	resetState(t)
	New("a")
	if got := maxNameLen.Load(); got != 1 {
		t.Errorf("maxNameLen = %d, want 1", got)
	}
	New("auth:discovery")
	if got := maxNameLen.Load(); got != 14 {
		t.Errorf("maxNameLen = %d, want 14", got)
	}
	// Shorter name should not reduce maxNameLen
	New("b")
	if got := maxNameLen.Load(); got != 14 {
		t.Errorf("maxNameLen = %d, want 14 (should not shrink)", got)
	}
}

func BenchmarkLog_Disabled(b *testing.B) {
	// Reset state to ensure no patterns
	patterns = nil
	globalEnabled.Store(false)
	var a atomic.Bool
	_ = a.Load() // warm up

	l := New("bench:disabled")
	b.ResetTimer()
	for b.Loop() {
		l.Log("message")
	}
}

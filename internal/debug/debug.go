// Package debug provides namespaced debug logging inspired by npm's debug package.
//
// Loggers are created with a namespace (e.g. "auth:token", "http:request") and
// write to stderr only when enabled. Enable via:
//
//   - GLEAN_DEBUG env var with glob patterns: GLEAN_DEBUG=auth:* or GLEAN_DEBUG=*
//   - Programmatically via Enable() (used by the --verbose flag)
//
// When disabled, Log is a near-zero-cost no-op (single atomic bool check).
package debug

import (
	"fmt"
	"hash/fnv"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fatih/color"
)

var (
	globalEnabled atomic.Bool
	patterns      []pattern
	lastTime      atomic.Int64
	mu            sync.Mutex
	maxNameLen    atomic.Int32

	palette = []*color.Color{
		color.New(color.FgCyan),
		color.New(color.FgMagenta),
		color.New(color.FgYellow),
		color.New(color.FgGreen),
		color.New(color.FgBlue),
		color.New(color.FgRed),
	}
)

type pattern struct {
	negate bool
	prefix string // "auth:" for "auth:*", or full name for exact match
	glob   bool   // true if pattern ends with *
}

func init() {
	raw := os.Getenv("GLEAN_DEBUG")
	if raw == "" {
		return
	}
	for p := range strings.SplitSeq(raw, ",") {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		pat := pattern{}
		if strings.HasPrefix(p, "-") {
			pat.negate = true
			p = p[1:]
		}
		switch {
		case p == "*":
			pat.glob = true
			pat.prefix = ""
		case strings.HasSuffix(p, "*"):
			pat.glob = true
			pat.prefix = p[:len(p)-1]
		default:
			pat.prefix = p
		}
		patterns = append(patterns, pat)
	}
}

// Logger is a namespaced debug logger. The zero value is a disabled no-op logger.
type Logger struct {
	namespace string
	enabled   bool
	clr       *color.Color
}

// New creates a debug logger for the given namespace.
// The logger is enabled if GLEAN_DEBUG patterns match the namespace.
func New(namespace string) Logger {
	if n := int32(len(namespace)); n > maxNameLen.Load() {
		maxNameLen.Store(n)
	}
	return Logger{
		namespace: namespace,
		enabled:   matchesPatterns(namespace),
		clr:       pickColor(namespace),
	}
}

// Enable turns on all debug loggers globally. Called by --verbose flag handling.
// Loggers created before Enable() is called are retroactively activated.
func Enable() {
	globalEnabled.Store(true)
}

// Enabled reports whether this logger will produce output.
func (l Logger) Enabled() bool {
	return l.enabled || globalEnabled.Load()
}

// Log writes a debug message to stderr if this logger is enabled.
// Format and args follow fmt.Sprintf conventions.
func (l Logger) Log(format string, args ...any) {
	if !l.enabled && !globalEnabled.Load() {
		return
	}

	now := time.Now()
	prev := lastTime.Swap(now.UnixMilli())
	delta := time.Duration(0)
	if prev > 0 {
		delta = now.Sub(time.UnixMilli(prev))
	}

	msg := fmt.Sprintf(format, args...)
	deltaStr := formatDelta(delta)

	mu.Lock()
	defer mu.Unlock()

	padded := fmt.Sprintf("%-*s", maxNameLen.Load(), l.namespace)
	ns := l.clr.Sprint(padded)
	dt := l.clr.Sprint(deltaStr)
	fmt.Fprintf(os.Stderr, "  %s %s %s\n", ns, msg, dt)
}

func matchesPatterns(namespace string) bool {
	matched := false
	for _, p := range patterns {
		if p.matches(namespace) {
			matched = !p.negate
		}
	}
	return matched
}

func (p pattern) matches(namespace string) bool {
	if p.glob {
		return strings.HasPrefix(namespace, p.prefix)
	}
	return namespace == p.prefix
}

func pickColor(namespace string) *color.Color {
	h := fnv.New32a()
	h.Write([]byte(namespace))
	return palette[h.Sum32()%uint32(len(palette))]
}

func formatDelta(d time.Duration) string {
	switch {
	case d < time.Millisecond:
		return "+0ms"
	case d < time.Second:
		return fmt.Sprintf("+%dms", d.Milliseconds())
	case d < time.Minute:
		return fmt.Sprintf("+%.1fs", d.Seconds())
	default:
		return fmt.Sprintf("+%dm%ds", int(d.Minutes()), int(d.Seconds())%60)
	}
}

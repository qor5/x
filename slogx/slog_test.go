package slogx_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"

	"github.com/qor5/x/v3/slogx"
)

// withCapturedDefault redirects slogx's package writer to a buffer and
// resets slog.Default() at test cleanup. Returns the buffer for asserts.
//
// Note: SetupDefaultLogger mutates the process-wide slog default, so
// tests touching it MUST NOT run with t.Parallel().
func withCapturedDefault(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	prevWriter := slogx.SetDefaultWriter(&buf)
	savedDefault := slog.Default()
	t.Cleanup(func() {
		slogx.SetDefaultWriter(prevWriter)
		slog.SetDefault(savedDefault)
	})
	return &buf
}

// TestSetupDefaultLogger_DefaultsToText verifies that with no Format
// in Config the installed default handler is *slog.TextHandler.
// Regression guard against accidentally flipping the default.
func TestSetupDefaultLogger_DefaultsToText(t *testing.T) {
	withCapturedDefault(t)

	logger, err := slogx.SetupDefaultLogger(&slogx.Config{Level: slogx.LevelInfo})
	if err != nil {
		t.Fatalf("SetupDefaultLogger: %v", err)
	}
	if _, ok := logger.Handler().(*slog.TextHandler); !ok {
		t.Errorf("returned logger handler = %T, want *slog.TextHandler", logger.Handler())
	}
	if _, ok := slog.Default().Handler().(*slog.TextHandler); !ok {
		t.Errorf("slog.Default() handler = %T, want *slog.TextHandler", slog.Default().Handler())
	}
}

// TestSetupDefaultLogger_FormatText explicitly selects text and checks
// the captured output uses slog's key=value layout (not JSON).
func TestSetupDefaultLogger_FormatText(t *testing.T) {
	buf := withCapturedDefault(t)

	if _, err := slogx.SetupDefaultLogger(&slogx.Config{
		Level:  slogx.LevelInfo,
		Format: slogx.FormatText,
	}); err != nil {
		t.Fatalf("SetupDefaultLogger: %v", err)
	}
	slog.Info("hi", "key", "value")
	out := buf.String()

	if !strings.Contains(out, "level=INFO") || !strings.Contains(out, "key=value") {
		t.Errorf("text output missing expected tokens: %q", out)
	}
	if len(out) > 0 && out[0] == '{' {
		t.Errorf("text output should not look like JSON, got: %q", out)
	}
}

// TestSetupDefaultLogger_FormatJSON end-to-end: Format=json selects
// JSONHandler and the captured output parses as JSON with expected
// fields. Covers the production "I need machine-parseable logs" path.
func TestSetupDefaultLogger_FormatJSON(t *testing.T) {
	buf := withCapturedDefault(t)

	logger, err := slogx.SetupDefaultLogger(&slogx.Config{
		Level:  slogx.LevelInfo,
		Format: slogx.FormatJSON,
	})
	if err != nil {
		t.Fatalf("SetupDefaultLogger: %v", err)
	}
	if _, ok := logger.Handler().(*slog.JSONHandler); !ok {
		t.Errorf("handler = %T, want *slog.JSONHandler", logger.Handler())
	}
	slog.Info("hello", "key", "value", "n", 42)

	var rec map[string]any
	if err := json.Unmarshal(bytes.TrimSpace(buf.Bytes()), &rec); err != nil {
		t.Fatalf("output is not JSON: %v; raw=%q", err, buf.String())
	}
	if got := rec["level"]; got != "INFO" {
		t.Errorf("level = %v, want INFO", got)
	}
	if got := rec["msg"]; got != "hello" {
		t.Errorf("msg = %v, want hello", got)
	}
	if got := rec["key"]; got != "value" {
		t.Errorf("key = %v, want value", got)
	}
	if got, _ := rec["n"].(float64); got != 42 {
		t.Errorf("n = %v, want 42", got)
	}
}

// TestSetupDefaultLogger_RespectsLevel verifies the configured level
// filters lower-severity records.
func TestSetupDefaultLogger_RespectsLevel(t *testing.T) {
	withCapturedDefault(t)

	if _, err := slogx.SetupDefaultLogger(&slogx.Config{Level: slogx.LevelWarn}); err != nil {
		t.Fatalf("SetupDefaultLogger: %v", err)
	}
	ctx := context.Background()
	if slog.Default().Enabled(ctx, slog.LevelInfo) {
		t.Error("Info should be filtered when level is Warn")
	}
	if !slog.Default().Enabled(ctx, slog.LevelWarn) {
		t.Error("Warn should pass when level is Warn")
	}
	if !slog.Default().Enabled(ctx, slog.LevelError) {
		t.Error("Error should pass when level is Warn")
	}
}

// TestSetupDefaultLogger_InvalidLevel returns an error for an unknown
// level string rather than silently falling back. Guards against the
// confx validator being relaxed at some point.
func TestSetupDefaultLogger_InvalidLevel(t *testing.T) {
	withCapturedDefault(t)

	_, err := slogx.SetupDefaultLogger(&slogx.Config{Level: "verbose"})
	if err == nil {
		t.Error("expected error for invalid level, got nil")
	}
}

// TestSetupDefaultLogger_FormatInvalid an unknown format string must
// error out, mirroring the behavior for invalid Level. Code paths that
// hand-build Config (tests, internal tools bypassing confx) get caught
// at the Go API level instead of silently defaulting.
func TestSetupDefaultLogger_FormatInvalid(t *testing.T) {
	withCapturedDefault(t)

	_, err := slogx.SetupDefaultLogger(&slogx.Config{
		Level:  slogx.LevelInfo,
		Format: "verbose",
	})
	if err == nil {
		t.Error("expected error for invalid format, got nil")
	}
}

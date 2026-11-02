// Package slogx installs a default slog handler driven by confx config:
// text by default for readability, JSON via Format=json for production
// log aggregators.
package slogx

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

// Format selects the slog handler. Zero value = text.
type Format string

const (
	FormatJSON Format = "json"
	FormatText Format = "text"
)

type Config struct {
	Level  Level  `confx:"level" usage:"Logging level" validate:"required,oneof=debug info warn error"`
	Format Format `confx:"format" usage:"Logging format (json or text); empty = text" validate:"omitempty,oneof=json text"`
}

// defaultWriter is the destination for the default handler. Unexported
// so production callers can't rewire it; tests use the hook in
// export_test.go.
var defaultWriter io.Writer = os.Stderr

// SetupDefaultLogger installs a text (or JSON, if conf.Format=json)
// slog handler at the configured level as the process default and
// returns it. Safe to call multiple times; the last call wins.
func SetupDefaultLogger(conf *Config) (*slog.Logger, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(conf.Level)); err != nil {
		return nil, err
	}
	opts := &slog.HandlerOptions{Level: level}
	var handler slog.Handler
	switch conf.Format {
	case "", FormatText:
		handler = slog.NewTextHandler(defaultWriter, opts)
	case FormatJSON:
		handler = slog.NewJSONHandler(defaultWriter, opts)
	default:
		return nil, fmt.Errorf("slogx: unknown format %q", conf.Format)
	}
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(level)
	return logger, nil
}

package slogx

import (
	"log/slog"
)

type Level string

const (
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
)

type Config struct {
	Level Level `confx:"level" usage:"Logging level" validate:"required,oneof=debug info warn error"`
}

func SetupDefaultLogger(conf *Config) (*slog.Logger, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(conf.Level)); err != nil {
		return nil, err
	}
	slog.SetLogLoggerLevel(level)
	return slog.Default(), nil
}

package slogx

import (
	"log/slog"
)

type Config struct {
	Level string `confx:"level" usage:"Logging level" validate:"required,oneof=debug info warn error"`
}

func SetupDefaultLogger(conf *Config) (*slog.Logger, error) {
	var level slog.Level
	if err := level.UnmarshalText([]byte(conf.Level)); err != nil {
		return nil, err
	}
	slog.SetLogLoggerLevel(level)
	return slog.Default(), nil
}

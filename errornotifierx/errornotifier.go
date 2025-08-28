package errornotifierx

import (
	"io"

	"github.com/pkg/errors"
	"github.com/theplant/appkit/errornotifier"
	kitlog "github.com/theplant/appkit/log"
	"github.com/theplant/inject/lifecycle"
)

type Kind string

const (
	KindLog      Kind = "log"
	KindAirbrake Kind = "airbrake"
)

type Config struct {
	Kind     Kind           `confx:"kind" validate:"required,oneof=log airbrake" usage:"Error notifier type, can be 'log' or 'airbrake'"`
	Airbrake AirbrakeConfig `confx:"airbrake" validate:"skip_nested_unless=Kind airbrake" usage:"Airbrake configuration, required when kind is 'airbrake'"`
	Logger   *kitlog.Logger `confx:"-" json:"-" inject:""`
}

func Setup(lc *lifecycle.Lifecycle, conf *Config) (errornotifier.Notifier, error) {
	switch conf.Kind {
	case KindLog:
		if conf.Logger == nil {
			return nil, errors.New("logger is required when kind is log")
		}
		return errornotifier.NewLogNotifier(*conf.Logger), nil
	case KindAirbrake:
		return SetupAirbrakeNotifier(lc, &conf.Airbrake)
	default:
		return nil, errors.Errorf("invalid error notifier kind: %s", conf.Kind)
	}
}

type nopCloser struct{}

func (nopCloser) Close() error { return nil }

func New(c *Config) (errornotifier.Notifier, io.Closer, error) {
	if c == nil {
		return nil, nil, errors.New("config cannot be nil")
	}
	switch c.Kind {
	case KindLog:
		if c.Logger == nil {
			return nil, nil, errors.New("logger is required when kind is log")
		}
		return errornotifier.NewLogNotifier(*c.Logger), &nopCloser{}, nil
	case KindAirbrake:
		airbrakeCfg := errornotifier.AirbrakeConfig{
			ProjectID:   c.Airbrake.ProjectID,
			Token:       c.Airbrake.Token,
			Environment: c.Airbrake.Environment,
		}
		notifier, closer, err := errornotifier.NewAirbrakeNotifier(airbrakeCfg)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to create airbrake notifier")
		}
		return notifier, closer, nil
	default:
		return nil, nil, errors.Errorf("invalid error notifier kind: %s", c.Kind)
	}
}

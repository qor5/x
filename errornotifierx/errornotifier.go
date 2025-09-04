package errornotifierx

import (
	"io"

	"github.com/pkg/errors"
	"github.com/theplant/appkit/errornotifier"
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
}

func SetupNotifier(lc *lifecycle.Lifecycle, conf *Config) (errornotifier.Notifier, error) {
	switch conf.Kind {
	case KindLog:
		return SetupLogNotifier(lc)
	case KindAirbrake:
		return SetupAirbrakeNotifier(lc, &conf.Airbrake)
	default:
		return nil, errors.Errorf("invalid error notifier kind: %s", conf.Kind)
	}
}

func New(c *Config) (errornotifier.Notifier, io.Closer, error) {
	if c == nil {
		return nil, nil, errors.New("config cannot be nil")
	}
	switch c.Kind {
	case KindLog:
		notifier := NewLogNotifier()
		return notifier, notifier, nil
	case KindAirbrake:
		notifier, err := NewAirbrakeNotifier(errornotifier.AirbrakeConfig{
			ProjectID:   c.Airbrake.ProjectID,
			Token:       c.Airbrake.Token,
			Environment: c.Airbrake.Environment,
		})
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to create airbrake notifier")
		}
		return notifier, notifier, nil
	default:
		return nil, nil, errors.Errorf("invalid error notifier kind: %s", c.Kind)
	}
}

type Notifier interface {
	errornotifier.Notifier
	io.Closer
}

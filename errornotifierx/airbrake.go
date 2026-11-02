package errornotifierx

import (
	"context"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/theplant/appkit/errornotifier"
	"github.com/theplant/inject/lifecycle"
)

type AirbrakeConfig struct {
	ProjectID   int64  `confx:"projectID" validate:"required" usage:"Airbrake project ID"`
	Token       string `confx:"token" validate:"required,len=32" usage:"Airbrake project token, must be exactly 32 characters"`
	Environment string `confx:"environment" validate:"required,oneof=development test staging production" usage:"Environment name for error reporting (e.g., development, test, staging, production)"`
}

func SetupAirbrakeNotifier(lc *lifecycle.Lifecycle, conf *AirbrakeConfig) (errornotifier.Notifier, error) {
	notifier, err := NewAirbrakeNotifier(errornotifier.AirbrakeConfig{
		ProjectID:   conf.ProjectID,
		Token:       conf.Token,
		Environment: conf.Environment,
	})
	if err != nil {
		return nil, err
	}

	lc.Add(lifecycle.NewFuncActor(nil, func(_ context.Context) error {
		if err := notifier.Close(); err != nil {
			return errors.Wrap(err, "failed to close airbrake notifier")
		}
		return nil
	}).WithName("airbrake-notifier"))

	return notifier, nil
}

func NewAirbrakeNotifier(conf errornotifier.AirbrakeConfig) (Notifier, error) {
	notifier, closer, err := errornotifier.NewAirbrakeNotifier(conf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create airbrake notifier")
	}
	return &notifierWrapper{Notifier: notifier, closer: closer}, nil
}

type notifierWrapper struct {
	errornotifier.Notifier
	closer io.Closer
}

func (n *notifierWrapper) Notify(aErr any, r *http.Request, context map[string]any) {
	if err, ok := aErr.(error); ok {
		n.Notifier.Notify(newTrackedError(err), r, context)
	} else {
		n.Notifier.Notify(aErr, r, context)
	}
}

func (n *notifierWrapper) Close() error {
	return n.closer.Close()
}

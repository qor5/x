package healthz

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"github.com/qor5/x/v3/httpx"
	"github.com/qor5/x/v3/timex"
	"github.com/theplant/inject"
	"github.com/theplant/inject/lifecycle"
)

// Deprecated: use SetupReadinessProbe
var SetupWaitForReady = SetupReadinessProbe

func SetupReadinessProbe(lc *lifecycle.Lifecycle, httpListener httpx.Listener) *inject.Element[*lifecycle.ReadinessProbe] {
	endpoint := fmt.Sprintf("http://%s%s", httpListener.Addr().String(), Path)
	return SetupReadinessProbeFactory(endpoint)(lc)
}

func SetupReadinessProbeFactory(endpoint string) func(lc *lifecycle.Lifecycle) *inject.Element[*lifecycle.ReadinessProbe] {
	return func(lc *lifecycle.Lifecycle) *inject.Element[*lifecycle.ReadinessProbe] {
		probe := lifecycle.NewReadinessProbe()
		lc.Add(lifecycle.NewFuncActor(func(ctx context.Context) (xerr error) {
			defer func() { probe.Signal(xerr) }()
			return WaitForReady(ctx, endpoint)
		}, nil))
		return inject.NewElement(probe)
	}
}

func WaitForReady(ctx context.Context, endpoint string) error {
	if err := timex.Sleep(ctx, 10*time.Millisecond); err != nil {
		return err
	}
	client := &http.Client{Timeout: 2 * time.Second}
	err := backoff.Retry(
		func() error {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
			if err != nil {
				return err
			}

			resp, err := client.Do(req)
			if resp != nil {
				defer resp.Body.Close()
			}
			if err != nil {
				return errors.Wrapf(err, "failed to access ready check endpoint")
			}

			if resp.StatusCode != http.StatusOK {
				return errors.Errorf("unexpected status code: %d", resp.StatusCode)
			}
			return nil
		},
		backoff.WithContext(backoff.NewExponentialBackOff(func(eb *backoff.ExponentialBackOff) {
			eb.InitialInterval = 100 * time.Millisecond
			eb.RandomizationFactor = 0.5
			eb.Multiplier = 2
			eb.MaxInterval = 1 * time.Second
			eb.MaxElapsedTime = 0
		}), ctx),
	)
	if err != nil {
		return errors.Wrap(err, "failed to access ready check endpoint")
	}
	return nil
}

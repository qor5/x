package healthz

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/pkg/errors"
	"github.com/qor5/x/v3/httpx"
	"github.com/qor5/x/v3/netx"
	"github.com/qor5/x/v3/timex"
	"github.com/theplant/inject/lifecycle"
)

// Deprecated: use SetupReadinessProbe
var SetupWaitForReady = SetupReadinessProbe

// DefaultReadinessTimeout is the default timeout for the readiness probe.
var DefaultReadinessTimeout = 30 * time.Second

// ReadinessProbe is a marker type for readiness probe dependency.
// Other components can depend on this type to ensure they execute after WaitForReady completes.
type ReadinessProbe struct{}

// SetupReadinessProbe sets up a readiness probe that checks the health endpoint.
func SetupReadinessProbe(lc *lifecycle.Lifecycle, httpListener httpx.Listener) *ReadinessProbe {
	endpoint := fmt.Sprintf("http://%s%s", netx.ConnectableString(httpListener.Addr()), Path)
	return SetupReadinessProbeFactory(endpoint)(lc)
}

// SetupReadinessProbeFactory creates a factory for setting up a readiness probe with custom endpoint.
func SetupReadinessProbeFactory(endpoint string) func(lc *lifecycle.Lifecycle) *ReadinessProbe {
	return func(lc *lifecycle.Lifecycle) *ReadinessProbe {
		lc.Add(lifecycle.NewFuncActor(func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, DefaultReadinessTimeout)
			defer cancel()
			return WaitForReady(ctx, endpoint)
		}, nil).WithName("healthz-readiness-probe").WithReadiness())
		return &ReadinessProbe{}
	}
}

// WaitForReady waits until the endpoint responds with HTTP 200 OK.
func WaitForReady(ctx context.Context, endpoint string) error {
	if err := timex.Sleep(ctx, 10*time.Millisecond); err != nil {
		return err
	}
	client := &http.Client{Timeout: 2 * time.Second}
	err := backoff.Retry(
		func() error {
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
			if err != nil {
				return errors.Wrap(err, "failed to create request")
			}

			resp, err := client.Do(req)
			if resp != nil {
				defer resp.Body.Close()
			}
			if err != nil {
				return errors.Wrap(err, "failed to do request")
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

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
	"github.com/theplant/inject/lifecycle"
)

type setupWaitForReady struct{}

func SetupWaitForReady(lc *lifecycle.Lifecycle, httpListener httpx.Listener) *setupWaitForReady {
	endpoint := fmt.Sprintf("http://%s%s", httpListener.Addr().String(), Path)
	return SetupWaitForReadyFactory(endpoint)(lc)
}

func SetupWaitForReadyFactory(endpoint string) func(lc *lifecycle.Lifecycle) *setupWaitForReady {
	return func(lc *lifecycle.Lifecycle) *setupWaitForReady {
		lc.Add(lifecycle.NewFuncActor(func(ctx context.Context) error {
			return WaitForReady(ctx, endpoint)
		}, nil))
		return &setupWaitForReady{}
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
			if err != nil {
				return errors.Wrapf(err, "failed to access ready check endpoint")
			}
			defer resp.Body.Close()

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

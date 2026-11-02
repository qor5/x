package timex

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func Sleep(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return nil
	}
	if ctx == nil {
		time.Sleep(d)
		return nil
	}

	t := time.NewTimer(d)
	select {
	case <-ctx.Done():
		t.Stop()
		return errors.Wrap(ctx.Err(), "timex.Sleep")
	case <-t.C:
	}
	return nil
}

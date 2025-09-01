package netx

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"github.com/theplant/inject/lifecycle"
)

func SetupListenerFactory(name, address string) func(lc *lifecycle.Lifecycle) (net.Listener, error) {
	return func(lc *lifecycle.Lifecycle) (net.Listener, error) {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to listen on %s", address)
		}
		lc.Add(lifecycle.NewFuncActor(nil, func(ctx context.Context) error {
			if err := listener.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
				return errors.Wrap(err, "failed to close listener")
			}
			return nil
		}).WithName(name))
		return listener, nil
	}
}

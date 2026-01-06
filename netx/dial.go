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

		// Wrap listener to replace [::] with 127.0.0.1
		wrappedListener := &addrWrapper{Listener: listener}

		lc.Add(lifecycle.NewFuncActor(nil, func(ctx context.Context) error {
			if err := listener.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
				return errors.Wrap(err, "failed to close listener")
			}
			return nil
		}).WithName(name))
		return wrappedListener, nil
	}
}

type addrWrapper struct {
	net.Listener
}

func (w *addrWrapper) Addr() net.Addr {
	addr := w.Listener.Addr()
	if tcpAddr, ok := addr.(*net.TCPAddr); ok && tcpAddr.IP.IsUnspecified() {
		// Return a new TCPAddr with 127.0.0.1 instead of [::] or 0.0.0.0
		return &net.TCPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: tcpAddr.Port,
			Zone: tcpAddr.Zone,
		}
	}
	return addr
}

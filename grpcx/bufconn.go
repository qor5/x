package grpcx

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"github.com/theplant/inject/lifecycle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var SetupBufconnListener = SetupBufconnListenerFactory("bufconn-listener")

func SetupBufconnListenerFactory(name string) func(lc *lifecycle.Lifecycle) (Listener, error) {
	return func(lc *lifecycle.Lifecycle) (Listener, error) {
		listener := bufconn.Listen(1024 * 1024)
		lc.Add(lifecycle.NewFuncActor(nil, func(ctx context.Context) error {
			if err := listener.Close(); err != nil && !errors.Is(err, net.ErrClosed) {
				return errors.Wrap(err, "failed to close bufconn listener")
			}
			return nil
		}).WithName(name))
		return listener, nil
	}
}

func SetupBufconnFactory(name string, dialOpts ...grpc.DialOption) func(lc *lifecycle.Lifecycle, listener Listener) (*grpc.ClientConn, error) {
	return func(lc *lifecycle.Lifecycle, listener Listener) (*grpc.ClientConn, error) {
		l, ok := listener.(interface {
			DialContext(context.Context) (net.Conn, error)
		})
		if !ok {
			return nil, errors.New("listener is not a bufconn listener")
		}

		opts := []grpc.DialOption{
			grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
				return l.DialContext(ctx)
			}),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		conn, err := grpc.NewClient("passthrough://bufnet", append(opts, dialOpts...)...)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create gRPC buf connection")
		}

		lc.Add(lifecycle.NewFuncActor(nil, func(_ context.Context) error {
			if err := conn.Close(); err != nil {
				return errors.Wrap(err, "failed to close gRPC buf connection")
			}
			return nil
		}).WithName(name))

		return conn, nil
	}
}

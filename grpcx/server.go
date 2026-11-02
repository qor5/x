package grpcx

import (
	"context"
	"log/slog"
	"net"

	"github.com/pkg/errors"
	"github.com/qor5/x/v3/netx"
	"github.com/theplant/inject/lifecycle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ServerConfig struct {
	Address            string `confx:"address" usage:"gRPC server address" validate:"required"`
	RegisterReflection bool   `confx:"registerReflection" usage:"register the server reflection service"`
}

type Listener net.Listener

func SetupListener(lc *lifecycle.Lifecycle, conf *ServerConfig) (Listener, error) {
	return netx.SetupListenerFactory("grpc-listener", conf.Address)(lc)
}

func SetupServerFactory(name string, opts ...grpc.ServerOption) func(ctx context.Context, lc *lifecycle.Lifecycle, listener Listener, conf *ServerConfig) (*grpc.Server, error) {
	return func(ctx context.Context, lc *lifecycle.Lifecycle, listener Listener, conf *ServerConfig) (*grpc.Server, error) {
		grpcServer := grpc.NewServer(opts...)

		if conf.RegisterReflection {
			reflection.Register(grpcServer)
		}

		lc.Add(lifecycle.NewFuncService(func(_ context.Context) error {
			slog.InfoContext(ctx, "gRPC server listening", "address", listener.Addr().String())
			if err := grpcServer.Serve(listener); err != nil {
				return errors.Wrap(err, "failed to serve gRPC")
			}
			return nil
		}).WithStop(func(_ context.Context) error {
			grpcServer.GracefulStop()
			return nil
		}).WithName(name))

		return grpcServer, nil
	}
}

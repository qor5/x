package grpcx

import (
	"context"
	"net"

	"github.com/pkg/errors"
	"github.com/qor5/x/v3/netx"
	kitlog "github.com/theplant/appkit/log"
	"github.com/theplant/inject/lifecycle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Listener net.Listener

func SetupListener(lc *lifecycle.Lifecycle, conf *Config) (Listener, error) {
	return netx.SetupListenerFactory("grpc-listener", conf.Address)(lc)
}

func SetupServerFactory(name string, opts ...grpc.ServerOption) func(lc *lifecycle.Lifecycle, listener Listener, conf *Config, logger *kitlog.Logger) (*grpc.Server, error) {
	return func(lc *lifecycle.Lifecycle, listener Listener, conf *Config, logger *kitlog.Logger) (*grpc.Server, error) {
		grpcServer := grpc.NewServer(opts...)

		if conf.RegisterReflection {
			reflection.Register(grpcServer)
		}

		lc.Add(lifecycle.NewFuncService(func(_ context.Context) error {
			logger.Info().Log("msg", "gRPC server listening on %s", "addr", listener.Addr().String())
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

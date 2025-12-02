package grpcx

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/theplant/appkit/logtracing"
	"github.com/theplant/inject/lifecycle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LoadBalancingPolicy string

const (
	// pick-first is the default load balancing policy
	LoadBalancingPolicyPickFirst LoadBalancingPolicy = "pick-first"
	// round-robin should be used with dns:///address and the target address should be a k8s headless service
	LoadBalancingPolicyRoundRobin LoadBalancingPolicy = "round-robin"
)

type ConnConfig struct {
	Address             string              `confx:"address" validate:"required" usage:"gRPC server address"`
	LoadBalancingPolicy LoadBalancingPolicy `confx:"loadBalancingPolicy" validate:"oneof=pick-first round-robin" usage:"gRPC load balancing policy, pick-first or round-robin"`
}

func SetupConnFactory(name string, dialOpts ...grpc.DialOption) func(lc *lifecycle.Lifecycle, conf *ConnConfig) (*grpc.ClientConn, error) {
	return func(lc *lifecycle.Lifecycle, conf *ConnConfig) (*grpc.ClientConn, error) {
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithChainUnaryInterceptor(
				logtracing.UnaryClientInterceptor(),
			),
		}

		if conf.LoadBalancingPolicy == LoadBalancingPolicyRoundRobin {
			if !strings.HasPrefix(conf.Address, "dns:///") {
				return nil, errors.Errorf("load balancing policy %q requires dns:///address, got %q", conf.LoadBalancingPolicy, conf.Address)
			}
			opts = append(opts, grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`))
		}

		conn, err := grpc.NewClient(conf.Address, append(opts, dialOpts...)...)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create gRPC connection to %s", conf.Address)
		}

		lc.Add(lifecycle.NewFuncActor(nil, func(_ context.Context) error {
			if err := conn.Close(); err != nil {
				return errors.Wrap(err, "failed to close gRPC connection")
			}
			return nil
		}).WithName(name))

		return conn, nil
	}
}

package healthz

import (
	"context"

	"google.golang.org/grpc"

	"google.golang.org/grpc/health/grpc_health_v1"
)

func UnaryServerInterceptor(healthServer grpc_health_v1.HealthServer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// For health service methods, bypass other middlewares
		switch info.FullMethod {
		case grpc_health_v1.Health_Check_FullMethodName:
			if _, ok := req.(*grpc_health_v1.HealthCheckRequest); !ok {
				req = &grpc_health_v1.HealthCheckRequest{}
			}
			return healthServer.Check(ctx, req.(*grpc_health_v1.HealthCheckRequest))

		case grpc_health_v1.Health_List_FullMethodName:
			if _, ok := req.(*grpc_health_v1.HealthListRequest); !ok {
				req = &grpc_health_v1.HealthListRequest{}
			}
			return healthServer.List(ctx, req.(*grpc_health_v1.HealthListRequest))
		}

		return handler(ctx, req)
	}
}

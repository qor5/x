package grpcx

type Config struct {
	Address            string `confx:"address" usage:"gRPC server address" validate:"required"`
	RegisterReflection bool   `confx:"registerReflection" usage:"register the server reflection service"`
}

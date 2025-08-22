package statusx

import (
	"fmt"

	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"google.golang.org/grpc/codes"
)

func Internal(message string) *Status {
	return New(codes.Internal, statusv1.ErrorReason_INTERNAL.String(), message)
}

func Internalf(format string, a ...any) *Status {
	return Internal(fmt.Sprintf(format, a...))
}

func WrapInternal(err error, message string) *Status {
	return Wrap(err, codes.Internal, statusv1.ErrorReason_INTERNAL.String(), message)
}

func WrapInternalf(err error, format string, a ...any) *Status {
	return WrapInternal(err, fmt.Sprintf(format, a...))
}

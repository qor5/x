package statusx

import (
	"fmt"

	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"google.golang.org/grpc/codes"
)

// ReasonFromCode maps gRPC codes to their corresponding ErrorReason values
func ReasonFromCode(code codes.Code) statusv1.ErrorReason {
	switch code {
	case codes.OK:
		return statusv1.ErrorReason_OK
	case codes.Canceled:
		return statusv1.ErrorReason_CANCELED
	case codes.Unknown:
		return statusv1.ErrorReason_UNKNOWN
	case codes.InvalidArgument:
		return statusv1.ErrorReason_INVALID_ARGUMENT
	case codes.DeadlineExceeded:
		return statusv1.ErrorReason_DEADLINE_EXCEEDED
	case codes.NotFound:
		return statusv1.ErrorReason_NOT_FOUND
	case codes.AlreadyExists:
		return statusv1.ErrorReason_ALREADY_EXISTS
	case codes.PermissionDenied:
		return statusv1.ErrorReason_PERMISSION_DENIED
	case codes.ResourceExhausted:
		return statusv1.ErrorReason_RESOURCE_EXHAUSTED
	case codes.FailedPrecondition:
		return statusv1.ErrorReason_FAILED_PRECONDITION
	case codes.Aborted:
		return statusv1.ErrorReason_ABORTED
	case codes.OutOfRange:
		return statusv1.ErrorReason_OUT_OF_RANGE
	case codes.Unimplemented:
		return statusv1.ErrorReason_UNIMPLEMENTED
	case codes.Internal:
		return statusv1.ErrorReason_INTERNAL
	case codes.Unavailable:
		return statusv1.ErrorReason_UNAVAILABLE
	case codes.DataLoss:
		return statusv1.ErrorReason_DATA_LOSS
	case codes.Unauthenticated:
		return statusv1.ErrorReason_UNAUTHENTICATED
	default:
		return statusv1.ErrorReason_UNKNOWN
	}
}

// NewCode creates a Status with the specified code and message.
// The reason is automatically derived from the code.
func NewCode(code codes.Code, message string) *Status {
	return New(code, ReasonFromCode(code).String(), message)
}

// NewCodef creates a Status with the specified code and formatted message.
// The reason is automatically derived from the code.
func NewCodef(code codes.Code, format string, a ...any) *Status {
	return NewCode(code, fmt.Sprintf(format, a...))
}

// WrapCode wraps an existing error with the specified code and message.
// The reason is automatically derived from the code.
func WrapCode(err error, code codes.Code, message string) *Status {
	return Wrap(err, code, ReasonFromCode(code).String(), message)
}

// WrapCodef wraps an existing error with the specified code and formatted message.
// The reason is automatically derived from the code.
func WrapCodef(err error, code codes.Code, format string, a ...any) *Status {
	return WrapCode(err, code, fmt.Sprintf(format, a...))
}

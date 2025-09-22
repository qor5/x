package statusx

import (
	"testing"

	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestReasonFromCode(t *testing.T) {
	testCases := []struct {
		name     string
		code     codes.Code
		expected statusv1.ErrorReason
	}{
		{"OK", codes.OK, statusv1.ErrorReason_OK},
		{"Canceled", codes.Canceled, statusv1.ErrorReason_CANCELED},
		{"Unknown", codes.Unknown, statusv1.ErrorReason_UNKNOWN},
		{"InvalidArgument", codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT},
		{"DeadlineExceeded", codes.DeadlineExceeded, statusv1.ErrorReason_DEADLINE_EXCEEDED},
		{"NotFound", codes.NotFound, statusv1.ErrorReason_NOT_FOUND},
		{"AlreadyExists", codes.AlreadyExists, statusv1.ErrorReason_ALREADY_EXISTS},
		{"PermissionDenied", codes.PermissionDenied, statusv1.ErrorReason_PERMISSION_DENIED},
		{"ResourceExhausted", codes.ResourceExhausted, statusv1.ErrorReason_RESOURCE_EXHAUSTED},
		{"FailedPrecondition", codes.FailedPrecondition, statusv1.ErrorReason_FAILED_PRECONDITION},
		{"Aborted", codes.Aborted, statusv1.ErrorReason_ABORTED},
		{"OutOfRange", codes.OutOfRange, statusv1.ErrorReason_OUT_OF_RANGE},
		{"Unimplemented", codes.Unimplemented, statusv1.ErrorReason_UNIMPLEMENTED},
		{"Internal", codes.Internal, statusv1.ErrorReason_INTERNAL},
		{"Unavailable", codes.Unavailable, statusv1.ErrorReason_UNAVAILABLE},
		{"DataLoss", codes.DataLoss, statusv1.ErrorReason_DATA_LOSS},
		{"Unauthenticated", codes.Unauthenticated, statusv1.ErrorReason_UNAUTHENTICATED},
		{"Undefined code", codes.Code(99), statusv1.ErrorReason_UNKNOWN}, // Default case
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReasonFromCode(tc.code)
			assert.Equal(t, tc.expected, result)
		})
	}
}

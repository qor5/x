package statusx

import (
	"errors"
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

func TestNewAutoReason(t *testing.T) {
	t.Run("creates status with derived reason", func(t *testing.T) {
		status := New(codes.NotFound, "Resource not found")

		assert.Equal(t, codes.NotFound, status.Code())
		assert.Equal(t, statusv1.ErrorReason_NOT_FOUND.String(), status.Reason())
		assert.Equal(t, "Resource not found", status.Message())
	})

	t.Run("handles OK code", func(t *testing.T) {
		status := New(codes.OK, "Success")

		assert.Equal(t, codes.OK, status.Code())
		assert.Equal(t, statusv1.ErrorReason_OK.String(), status.Reason())
		assert.Equal(t, "Success", status.Message())
	})
}

func TestNewfAutoReason(t *testing.T) {
	t.Run("creates status with formatted message", func(t *testing.T) {
		status := Newf(codes.InvalidArgument, "Invalid %s: %d", "user ID", 123)

		assert.Equal(t, codes.InvalidArgument, status.Code())
		assert.Equal(t, statusv1.ErrorReason_INVALID_ARGUMENT.String(), status.Reason())
		assert.Equal(t, "Invalid user ID: 123", status.Message())
	})

	t.Run("handles empty format arguments", func(t *testing.T) {
		status := Newf(codes.Internal, "Internal error")

		assert.Equal(t, codes.Internal, status.Code())
		assert.Equal(t, "Internal error", status.Message())
	})
}

func TestWrapAutoReason(t *testing.T) {
	t.Run("wraps error with derived reason", func(t *testing.T) {
		originalErr := errors.New("original error")
		status := Wrap(originalErr, codes.PermissionDenied, "Access denied")

		assert.Equal(t, codes.PermissionDenied, status.Code())
		assert.Equal(t, statusv1.ErrorReason_PERMISSION_DENIED.String(), status.Reason())
		assert.Equal(t, "Access denied", status.Message())
		assert.True(t, errors.Is(status.Err(), originalErr))
	})

	t.Run("handles nil error", func(t *testing.T) {
		status := Wrap(nil, codes.Internal, "Internal error")

		assert.Equal(t, codes.OK, status.Code())
		assert.Equal(t, statusv1.ErrorReason_OK.String(), status.Reason())
		assert.Equal(t, "", status.Message())
	})
}

func TestWrapf(t *testing.T) {
	t.Run("wraps error with formatted message", func(t *testing.T) {
		originalErr := errors.New("database connection failed")
		status := Wrapf(originalErr, codes.Unavailable, "Service unavailable: %s", "database down")

		assert.Equal(t, codes.Unavailable, status.Code())
		assert.Equal(t, statusv1.ErrorReason_UNAVAILABLE.String(), status.Reason())
		assert.Equal(t, "Service unavailable: database down", status.Message())
		assert.True(t, errors.Is(status.Err(), originalErr))
	})
}

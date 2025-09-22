package statusx

import (
	"testing"

	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

func TestAssertErrorFieldViolation(t *testing.T) {
	t.Run("validates error with expected field violations", func(t *testing.T) {
		// Create error with field violations directly using New + WithDetails
		expectedFV := &errdetails.BadRequest_FieldViolation{
			Field:       "email",
			Description: "Email is required",
			Reason:      "REQUIRED",
		}
		err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "invalid argument").
			WithDetails(&errdetails.BadRequest{FieldViolations: []*errdetails.BadRequest_FieldViolation{expectedFV}}).
			GRPCStatus().Err()

		// This should not panic or fail assertions
		AssertFieldViolations(t, err, expectedFV)
	})

	t.Run("works with multiple field violations", func(t *testing.T) {
		fv1 := &errdetails.BadRequest_FieldViolation{
			Field:       "email",
			Description: "Email is required",
			Reason:      "REQUIRED",
		}
		fv2 := &errdetails.BadRequest_FieldViolation{
			Field:       "password",
			Description: "Password is too short",
			Reason:      "TOO_SHORT",
		}
		err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "invalid argument").
			WithDetails(&errdetails.BadRequest{FieldViolations: []*errdetails.BadRequest_FieldViolation{fv1, fv2}}).
			GRPCStatus().Err()

		AssertFieldViolations(t, err, fv1, fv2)
	})
}

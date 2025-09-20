package statusx

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/qor5/x/v3/i18nx"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestNew(t *testing.T) {
	t.Run("basic properties", func(t *testing.T) {
		s := New(codes.InvalidArgument, "invalid input provided").WithReason("INVALID_INPUT")
		require.NotNil(t, s)

		assert.Equal(t, codes.InvalidArgument, s.Code())
		assert.Equal(t, "INVALID_INPUT", s.Reason())
		assert.Equal(t, "invalid input provided", s.Message())
	})

	t.Run("stacktrace behavior", func(t *testing.T) {
		// Test OK status has no stacktrace
		s := New(codes.OK, "success")
		require.NotNil(t, s)
		assert.Nil(t, s.cause)

		// Test non-OK status has stacktrace
		s = New(codes.Internal, "error")
		require.NotNil(t, s)
		require.NotNil(t, s.cause)

		// Verify stacktrace exists
		stackErr := s.cause
		assert.Contains(t, fmt.Sprintf("%+v", stackErr), "github.com/qor5/x/v3/statusx.TestNew")
	})
}

func TestNewf(t *testing.T) {
	s := Newf(codes.NotFound, "resource %s not found", "user")
	require.NotNil(t, s)

	assert.Equal(t, codes.NotFound, s.Code())
	assert.Equal(t, "NOT_FOUND", s.Reason())
	assert.Equal(t, "resource user not found", s.Message())
}

func TestErr(t *testing.T) {
	err := Error(codes.PermissionDenied, "permission denied")
	require.NotNil(t, err)

	s, ok := FromError(err)
	require.True(t, ok)
	require.NotNil(t, s)

	assert.Equal(t, codes.PermissionDenied, s.Code())
	assert.Equal(t, "PERMISSION_DENIED", s.Reason())
	assert.Equal(t, "permission denied", s.Message())

	require.NotNil(t, s.Err())
	assert.True(t, errors.Is(s.Err(), err))
}

func TestErrorf(t *testing.T) {
	err := Errorf(codes.InvalidArgument, "invalid input provided for %s", "username")
	require.NotNil(t, err)

	s, ok := FromError(err)
	require.True(t, ok)
	require.NotNil(t, s)

	assert.Equal(t, codes.InvalidArgument, s.Code())
	assert.Equal(t, "INVALID_ARGUMENT", s.Reason())
	assert.Equal(t, "invalid input provided for username", s.Message())
}

func TestWithMetadata(t *testing.T) {
	md := map[string]string{"field": "username", "error": "missing"}
	s := New(codes.InvalidArgument, "required field is missing").WithReason("MISSING_FIELD").WithMetadata(md)

	require.NotNil(t, s)
	assert.Equal(t, md, s.Metadata())
}

func TestWithLocalized(t *testing.T) {
	s := New(codes.InvalidArgument, "invalid input").WithLocalized("error.invalid_input", "username", "required")

	require.NotNil(t, s)
	localized := s.Localized()
	assert.NotNil(t, localized)
	assert.Equal(t, "error.invalid_input", localized.Key)

	// Extract values from Any for comparison
	require.Len(t, localized.Args, 2)
	assert.Equal(t, "username", extractStringFromAny(t, localized.Args[0]))
	assert.Equal(t, "required", extractStringFromAny(t, localized.Args[1]))
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrapped := Wrap(originalErr, codes.Internal, "internal server error").WithReason("INTERNAL_ERROR")
	require.NotNil(t, wrapped)

	assert.Equal(t, codes.Internal, wrapped.Code())
	assert.Equal(t, "INTERNAL_ERROR", wrapped.Reason())
	assert.Equal(t, "internal server error", wrapped.Message())
	assert.True(t, errors.Is(wrapped.Err(), originalErr))

	{
		wrapped := Wrap(nil, codes.Internal, "internal server error")
		assert.NotNil(t, wrapped)
		assert.Equal(t, codes.OK, wrapped.Code())
		assert.Equal(t, statusv1.ErrorReason_OK.String(), wrapped.Reason())
		assert.Equal(t, "", wrapped.Message())
	}

	{
		wrapped := Wrap(originalErr, codes.OK, "success")
		assert.NotNil(t, wrapped)
		// Because the cause is not nil, the code will be Unknown finally.
		assert.Equal(t, codes.Unknown, wrapped.Code())
		assert.Equal(t, statusv1.ErrorReason_UNKNOWN.String(), wrapped.Reason())
		assert.Equal(t, "success", wrapped.Message())
	}

	{
		wrapped := Wrapf(originalErr, codes.Internal, "internal server error for %s", "user").WithReason("INTERNAL_ERROR")
		assert.NotNil(t, wrapped)
		assert.Equal(t, codes.Internal, wrapped.Code())
		assert.Equal(t, "INTERNAL_ERROR", wrapped.Reason())
		assert.Equal(t, "internal server error for user", wrapped.Message())
	}

	{
		status := New(codes.NotFound, "resource not found")
		wrapped := Wrap(status.Err(), codes.Internal, "internal server error")
		assert.Equal(t, wrapped, status)
		assert.Equal(t, codes.NotFound, wrapped.Code())
		assert.Equal(t, "NOT_FOUND", wrapped.Reason())
		assert.Equal(t, "resource not found", wrapped.Message())
	}

	{
		status, _ := status.New(codes.NotFound, "resource not found").WithDetails(&errdetails.ErrorInfo{
			Reason: "NOT_FOUND",
		})
		wrapped := Wrap(status.Err(), codes.Internal, "internal server error")
		assert.Equal(t, codes.NotFound, wrapped.Code())
		assert.Equal(t, "NOT_FOUND", wrapped.Reason())
		assert.Equal(t, "resource not found", wrapped.Message())
	}
}

func TestGRPCStatus(t *testing.T) {
	s := New(codes.PermissionDenied, "permission denied").
		WithLocalized("error.permission_denied", "access", "user").
		WithMetadata(map[string]string{"reason": "policy_violation"})

	st := s.GRPCStatus()
	require.NotNil(t, st)

	details := st.Details()
	require.Len(t, details, 2)

	errorInfo := details[0].(*errdetails.ErrorInfo)
	require.Equal(t, "PERMISSION_DENIED", errorInfo.Reason)
	require.Equal(t, map[string]string{"reason": "policy_violation"}, errorInfo.Metadata)

	localized := details[1].(*statusv1.Localized)
	require.Equal(t, "error.permission_denied", localized.Key)
	require.Len(t, localized.Args, 2)
	assert.Equal(t, "access", extractStringFromAny(t, localized.Args[0]))
	assert.Equal(t, "user", extractStringFromAny(t, localized.Args[1]))

	{
		s := New(codes.OK, "ok")
		st := s.GRPCStatus()
		require.NotNil(t, st)
		assert.Empty(t, st.Details())
		assert.Equal(t, codes.OK, st.Code())
		assert.Empty(t, st.Message())
	}
}

func TestFromError(t *testing.T) {
	s := New(codes.Unauthenticated, "authentication required").
		WithLocalized("error.unauthenticated", "access", "user")
	ss := s.GRPCStatus()
	grpcErr := ss.Err()

	converted, ok := FromError(grpcErr)
	require.True(t, ok)
	require.NotNil(t, converted)

	assert.Equal(t, codes.Unauthenticated, converted.Code())
	assert.Equal(t, "UNAUTHENTICATED", converted.Reason())
	assert.Equal(t, "authentication required", converted.Message())

	{
		ss, err := ss.WithDetails(&errdetails.LocalizedMessage{
			Locale:  "en-US",
			Message: "authentication required",
		})
		require.NoError(t, err)

		s, ok := FromError(ss.Err())
		require.True(t, ok)

		assert.Equal(t, codes.Unauthenticated, s.Code())
		assert.Equal(t, "UNAUTHENTICATED", s.Reason())
		assert.Equal(t, "authentication required", s.Message())
		assert.Len(t, s.details, 1)
		assert.True(t, proto.Equal(s.details[0], &errdetails.LocalizedMessage{
			Locale:  "en-US",
			Message: "authentication required",
		}))

		cloned := Clone(s)
		require.NotNil(t, cloned)
		assert.True(t, proto.Equal(cloned.GRPCStatus().Proto(), s.GRPCStatus().Proto()))

		ss, ok = status.FromError(s.Err())
		require.True(t, ok)
		assert.True(t, proto.Equal(ss.Proto(), s.GRPCStatus().Proto()))
	}
}

func TestClone(t *testing.T) {
	s := New(codes.InvalidArgument, "invalid input provided").WithReason("INVALID_INPUT").
		WithMetadata(map[string]string{"field": "username"}).
		WithLocalized("error.invalid_input", "username", "required").
		WithCause(errors.New("original error"))

	cloned := Clone(s)
	require.NotNil(t, cloned)

	assert.Equal(t, s.Code(), cloned.Code())
	assert.Equal(t, s.Reason(), cloned.Reason())
	assert.Equal(t, s.Message(), cloned.Message())
	assert.Equal(t, s.Metadata(), cloned.Metadata())
	assert.Equal(t, s.Localized(), cloned.Localized())
	assert.True(t, errors.Is(cloned.Err(), s.Cause()))

	{
		cloned := Clone(nil)
		assert.Nil(t, cloned)
	}
}

func TestNil(t *testing.T) {
	var s *Status
	assert.Equal(t, codes.OK, s.Code())
	assert.Equal(t, "", s.Message())
	assert.Equal(t, statusv1.ErrorReason_OK.String(), s.Reason())
	assert.Nil(t, s.Metadata())
	assert.Nil(t, s.Localized())
	assert.Nil(t, s.Cause())
	assert.Nil(t, s.Err())
	assert.Nil(t, s.GRPCStatus())
}

func TestErrorFormat(t *testing.T) {
	s := New(codes.InvalidArgument, "invalid input provided").WithReason("INVALID_INPUT").
		WithMetadata(map[string]string{"field": "username"}).
		WithLocalized("error.invalid_input", "username", "required").
		WithCause(errors.New("original error"))

	{
		formatted := fmt.Sprintf("%+v", s.Err())
		assert.Contains(t, formatted, "original error")
		assert.Contains(t, formatted, "rpc error: code = InvalidArgument reason = INVALID_INPUT message = invalid input provided")
		assert.Contains(t, formatted, "statusx.TestErrorFormat")
	}
	{
		formatted := fmt.Sprintf("%v", s.Err())
		assert.Contains(t, formatted, "rpc error: code = InvalidArgument reason = INVALID_INPUT message = invalid input provided: original error")
	}
	{
		formatted := fmt.Sprintf("%s", s.Err())
		assert.Contains(t, formatted, "rpc error: code = InvalidArgument reason = INVALID_INPUT message = invalid input provided: original error")
	}
	{
		formatted := fmt.Sprintf("%q", s.Err())
		assert.Contains(t, formatted, `"rpc error: code = InvalidArgument reason = INVALID_INPUT message = invalid input provided: original error"`)
	}
}

func TestCode(t *testing.T) {
	assert.Equal(t, codes.OK, Code(nil))
	assert.Equal(t, codes.Unknown, Code(errors.New("original error")))

	err := New(codes.InvalidArgument, "invalid input provided")
	assert.Equal(t, codes.InvalidArgument, Code(err.Err()))

	err = New(codes.OK, "success").WithCause(errors.New("original error"))
	assert.Equal(t, codes.Unknown, Code(err.Err()))
}

func TestReason(t *testing.T) {
	assert.Equal(t, statusv1.ErrorReason_OK.String(), Reason(nil))
	assert.Equal(t, statusv1.ErrorReason_UNKNOWN.String(), Reason(errors.New("original error")))

	err := New(codes.InvalidArgument, "invalid input provided").WithReason("INVALID_INPUT")
	assert.Equal(t, "INVALID_INPUT", Reason(err.Err()))
}

func TestConvert(t *testing.T) {
	{
		err := errors.New("original error")
		s := Convert(err)
		require.NotNil(t, s)
		assert.Equal(t, codes.Unknown, s.Code())
		assert.Equal(t, statusv1.ErrorReason_UNKNOWN.String(), s.Reason())
		assert.Equal(t, err.Error(), s.Message())
	}
	{
		err := errors.Wrap(context.Canceled, "original error")
		s := Convert(err)
		require.NotNil(t, s)
		assert.Equal(t, codes.Canceled, s.Code())
		assert.Equal(t, statusv1.ErrorReason_CANCELED.String(), s.Reason())
		assert.Equal(t, err.Error(), s.Message())
	}
	{
		err := errors.Wrap(context.DeadlineExceeded, "original error")
		s := Convert(err)
		require.NotNil(t, s)
		assert.Equal(t, codes.DeadlineExceeded, s.Code())
		assert.Equal(t, statusv1.ErrorReason_DEADLINE_EXCEEDED.String(), s.Reason())
		assert.Equal(t, err.Error(), s.Message())
	}
}

// extractStringFromAny is a test helper to extract string values from protobuf Any
func extractStringFromAny(t *testing.T, anyVal *anypb.Any) string {
	t.Helper()
	if anyVal == nil {
		return ""
	}

	if anyVal.MessageIs(&wrapperspb.StringValue{}) {
		var val wrapperspb.StringValue
		err := anyVal.UnmarshalTo(&val)
		require.NoError(t, err)
		return val.GetValue()
	}

	t.Fatalf("Expected StringValue in Any, got %s", anyVal.GetTypeUrl())
	return ""
}

func TestWithLocalizedWellKnownTypes(t *testing.T) {
	now := time.Now()
	duration := 5 * time.Minute
	structData := map[string]any{
		"user": "john",
		"age":  30,
	}

	s := New(codes.InvalidArgument, "validation failed").
		WithLocalized("validation.essential",
			"string_arg",    // 1. string
			int64(123),      // 2. int64
			int32(456),      // 3. int32
			uint64(789),     // 4. uint64
			uint32(101),     // 5. uint32
			float32(3.14),   // 6. float32
			float64(2.718),  // 7. float64
			true,            // 8. bool
			[]byte("bytes"), // 9. bytes
			now,             // 10. time.Time -> Timestamp
			duration,        // 11. time.Duration -> Duration
			structData,      // 12. map[string]any -> Struct
			nil,             // 13. nil -> Empty
		)

	require.NotNil(t, s)
	localized := s.Localized()
	assert.NotNil(t, localized)
	assert.Equal(t, "validation.essential", localized.Key)
	require.Len(t, localized.Args, 13)

	// Test that all types are properly stored and can be extracted
	args := localized.Args
	assert.Equal(t, "string_arg", extractStringFromAny(t, args[0]))
	// Note: We could add more specific extraction helpers for other types,
	// but the main point is that they're stored without errors
}

func TestWithLocalizedUnsupportedType(t *testing.T) {
	// Test that unsupported types cause panic
	assert.Panics(t, func() {
		New(codes.InvalidArgument, "test").WithLocalized("test.key", make(chan int)) // channels are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, "test").WithLocalized("test.key", func() {}) // functions are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, "test").WithLocalized("test.key", []string{"a", "b"}) // arrays are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, "test").WithLocalized("test.key", []any{"a", "b"}) // slices are not supported
	})
}

func TestWithDetailsTypeValidation(t *testing.T) {
	// Valid usage should work
	t.Run("valid details", func(t *testing.T) {
		err := New(codes.InvalidArgument, "test").
			WithDetails(&errdetails.BadRequest{
				FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{Field: "email", Description: "invalid"},
				},
			}).Err()

		assert.NotNil(t, err)

		// Should have the BadRequest detail
		details := status.Convert(err).Details()
		badRequestFound := false
		for _, detail := range details {
			if _, ok := detail.(*errdetails.BadRequest); ok {
				badRequestFound = true
				break
			}
		}
		assert.True(t, badRequestFound, "should contain BadRequest detail")
	})

	// Should panic when trying to add reserved types
	t.Run("panic on statusv1.Localized", func(t *testing.T) {
		assert.Panics(t, func() {
			New(codes.InvalidArgument, "test").
				WithDetails(&statusv1.Localized{
					Key: "test.key",
				})
		}, "should panic when adding *statusv1.Localized")
	})

	t.Run("panic on errdetails.ErrorInfo", func(t *testing.T) {
		assert.Panics(t, func() {
			New(codes.InvalidArgument, "test").
				WithDetails(&errdetails.ErrorInfo{
					Reason: "TEST_REASON",
				})
		}, "should panic when adding *errdetails.ErrorInfo")
	})
}

func TestWithFieldViolationsAPI(t *testing.T) {
	t.Run("single field via WithFieldViolations", func(t *testing.T) {
		status := New(codes.InvalidArgument, "validation failed").
			WithFieldViolations(&FieldViolation{Field: "email", Reason: "REQUIRED", Localized: &Localized{Key: "field.email.required"}})

		assert.NotNil(t, status.badRequest)
		assert.Len(t, status.badRequest.FieldViolations, 1)

		field := status.badRequest.FieldViolations[0]
		assert.Equal(t, "email", field.Field)
		assert.Equal(t, "REQUIRED", field.Reason)
	})

	t.Run("WithFieldViolations with args data", func(t *testing.T) {
		data := map[string]any{"minLength": 8}
		status := New(codes.InvalidArgument, "validation failed").
			WithFieldViolations(&FieldViolation{Field: "password", Reason: "TOO_WEAK", Localized: &Localized{Key: "field.password.min_length", Args: []any{data}}})

		field := status.badRequest.FieldViolations[0]
		assert.Equal(t, "password", field.Field)
		assert.Equal(t, "TOO_WEAK", field.Reason)
	})

	t.Run("batch WithFieldViolations", func(t *testing.T) {
		status := New(codes.InvalidArgument, "multiple field errors").
			WithFieldViolations(
				&FieldViolation{Field: "email", Reason: "REQUIRED", Localized: &Localized{Key: "field.email.required"}},
				&FieldViolation{Field: "password", Reason: "WEAK", Localized: &Localized{Key: "field.password.weak", Args: []any{map[string]any{"strength": "low"}}}},
				&FieldViolation{Field: "age", Reason: "INVALID_RANGE", Localized: &Localized{Key: "field.age.invalid_range", Args: []any{18, 65}}},
			)

		assert.NotNil(t, status.badRequest)
		assert.Len(t, status.badRequest.FieldViolations, 3)
		assert.Equal(t, "email", status.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "password", status.badRequest.FieldViolations[1].Field)
		assert.Equal(t, "age", status.badRequest.FieldViolations[2].Field)
	})

	t.Run("combine WithLocalized and WithFieldViolations", func(t *testing.T) {
		status := New(codes.InvalidArgument, "validation failed").
			WithLocalized("global.validation_failed").
			WithFieldViolations(
				&FieldViolation{Field: "email", Reason: "INVALID", Localized: &Localized{Key: "field.email.invalid"}},
				&FieldViolation{Field: "phone", Reason: "INVALID", Localized: &Localized{Key: "field.phone.invalid"}},
			)

		assert.NotNil(t, status.Localized())
		assert.Equal(t, "global.validation_failed", status.Localized().Key)
		assert.Len(t, status.badRequest.FieldViolations, 2)
	})

	t.Run("multiple WithFieldViolations calls", func(t *testing.T) {
		status := New(codes.InvalidArgument, "validation failed").
			WithFieldViolations(&FieldViolation{Field: "email", Reason: "INVALID", Localized: &Localized{Key: "field.email.invalid"}}).
			WithFieldViolations(&FieldViolation{Field: "password", Reason: "WEAK", Localized: &Localized{Key: "field.password.weak"}})

		assert.NotNil(t, status.badRequest)
		assert.Len(t, status.badRequest.FieldViolations, 2)

		assert.Equal(t, "email", status.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "password", status.badRequest.FieldViolations[1].Field)
	})

	t.Run("FieldViolation Proto mapping", func(t *testing.T) {
		fv := &FieldViolation{Field: "email", Reason: "INVALID", Localized: &Localized{Key: "field.email.invalid", Args: []any{"test@example.com", 25}}}
		proto := fv.Proto()
		assert.Equal(t, "email", proto.Field)
		assert.Equal(t, "field.email.invalid", proto.Localized.Key)
		assert.Len(t, proto.Localized.Args, 2)
	})

	t.Run("WithFieldsLocalized duplicate field validation", func(t *testing.T) {
		t.Run("duplicate within new fields", func(t *testing.T) {
			assert.Panics(t, func() {
				New(codes.InvalidArgument, "validation failed").
					WithFieldViolations(
						&FieldViolation{Field: "email", Reason: "REQUIRED"},
						&FieldViolation{Field: "password", Reason: "WEAK"},
						&FieldViolation{Field: "email", Reason: "INVALID"},
					)
			}, "should panic when duplicate field names are provided")
		})

		t.Run("replacement preserves original position with WithFieldViolations", func(t *testing.T) {
			// First add some fields
			status := New(codes.InvalidArgument, "validation failed").
				WithFieldViolations(
					&FieldViolation{Field: "email", Reason: "REQUIRED", Localized: &Localized{Key: "field.email.required"}},
					&FieldViolation{Field: "password", Reason: "WEAK", Localized: &Localized{Key: "field.password.weak"}},
					&FieldViolation{Field: "age", Reason: "REQUIRED", Localized: &Localized{Key: "field.age.required"}},
				)

				// Verify initial order
			assert.Len(t, status.badRequest.FieldViolations, 3)
			assert.Equal(t, "email", status.badRequest.FieldViolations[0].Field)
			assert.Equal(t, "password", status.badRequest.FieldViolations[1].Field)
			assert.Equal(t, "age", status.badRequest.FieldViolations[2].Field)

			// Replace middle field and add new field
			updatedStatus := status.WithFieldViolations(
				&FieldViolation{Field: "phone", Reason: "INVALID", Localized: &Localized{Key: "field.phone.invalid"}},
				&FieldViolation{Field: "password", Reason: "STRONG", Localized: &Localized{Key: "field.password.strong"}},
			)

			// Should have 4 fields total
			assert.Len(t, updatedStatus.badRequest.FieldViolations, 4)

			// Verify order is preserved for existing fields, new fields appended
			fields := updatedStatus.badRequest.FieldViolations
			assert.Equal(t, "email", fields[0].Field)
			assert.Equal(t, "field.email.required", fields[0].Localized.Key)

			assert.Equal(t, "password", fields[1].Field)
			assert.Equal(t, "field.password.strong", fields[1].Localized.Key) // Replaced in place

			assert.Equal(t, "age", fields[2].Field)
			assert.Equal(t, "field.age.required", fields[2].Localized.Key)

			assert.Equal(t, "phone", fields[3].Field) // New field appended
			assert.Equal(t, "field.phone.invalid", fields[3].Localized.Key)
		})

		t.Run("WithFieldViolations should replace existing field", func(t *testing.T) {
			status := New(codes.InvalidArgument, "validation failed").
				WithFieldViolations(&FieldViolation{Field: "email", Reason: "REQUIRED", Localized: &Localized{Key: "field.email.required"}})

			// Replace the existing field
			updatedStatus := status.WithFieldViolations(&FieldViolation{Field: "email", Reason: "FORMAT_ERROR", Localized: &Localized{Key: "field.email.format_error"}})

			assert.Len(t, updatedStatus.badRequest.FieldViolations, 1)
			assert.Equal(t, "email", updatedStatus.badRequest.FieldViolations[0].Field)
			assert.Equal(t, "field.email.format_error", updatedStatus.badRequest.FieldViolations[0].Localized.Key)
		})

		t.Run("duplicate within new fields should still panic", func(t *testing.T) {
			status := New(codes.InvalidArgument, "validation failed").
				WithFieldViolations(&FieldViolation{Field: "email", Reason: "REQUIRED", Localized: &Localized{Key: "field.email.required"}})

			assert.Panics(t, func() {
				status.WithFieldViolations(
					&FieldViolation{Field: "password", Reason: "WEAK"},
					&FieldViolation{Field: "phone", Reason: "INVALID"},
					&FieldViolation{Field: "password", Reason: "SHORT"},
				)
			}, "should panic when duplicate field names exist within new fields")
		})

		t.Run("new fields maintain input order", func(t *testing.T) {
			status := New(codes.InvalidArgument, "validation failed").
				WithFieldViolations(&FieldViolation{Field: "email", Reason: "REQUIRED", Localized: &Localized{Key: "field.email.required"}})

			// Add multiple new fields in specific order
			updatedStatus := status.WithFieldViolations(
				&FieldViolation{Field: "phone", Reason: "INVALID", Localized: &Localized{Key: "field.phone.invalid"}},
				&FieldViolation{Field: "address", Reason: "REQUIRED", Localized: &Localized{Key: "field.address.required"}},
				&FieldViolation{Field: "age", Reason: "INVALID", Localized: &Localized{Key: "field.age.invalid"}},
				&FieldViolation{Field: "country", Reason: "REQUIRED", Localized: &Localized{Key: "field.country.required"}},
			)

			// Should have 5 fields: email (existing) + 4 new fields
			fields := updatedStatus.badRequest.FieldViolations
			assert.Len(t, fields, 5)
			assert.Equal(t, "email", fields[0].Field) // Existing field preserved in position

			// New fields should be appended in the exact order they were provided
			assert.Equal(t, "phone", fields[1].Field)
			assert.Equal(t, "address", fields[2].Field)
			assert.Equal(t, "age", fields[3].Field)
			assert.Equal(t, "country", fields[4].Field)
		})

		t.Run("FieldViolation with Description and Reason", func(t *testing.T) {
			fieldViolation := &FieldViolation{
				Field:       "email",
				Localized:   &Localized{Key: "field.email.invalid_format", Args: []any{"user@example", "email format"}},
				Description: "Email address format is invalid",
				Reason:      "INVALID_EMAIL_FORMAT",
			}

			proto := fieldViolation.Proto()
			assert.Equal(t, "email", proto.Field)
			assert.Equal(t, "field.email.invalid_format", proto.Localized.Key)
			assert.Len(t, proto.Localized.Args, 2)
			assert.Equal(t, "Email address format is invalid", proto.Description)
			assert.Equal(t, "INVALID_EMAIL_FORMAT", proto.Reason)
		})

		t.Run("New WithFieldViolation API", func(t *testing.T) {
			status := New(codes.InvalidArgument, "validation failed").
				WithFieldViolations(&FieldViolation{Field: "email", Reason: "INVALID_EMAIL_FORMAT", Description: "Email address format is invalid", Localized: &Localized{Key: "field.email.invalid_format", Args: []any{"user@example"}}})

			assert.NotNil(t, status.badRequest)
			assert.Len(t, status.badRequest.FieldViolations, 1)

			fieldViol := status.badRequest.FieldViolations[0]
			assert.Equal(t, "email", fieldViol.Field)
			assert.Equal(t, "INVALID_EMAIL_FORMAT", fieldViol.Reason)
			assert.Equal(t, "Email address format is invalid", fieldViol.Description)
			assert.Equal(t, "field.email.invalid_format", fieldViol.Localized.Key)
			assert.Len(t, fieldViol.Localized.Args, 1)
		})

		t.Run("New WithFieldViolations API", func(t *testing.T) {
			violations := []*FieldViolation{
				{
					Field:       "email",
					Reason:      "REQUIRED",
					Description: "Email is required",
					Localized:   &Localized{Key: "field.email.required"},
				},
				{
					Field:       "password",
					Reason:      "TOO_WEAK",
					Description: "Password is too weak",
					Localized:   &Localized{Key: "field.password.weak", Args: []any{"minLength", 8}},
				},
			}

			status := New(codes.InvalidArgument, "validation failed").
				WithFieldViolations(violations...)

			assert.NotNil(t, status.badRequest)
			assert.Len(t, status.badRequest.FieldViolations, 2)

			// Check first violation
			emailViol := status.badRequest.FieldViolations[0]
			assert.Equal(t, "email", emailViol.Field)
			assert.Equal(t, "REQUIRED", emailViol.Reason)
			assert.Equal(t, "Email is required", emailViol.Description)

			// Check second violation
			passwordViol := status.badRequest.FieldViolations[1]
			assert.Equal(t, "password", passwordViol.Field)
			assert.Equal(t, "TOO_WEAK", passwordViol.Reason)
			assert.Equal(t, "Password is too weak", passwordViol.Description)
			assert.Len(t, passwordViol.Localized.Args, 2)
		})
	})
}

func TestWithFlattenFieldViolations(t *testing.T) {
	t.Run("mixed types in single call", func(t *testing.T) {
		// Create various inputs
		single := &FieldViolation{Field: "email", Reason: "INVALID_FORMAT"}
		slice := []*FieldViolation{
			{Field: "name", Reason: "REQUIRED"},
			{Field: "age", Reason: "TOO_YOUNG"},
		}

		// Create another status with violations
		otherStatus := New(codes.InvalidArgument, "nested validation").
			WithFieldViolations(&FieldViolation{Field: "nested.field", Reason: "INVALID"})

		// Use WithFlattenFieldViolations
		status := New(codes.InvalidArgument, "validation failed").
			WithFlattenFieldViolations(single, slice, otherStatus.ToFieldViolations("parent"))

		// Verify results
		assert.NotNil(t, status.badRequest)
		// Should have: single + 2 from slice + 1 main + 1 nested from other status = 5 total
		assert.Len(t, status.badRequest.FieldViolations, 5)

		// Verify all fields are present
		fieldNames := make([]string, 0, 5)
		for _, fv := range status.badRequest.FieldViolations {
			fieldNames = append(fieldNames, fv.Field)
		}
		assert.Contains(t, fieldNames, "email")
		assert.Contains(t, fieldNames, "name")
		assert.Contains(t, fieldNames, "age")
		assert.Contains(t, fieldNames, "parent")
		assert.Contains(t, fieldNames, "parent.nested.field")
	})

	t.Run("error input conversion", func(t *testing.T) {
		// Create a status error
		statusErr := New(codes.InvalidArgument, "validation failed").
			WithFieldViolations(
				&FieldViolation{Field: "field1", Reason: "INVALID"},
				&FieldViolation{Field: "field2", Reason: "REQUIRED"},
			).Err()

		// Use error via ToFieldViolations in WithFlattenFieldViolations
		status := New(codes.InvalidArgument, "main validation error").
			WithFlattenFieldViolations(ToFieldViolations(statusErr, "nested"), &FieldViolation{Field: "direct", Reason: "INVALID"})

		// Verify results
		assert.NotNil(t, status.badRequest)
		// Should have: 1 main error + 2 nested errors + 1 direct = 4 total
		assert.Len(t, status.badRequest.FieldViolations, 4)
	})

	t.Run("delegation to existing method", func(t *testing.T) {
		// Test that the new method delegates properly to WithFieldViolations
		// and maintains the same behavior for duplicate field handling

		existing := New(codes.InvalidArgument, "base").
			WithFieldViolations(&FieldViolation{Field: "email", Reason: "ORIGINAL"})

		// Add same field via WithFlattenFieldViolations - should replace
		updated := existing.WithFlattenFieldViolations(&FieldViolation{Field: "email", Reason: "UPDATED"})

		// Verify delegation behavior
		assert.Len(t, updated.badRequest.FieldViolations, 1)
		assert.Equal(t, "email", updated.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "UPDATED", updated.badRequest.FieldViolations[0].Reason)
	})

	t.Run("nil input handling", func(t *testing.T) {
		status := New(codes.InvalidArgument, "test").
			WithFlattenFieldViolations(nil, (*FieldViolation)(nil), []*FieldViolation(nil))

		// Should handle nil inputs gracefully
		if status.badRequest != nil {
			assert.Len(t, status.badRequest.FieldViolations, 0)
		}
	})
}

func TestWithFieldViolationsDuplicateStrategy(t *testing.T) {
	t.Run("design rationale - same call vs cross call duplicates", func(t *testing.T) {
		// ✅ Scenario 1: Same-call duplicates should panic (programming error)
		assert.Panics(t, func() {
			New(codes.InvalidArgument, "test").WithFieldViolations(
				&FieldViolation{Field: "email", Reason: "REQUIRED"},
				&FieldViolation{Field: "email", Reason: "INVALID"}, // Same field in same call - clear mistake
			)
		}, "Same-call duplicates should panic as this is clearly a programming error")

		// ✅ Scenario 2: Cross-call duplicates should overwrite (update semantics)
		base := New(codes.InvalidArgument, "test").WithFieldViolations(
			&FieldViolation{Field: "email", Reason: "REQUIRED"},
			&FieldViolation{Field: "name", Reason: "TOO_SHORT"},
		)

		updated := base.WithFieldViolations(
			&FieldViolation{Field: "email", Reason: "INVALID_FORMAT"}, // Different call - update existing
		)

		// Should have 2 fields total, email should be updated, name should remain
		assert.Len(t, updated.badRequest.FieldViolations, 2)

		// Email should be updated but maintain its original position (index 0)
		assert.Equal(t, "email", updated.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "INVALID_FORMAT", updated.badRequest.FieldViolations[0].Reason)

		// Name should remain unchanged
		assert.Equal(t, "name", updated.badRequest.FieldViolations[1].Field)
		assert.Equal(t, "TOO_SHORT", updated.badRequest.FieldViolations[1].Reason)
	})

	t.Run("position preservation demonstrates intent", func(t *testing.T) {
		// Position preservation shows this is intentional "update" not "replace"
		status := New(codes.InvalidArgument, "test").WithFieldViolations(
			&FieldViolation{Field: "field1", Reason: "ERROR1"},
			&FieldViolation{Field: "field2", Reason: "ERROR2"},
			&FieldViolation{Field: "field3", Reason: "ERROR3"},
		)

		// Update the middle field
		updated := status.WithFieldViolations(
			&FieldViolation{Field: "field2", Reason: "UPDATED_ERROR"},
		)

		// Verify order is preserved: field2 stays in position 1
		assert.Len(t, updated.badRequest.FieldViolations, 3)
		assert.Equal(t, "field1", updated.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "field2", updated.badRequest.FieldViolations[1].Field)
		assert.Equal(t, "UPDATED_ERROR", updated.badRequest.FieldViolations[1].Reason)
		assert.Equal(t, "field3", updated.badRequest.FieldViolations[2].Field)
	})
}

func TestTranslated(t *testing.T) {
	ib, _ := i18nx.New(strings.NewReader(`
key,en,zh-CN
invalid_request,Invalid request,无效请求
email.invalid_format,Invalid email format,邮箱格式错误
field.password.min_length,Password must be at least %d characters,密码至少需要 %d 个字符
VALIDATION_FAILED,Validation failed,验证失败
REQUIRED,Required field,必填字段
`))

	t.Run("translate main message with localized template", func(t *testing.T) {
		status := New(codes.InvalidArgument, "original message").
			WithLocalized("invalid_request")

		translated := status.Translated(ib, language.MustParse("zh-CN"))

		// Should have LocalizedMessage in details
		details := translated.Details()
		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](details)
		assert.NotNil(t, localizedMsg)
		assert.Equal(t, "zh-CN", localizedMsg.Locale)
		assert.Equal(t, "无效请求", localizedMsg.Message)

		// Original localized should be cleared
		assert.Nil(t, translated.Localized())
	})

	t.Run("translate main message with reason fallback", func(t *testing.T) {
		status := New(codes.InvalidArgument, "original message").WithReason("VALIDATION_FAILED")

		translated := status.Translated(ib, language.MustParse("zh-CN"))

		details := translated.Details()
		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](details)
		assert.NotNil(t, localizedMsg)
		assert.Equal(t, "zh-CN", localizedMsg.Locale)
		assert.Equal(t, "验证失败", localizedMsg.Message)
	})

	t.Run("translate field violations with localized templates", func(t *testing.T) {
		status := New(codes.InvalidArgument, "validation failed").
			WithFieldViolations(
				&FieldViolation{
					Field:     "email",
					Reason:    "INVALID_FORMAT",
					Localized: &Localized{Key: "email.invalid_format"},
				},
				&FieldViolation{
					Field:     "password",
					Reason:    "TOO_SHORT",
					Localized: &Localized{Key: "field.password.min_length", Args: []any{8}},
				},
			)

		translated := status.Translated(ib, language.MustParse("zh-CN"))

		details := translated.Details()
		badRequest := ExtractDetail[*errdetails.BadRequest](details)
		assert.NotNil(t, badRequest)
		assert.Len(t, badRequest.FieldViolations, 2)

		// Check first field violation
		emailViolation := badRequest.FieldViolations[0]
		assert.Equal(t, "email", emailViolation.Field)
		assert.NotNil(t, emailViolation.LocalizedMessage)
		assert.Equal(t, "zh-CN", emailViolation.LocalizedMessage.Locale)
		assert.Equal(t, "邮箱格式错误", emailViolation.LocalizedMessage.Message)

		// Check second field violation with arguments
		passwordViolation := badRequest.FieldViolations[1]
		assert.Equal(t, "password", passwordViolation.Field)
		assert.NotNil(t, passwordViolation.LocalizedMessage)
		assert.Equal(t, "密码至少需要 8 个字符", passwordViolation.LocalizedMessage.Message)

		// Original badRequest should be cleared
		assert.Nil(t, translated.badRequest)
	})

	t.Run("skip translation when LocalizedMessage already exists", func(t *testing.T) {
		status := New(codes.InvalidArgument, "original")

		// Add existing LocalizedMessage
		preTranslated := status.WithDetails(&errdetails.LocalizedMessage{
			Locale:  "en",
			Message: "Pre-translated message",
		})

		// Should not translate again
		result := preTranslated.Translated(ib, language.MustParse("zh-CN"))

		details := result.Details()
		var localizedMsgs []*errdetails.LocalizedMessage
		for _, d := range details {
			if localized, ok := d.(*errdetails.LocalizedMessage); ok {
				localizedMsgs = append(localizedMsgs, localized)
			}
		}

		// Should still have only the original LocalizedMessage
		assert.Len(t, localizedMsgs, 1)
		localizedMsg := localizedMsgs[0]
		assert.Equal(t, "en", localizedMsg.Locale)
		assert.Equal(t, "Pre-translated message", localizedMsg.Message)
	})

	t.Run("translateMainMessage and translateFieldViolations are idempotent", func(t *testing.T) {
		status := New(codes.InvalidArgument, "validation failed").
			WithLocalized("invalid_request").
			WithFieldViolations(&FieldViolation{
				Field:     "email",
				Reason:    "REQUIRED",
				Localized: &Localized{Key: "email.invalid_format"},
			})

		// First translation
		translated1 := status.Translated(ib, language.MustParse("zh-CN"))

		// Second translation should not change anything
		translated2 := translated1.Translated(ib, language.MustParse("zh-CN"))

		// Should be identical (no double-translation)
		assert.Equal(t, len(translated1.Details()), len(translated2.Details()))
	})

	t.Run("nil status returns nil", func(t *testing.T) {
		var status *Status
		result := status.Translated(ib, language.MustParse("zh-CN"))
		assert.Nil(t, result)
	})
}

func TestStatusMethods(t *testing.T) {
	t.Run("Details method", func(t *testing.T) {
		status := New(codes.InvalidArgument, "test")

		// Status automatically adds ErrorInfo, so details are not nil
		details := status.Details()
		assert.NotEmpty(t, details, "should have ErrorInfo automatically")

		// Verify ErrorInfo is present
		errorInfo := ExtractDetail[*errdetails.ErrorInfo](details)
		assert.NotNil(t, errorInfo, "should contain ErrorInfo")

		// Add additional details
		status = status.WithDetails(&errdetails.LocalizedMessage{
			Locale:  "en",
			Message: "test message",
		})
		details = status.Details()
		assert.Greater(t, len(details), 1, "should have more details after adding LocalizedMessage")

		// Check that our LocalizedMessage is present
		localizedMessage := ExtractDetail[*errdetails.LocalizedMessage](details)
		assert.NotNil(t, localizedMessage, "should contain the LocalizedMessage we added")
		assert.Equal(t, "en", localizedMessage.GetLocale())
		assert.Equal(t, "test message", localizedMessage.GetMessage())
	})

	t.Run("BadRequest method", func(t *testing.T) {
		// Nil status
		var nilStatus *Status
		assert.Nil(t, nilStatus.BadRequest())

		// Status without BadRequest
		status := New(codes.InvalidArgument, "test")
		assert.Nil(t, status.BadRequest())

		// Status with BadRequest
		status = status.WithFieldViolations(
			&FieldViolation{
				Field:       "email",
				Reason:      "REQUIRED",
				Description: "Email is required",
			},
		)
		badRequest := status.BadRequest()
		assert.NotNil(t, badRequest)
		assert.Len(t, badRequest.FieldViolations, 1)
		assert.Equal(t, "email", badRequest.FieldViolations[0].Field)
	})

	t.Run("ToFieldViolations method delegation", func(t *testing.T) {
		// Nil status
		var nilStatus *Status
		assert.Nil(t, nilStatus.ToFieldViolations("test"))

		// Valid status
		status := New(codes.InvalidArgument, "Validation failed").
			WithFieldViolations(
				&FieldViolation{Field: "name", Reason: "REQUIRED", Description: "Name required"},
			)
		violations := status.ToFieldViolations("user")

		assert.Len(t, violations, 2) // Main + nested
		assert.Equal(t, "user", violations[0].Field)
		assert.Equal(t, "user.name", violations[1].Field)
	})

	t.Run("WithCode method", func(t *testing.T) {
		originalStatus := New(codes.InvalidArgument, "test message").WithReason("TEST")
		newStatus := originalStatus.WithCode(codes.NotFound)

		// Original unchanged
		assert.Equal(t, codes.InvalidArgument, originalStatus.Code())

		// New status has updated code
		assert.Equal(t, codes.NotFound, newStatus.Code())
		assert.Equal(t, "TEST", newStatus.Reason())          // Reason unchanged
		assert.Equal(t, "test message", newStatus.Message()) // Message unchanged
	})

	t.Run("WithReason method", func(t *testing.T) {
		originalStatus := New(codes.InvalidArgument, "test message").WithReason("OLD_REASON")
		newStatus := originalStatus.WithReason("NEW_REASON")

		// Original unchanged
		assert.Equal(t, "OLD_REASON", originalStatus.Reason())

		// New status has updated reason
		assert.Equal(t, "NEW_REASON", newStatus.Reason())
		assert.Equal(t, codes.InvalidArgument, newStatus.Code()) // Code unchanged
		assert.Equal(t, "test message", newStatus.Message())     // Message unchanged
	})

	t.Run("Status method for StatusError", func(t *testing.T) {
		status := New(codes.NotFound, "Resource not found")
		err := status.Err().(*StatusError)

		grpcStatus := err.Status()
		assert.Equal(t, codes.NotFound, grpcStatus.Code())
		assert.Equal(t, "Resource not found", grpcStatus.Message())
	})
}

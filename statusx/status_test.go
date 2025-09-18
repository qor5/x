package statusx

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestNew(t *testing.T) {
	t.Run("basic properties", func(t *testing.T) {
		s := New(codes.InvalidArgument, "INVALID_INPUT", "invalid input provided")
		require.NotNil(t, s)

		assert.Equal(t, codes.InvalidArgument, s.Code())
		assert.Equal(t, "INVALID_INPUT", s.Reason())
		assert.Equal(t, "invalid input provided", s.Message())
	})

	t.Run("stacktrace behavior", func(t *testing.T) {
		// Test OK status has no stacktrace
		s := New(codes.OK, "OK", "success")
		require.NotNil(t, s)
		assert.Nil(t, s.cause)

		// Test non-OK status has stacktrace
		s = New(codes.Internal, "INTERNAL", "error")
		require.NotNil(t, s)
		require.NotNil(t, s.cause)

		// Verify stacktrace exists
		stackErr := s.cause
		assert.Contains(t, fmt.Sprintf("%+v", stackErr), "github.com/qor5/x/v3/statusx.TestNew")
	})
}

func TestNewf(t *testing.T) {
	s := Newf(codes.NotFound, "NOT_FOUND", "resource %s not found", "user")
	require.NotNil(t, s)

	assert.Equal(t, codes.NotFound, s.Code())
	assert.Equal(t, "NOT_FOUND", s.Reason())
	assert.Equal(t, "resource user not found", s.Message())
}

func TestErr(t *testing.T) {
	err := Error(codes.PermissionDenied, "PERMISSION_DENIED", "permission denied")
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
	err := Errorf(codes.InvalidArgument, "INVALID_INPUT", "invalid input provided for %s", "username")
	require.NotNil(t, err)

	s, ok := FromError(err)
	require.True(t, ok)
	require.NotNil(t, s)

	assert.Equal(t, codes.InvalidArgument, s.Code())
	assert.Equal(t, "INVALID_INPUT", s.Reason())
	assert.Equal(t, "invalid input provided for username", s.Message())
}

func TestWithMetadata(t *testing.T) {
	md := map[string]string{"field": "username", "error": "missing"}
	s := New(codes.InvalidArgument, "MISSING_FIELD", "required field is missing").WithMetadata(md)

	require.NotNil(t, s)
	assert.Equal(t, md, s.Metadata())
}

func TestWithLocalized(t *testing.T) {
	s := New(codes.InvalidArgument, "INVALID_INPUT", "invalid input").WithLocalized("error.invalid_input", "username", "required")

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
	wrapped := Wrap(originalErr, codes.Internal, "INTERNAL_ERROR", "internal server error")
	require.NotNil(t, wrapped)

	assert.Equal(t, codes.Internal, wrapped.Code())
	assert.Equal(t, "INTERNAL_ERROR", wrapped.Reason())
	assert.Equal(t, "internal server error", wrapped.Message())
	assert.True(t, errors.Is(wrapped.Err(), originalErr))

	{
		wrapped := Wrap(nil, codes.Internal, "INTERNAL_ERROR", "internal server error")
		assert.NotNil(t, wrapped)
		assert.Equal(t, codes.OK, wrapped.Code())
		assert.Equal(t, statusv1.ErrorReason_OK.String(), wrapped.Reason())
		assert.Equal(t, "", wrapped.Message())
	}

	{
		wrapped := Wrap(originalErr, codes.OK, "OK", "success")
		assert.NotNil(t, wrapped)
		// Because the cause is not nil, the code will be Unknown finally.
		assert.Equal(t, codes.Unknown, wrapped.Code())
		assert.Equal(t, statusv1.ErrorReason_UNKNOWN.String(), wrapped.Reason())
		assert.Equal(t, "success", wrapped.Message())
	}

	{
		wrapped := Wrapf(originalErr, codes.Internal, "INTERNAL_ERROR", "internal server error for %s", "user")
		assert.NotNil(t, wrapped)
		assert.Equal(t, codes.Internal, wrapped.Code())
		assert.Equal(t, "INTERNAL_ERROR", wrapped.Reason())
		assert.Equal(t, "internal server error for user", wrapped.Message())
	}

	{
		status := New(codes.NotFound, "NOT_FOUND", "resource not found")
		wrapped := Wrap(status.Err(), codes.Internal, "INTERNAL_ERROR", "internal server error")
		assert.Equal(t, wrapped, status)
		assert.Equal(t, codes.NotFound, wrapped.Code())
		assert.Equal(t, "NOT_FOUND", wrapped.Reason())
		assert.Equal(t, "resource not found", wrapped.Message())
	}

	{
		status, _ := status.New(codes.NotFound, "resource not found").WithDetails(&errdetails.ErrorInfo{
			Reason: "NOT_FOUND",
		})
		wrapped := Wrap(status.Err(), codes.Internal, "INTERNAL_ERROR", "internal server error")
		assert.Equal(t, codes.NotFound, wrapped.Code())
		assert.Equal(t, "NOT_FOUND", wrapped.Reason())
		assert.Equal(t, "resource not found", wrapped.Message())
	}
}

func TestGRPCStatus(t *testing.T) {
	s := New(codes.PermissionDenied, "PERMISSION_DENIED", "permission denied").
		WithLocalized("error.permission_denied", "access", "user").
		WithMetadata(map[string]string{"reason": "policy_violation"})

	st := s.GRPCStatus()
	require.NotNil(t, st)

	details := st.Details()
	require.Len(t, details, 2)

	errorInfo := details[0].(*errdetails.ErrorInfo)
	require.Equal(t, "PERMISSION_DENIED", errorInfo.Reason)
	require.Equal(t, map[string]string{"reason": "policy_violation"}, errorInfo.Metadata)

	localizedMessage := details[1].(*statusv1.Localized)
	require.Equal(t, "error.permission_denied", localizedMessage.Key)
	require.Len(t, localizedMessage.Args, 2)
	assert.Equal(t, "access", extractStringFromAny(t, localizedMessage.Args[0]))
	assert.Equal(t, "user", extractStringFromAny(t, localizedMessage.Args[1]))

	{
		s := New(codes.OK, "XXXX", "yyyy")
		st := s.GRPCStatus()
		require.NotNil(t, st)
		assert.Empty(t, st.Details())
		assert.Equal(t, codes.OK, st.Code())
		assert.Empty(t, st.Message())
	}
}

func TestFromError(t *testing.T) {
	s := New(codes.Unauthenticated, "UNAUTHENTICATED", "authentication required").
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
	s := New(codes.InvalidArgument, "INVALID_INPUT", "invalid input provided").
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
	s := New(codes.InvalidArgument, "INVALID_INPUT", "invalid input provided").
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

	err := New(codes.InvalidArgument, "INVALID_INPUT", "invalid input provided")
	assert.Equal(t, codes.InvalidArgument, Code(err.Err()))

	err = New(codes.OK, "OK", "success").WithCause(errors.New("original error"))
	assert.Equal(t, codes.Unknown, Code(err.Err()))
}

func TestReason(t *testing.T) {
	assert.Equal(t, statusv1.ErrorReason_OK.String(), Reason(nil))
	assert.Equal(t, statusv1.ErrorReason_UNKNOWN.String(), Reason(errors.New("original error")))

	err := New(codes.InvalidArgument, "INVALID_INPUT", "invalid input provided")
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

	s := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
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
		New(codes.InvalidArgument, "TEST", "test").
			WithLocalized("test.key", make(chan int)) // channels are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, "TEST", "test").
			WithLocalized("test.key", func() {}) // functions are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, "TEST", "test").
			WithLocalized("test.key", []string{"a", "b"}) // arrays are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, "TEST", "test").
			WithLocalized("test.key", []any{"a", "b"}) // slices are not supported
	})
}

func TestWithDetailsTypeValidation(t *testing.T) {
	// Valid usage should work
	t.Run("valid details", func(t *testing.T) {
		err := New(codes.InvalidArgument, "TEST", "test").
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
			New(codes.InvalidArgument, "TEST", "test").
				WithDetails(&statusv1.Localized{
					Key: "test.key",
				})
		}, "should panic when adding *statusv1.Localized")
	})

	t.Run("panic on errdetails.ErrorInfo", func(t *testing.T) {
		assert.Panics(t, func() {
			New(codes.InvalidArgument, "TEST", "test").
				WithDetails(&errdetails.ErrorInfo{
					Reason: "TEST_REASON",
				})
		}, "should panic when adding *errdetails.ErrorInfo")
	})
}

func TestWithFieldsLocalizedAPI(t *testing.T) {
	t.Run("WithFieldLocalized single field", func(t *testing.T) {
		status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
			WithFieldLocalized("email", "field.email.required")

		assert.NotNil(t, status.localized)
		assert.Len(t, status.localized.FieldViolations, 1)

		fieldLoc := status.localized.FieldViolations[0]
		assert.Equal(t, "email", fieldLoc.Field)
		assert.Equal(t, "field.email.required", fieldLoc.LocalizedKey)
		assert.Empty(t, fieldLoc.LocalizedArgs)
	})

	t.Run("WithFieldLocalized with args", func(t *testing.T) {
		data := map[string]any{"minLength": 8}
		status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
			WithFieldLocalized("password", "field.password.min_length", data)

		fieldLoc := status.localized.FieldViolations[0]
		assert.Equal(t, "password", fieldLoc.Field)
		assert.Equal(t, "field.password.min_length", fieldLoc.LocalizedKey)
		assert.Len(t, fieldLoc.LocalizedArgs, 1)
	})

	t.Run("WithFieldsLocalized batch operation", func(t *testing.T) {
		fields := []*FieldLocalized{
			{
				Field: "email",
				Key:   "field.email.required",
			},
			{
				Field: "password",
				Key:   "field.password.weak",
				Args:  []any{map[string]any{"strength": "low"}},
			},
			{
				Field: "age",
				Key:   "field.age.invalid_range",
				Args:  []any{18, 65},
			},
		}

		status := New(codes.InvalidArgument, "VALIDATION_ERROR", "multiple field errors").
			WithFieldsLocalized(fields...)

		assert.NotNil(t, status.localized)
		assert.Len(t, status.localized.FieldViolations, 3)

		// Check first field
		field1 := status.localized.FieldViolations[0]
		assert.Equal(t, "email", field1.Field)
		assert.Equal(t, "field.email.required", field1.LocalizedKey)
		assert.Empty(t, field1.LocalizedArgs)

		// Check second field
		field2 := status.localized.FieldViolations[1]
		assert.Equal(t, "password", field2.Field)
		assert.Equal(t, "field.password.weak", field2.LocalizedKey)
		assert.Len(t, field2.LocalizedArgs, 1)

		// Check third field
		field3 := status.localized.FieldViolations[2]
		assert.Equal(t, "age", field3.Field)
		assert.Equal(t, "field.age.invalid_range", field3.LocalizedKey)
		assert.Len(t, field3.LocalizedArgs, 2)
	})

	t.Run("combine WithLocalized and WithFieldsLocalized", func(t *testing.T) {
		fields := []*FieldLocalized{
			{Field: "email", Key: "field.email.invalid"},
			{Field: "phone", Key: "field.phone.invalid"},
		}

		status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
			WithLocalized("global.validation_failed").
			WithFieldsLocalized(fields...)

		assert.NotNil(t, status.localized)
		assert.Equal(t, "global.validation_failed", status.localized.Key)
		assert.Len(t, status.localized.FieldViolations, 2)
	})

	t.Run("multiple calls to WithFieldLocalized", func(t *testing.T) {
		status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
			WithFieldLocalized("email", "field.email.invalid").
			WithFieldLocalized("password", "field.password.weak")

		assert.NotNil(t, status.localized)
		assert.Len(t, status.localized.FieldViolations, 2)

		assert.Equal(t, "email", status.localized.FieldViolations[0].Field)
		assert.Equal(t, "password", status.localized.FieldViolations[1].Field)
	})

	t.Run("FieldLocalized Proto method", func(t *testing.T) {
		fieldLocalized := &FieldLocalized{
			Field: "email",
			Key:   "field.email.invalid",
			Args:  []any{"test@example.com", 25},
		}

		proto := fieldLocalized.Proto()

		assert.Equal(t, "email", proto.Field)
		assert.Equal(t, "field.email.invalid", proto.Key)
		assert.Len(t, proto.Args, 2)
	})

	t.Run("WithFieldsLocalized duplicate field validation", func(t *testing.T) {
		t.Run("duplicate within new fields", func(t *testing.T) {
			fields := []*FieldLocalized{
				{Field: "email", Key: "field.email.required"},
				{Field: "password", Key: "field.password.weak"},
				{Field: "email", Key: "field.email.invalid"}, // Duplicate field!
			}

			assert.Panics(t, func() {
				New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
					WithFieldsLocalized(fields...)
			}, "should panic when duplicate field names are provided")
		})

		t.Run("replacement preserves original position", func(t *testing.T) {
			// First add some fields
			status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithFieldLocalized("email", "field.email.required").
				WithFieldLocalized("password", "field.password.weak").
				WithFieldLocalized("age", "field.age.required")

			// Verify initial order
			assert.Len(t, status.localized.FieldViolations, 3)
			assert.Equal(t, "email", status.localized.FieldViolations[0].Field)
			assert.Equal(t, "password", status.localized.FieldViolations[1].Field)
			assert.Equal(t, "age", status.localized.FieldViolations[2].Field)

			// Replace middle field and add new field
			newFields := []*FieldLocalized{
				{Field: "phone", Key: "field.phone.invalid"},      // New field (should be appended)
				{Field: "password", Key: "field.password.strong"}, // Replace existing (should stay in position 1)
			}

			updatedStatus := status.WithFieldsLocalized(newFields...)

			// Should have 4 fields total
			assert.Len(t, updatedStatus.localized.FieldViolations, 4)

			// Verify order is preserved for existing fields, new fields appended
			fields := updatedStatus.localized.FieldViolations
			assert.Equal(t, "email", fields[0].Field)
			assert.Equal(t, "field.email.required", fields[0].Key)

			assert.Equal(t, "password", fields[1].Field)
			assert.Equal(t, "field.password.strong", fields[1].Key) // Replaced in place

			assert.Equal(t, "age", fields[2].Field)
			assert.Equal(t, "field.age.required", fields[2].Key)

			assert.Equal(t, "phone", fields[3].Field) // New field appended
			assert.Equal(t, "field.phone.invalid", fields[3].Key)
		})

		t.Run("WithFieldLocalized should replace existing field", func(t *testing.T) {
			status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithFieldLocalized("email", "field.email.required")

			// Replace the existing field
			updatedStatus := status.WithFieldLocalized("email", "field.email.format_error")

			assert.Len(t, updatedStatus.localized.FieldViolations, 1)
			assert.Equal(t, "email", updatedStatus.localized.FieldViolations[0].Field)
			assert.Equal(t, "field.email.format_error", updatedStatus.localized.FieldViolations[0].Key)
		})

		t.Run("duplicate within new fields should still panic", func(t *testing.T) {
			status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithFieldLocalized("email", "field.email.required")

			newFields := []*FieldLocalized{
				{Field: "password", Key: "field.password.weak"},
				{Field: "phone", Key: "field.phone.invalid"},
				{Field: "password", Key: "field.password.short"}, // Duplicate in new fields
			}

			assert.Panics(t, func() {
				status.WithFieldsLocalized(newFields...)
			}, "should panic when duplicate field names exist within new fields")
		})

		t.Run("new fields maintain input order", func(t *testing.T) {
			status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithFieldLocalized("email", "field.email.required")

			// Add multiple new fields in specific order
			newFields := []*FieldLocalized{
				{Field: "phone", Key: "field.phone.invalid"},
				{Field: "address", Key: "field.address.required"},
				{Field: "age", Key: "field.age.invalid"},
				{Field: "country", Key: "field.country.required"},
			}

			updatedStatus := status.WithFieldsLocalized(newFields...)

			// Should have 5 fields: email (existing) + 4 new fields
			assert.Len(t, updatedStatus.localized.FieldViolations, 5)

			fields := updatedStatus.localized.FieldViolations
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
				Key:         "field.email.invalid_format",
				Args:        []any{"user@example", "email format"},
				Description: "Email address format is invalid",
				Reason:      "INVALID_EMAIL_FORMAT",
			}

			proto := fieldViolation.Proto()
			assert.Equal(t, "email", proto.Field)
			assert.Equal(t, "field.email.invalid_format", proto.Key)
			assert.Len(t, proto.Args, 2)
			assert.Equal(t, "Email address format is invalid", proto.Description)
			assert.Equal(t, "INVALID_EMAIL_FORMAT", proto.Reason)
		})

		t.Run("New WithFieldViolation API", func(t *testing.T) {
			status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithFieldViolation("email", "INVALID_EMAIL_FORMAT", "Email address format is invalid", "field.email.invalid_format", "user@example")

			assert.NotNil(t, status.localized)
			assert.Len(t, status.localized.FieldViolations, 1)

			fieldViol := status.localized.FieldViolations[0]
			assert.Equal(t, "email", fieldViol.Field)
			assert.Equal(t, "INVALID_EMAIL_FORMAT", fieldViol.Reason)
			assert.Equal(t, "Email address format is invalid", fieldViol.Description)
			assert.Equal(t, "field.email.invalid_format", fieldViol.Key)
			assert.Len(t, fieldViol.Args, 1)
		})

		t.Run("New WithFieldsViolations API", func(t *testing.T) {
			violations := []*FieldViolation{
				{
					Field:       "email",
					Reason:      "REQUIRED",
					Description: "Email is required",
					Key:         "field.email.required",
				},
				{
					Field:       "password",
					Reason:      "TOO_WEAK",
					Description: "Password is too weak",
					Key:         "field.password.weak",
					Args:        []any{"minLength", 8},
				},
			}

			status := New(codes.InvalidArgument, "VALIDATION_ERROR", "validation failed").
				WithFieldsViolations(violations...)

			assert.NotNil(t, status.localized)
			assert.Len(t, status.localized.FieldViolations, 2)

			// Check first violation
			emailViol := status.localized.FieldViolations[0]
			assert.Equal(t, "email", emailViol.Field)
			assert.Equal(t, "REQUIRED", emailViol.Reason)
			assert.Equal(t, "Email is required", emailViol.Description)

			// Check second violation
			passwordViol := status.localized.FieldViolations[1]
			assert.Equal(t, "password", passwordViol.Field)
			assert.Equal(t, "TOO_WEAK", passwordViol.Reason)
			assert.Equal(t, "Password is too weak", passwordViol.Description)
			assert.Len(t, passwordViol.Args, 2)
		})
	})
}

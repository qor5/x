package statusx

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/qor5/x/v3/i18nx"
	"github.com/qor5/x/v3/lox"
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
		s := New(codes.InvalidArgument, "INVALID_INPUT", "invalid input provided")
		require.NotNil(t, s)

		assert.Equal(t, codes.InvalidArgument, s.Code())
		assert.Equal(t, "INVALID_INPUT", s.Reason())
		assert.Equal(t, "invalid input provided", s.Message())

		// Verify that localized key is immediately fixed to reason
		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "INVALID_INPUT", localized.Key)
		assert.Empty(t, localized.Args)
	})

	t.Run("stacktrace behavior", func(t *testing.T) {
		// Test OK status has no stacktrace
		s := New(codes.OK, statusv1.ErrorReason_OK.String(), "success")
		require.NotNil(t, s)
		assert.Nil(t, s.cause)

		// Test non-OK status has stacktrace
		s = New(codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "error")
		require.NotNil(t, s)
		require.NotNil(t, s.cause)

		// Verify stacktrace exists
		stackErr := s.cause
		assert.Contains(t, fmt.Sprintf("%+v", stackErr), "github.com/qor5/x/v3/statusx.TestNew")
	})
}

func TestNewf(t *testing.T) {
	s := Newf(codes.NotFound, statusv1.ErrorReason_NOT_FOUND.String(), "resource %s not found", "user")
	require.NotNil(t, s)

	assert.Equal(t, codes.NotFound, s.Code())
	assert.Equal(t, "NOT_FOUND", s.Reason())
	assert.Equal(t, "resource user not found", s.Message())
}

func TestErr(t *testing.T) {
	err := Error(codes.PermissionDenied, statusv1.ErrorReason_PERMISSION_DENIED.String(), "permission denied")
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
	err := Errorf(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "invalid input provided for %s", "username")
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
	s := New(codes.InvalidArgument, "MISSING_FIELD", "required field is missing").WithMetadata(md)

	require.NotNil(t, s)
	assert.Equal(t, md, s.Metadata())
}

func TestWithLocalized(t *testing.T) {
	s := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "invalid input").
		WithLocalized("error.invalid_input", "username", "required")

	require.NotNil(t, s)
	localized := s.Localized()
	assert.NotNil(t, localized)
	assert.Equal(t, "error.invalid_input", localized.Key)

	// Verify args using proto.Equal
	require.Len(t, localized.Args, 2)
	assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("username"))), localized.Args[0]))
	assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("required"))), localized.Args[1]))
}

func TestWithLocalizedArgs(t *testing.T) {
	t.Run("uses reason as key when no existing key", func(t *testing.T) {
		s := New(codes.InvalidArgument, "INVALID_EMAIL_FORMAT", "Email validation failed").
			WithLocalizedArgs("user@invalid.com")

		require.NotNil(t, s)
		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "INVALID_EMAIL_FORMAT", localized.Key) // Uses reason as key immediately
		require.Len(t, localized.Args, 1)
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("user@invalid.com"))), localized.Args[0]))
	})

	t.Run("preserves existing localized key", func(t *testing.T) {
		s := New(codes.PermissionDenied, "ACCESS_DENIED", "Permission check failed").
			WithLocalized("INSUFFICIENT_CREDITS", "placeholder"). // Set custom key first
			WithLocalizedArgs(100, 200)                           // Then add args

		require.NotNil(t, s)
		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "INSUFFICIENT_CREDITS", localized.Key) // Preserves existing key
		require.Len(t, localized.Args, 2)
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.Int64(100))), localized.Args[0]))
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.Int64(200))), localized.Args[1]))
	})

	t.Run("overwrites existing args", func(t *testing.T) {
		s := New(codes.NotFound, "USER_NOT_FOUND", "User lookup failed").
			WithLocalizedArgs("john", "initial"). // First set of args
			WithLocalizedArgs("jane", "updated")  // Second set overwrites

		require.NotNil(t, s)
		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "USER_NOT_FOUND", localized.Key) // Uses reason as key immediately
		require.Len(t, localized.Args, 2)
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("jane"))), localized.Args[0])) // Updated args
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("updated"))), localized.Args[1]))
	})

	t.Run("handles empty args", func(t *testing.T) {
		s := New(codes.Internal, "DATABASE_ERROR", "Connection failed").
			WithLocalizedArgs() // No args

		require.NotNil(t, s)
		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "DATABASE_ERROR", localized.Key) // Uses reason as key immediately
		assert.Empty(t, localized.Args)
	})

	t.Run("handles mixed types", func(t *testing.T) {
		s := New(codes.InvalidArgument, "VALIDATION_FAILED", "Field validation failed").
			WithLocalizedArgs("email", int64(5), true)

		require.NotNil(t, s)
		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "VALIDATION_FAILED", localized.Key) // Uses reason as key immediately
		require.Len(t, localized.Args, 3)
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("email"))), localized.Args[0]))
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.Int64(5))), localized.Args[1]))
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.Bool(true))), localized.Args[2]))
	})

	t.Run("preserves other status properties", func(t *testing.T) {
		metadata := map[string]string{"field": "email", "validator": "regex"}
		s := New(codes.InvalidArgument, "INVALID_EMAIL", "Email format check failed").
			WithMetadata(metadata).
			WithLocalizedArgs("test@invalid")

		require.NotNil(t, s)
		assert.Equal(t, codes.InvalidArgument, s.Code())
		assert.Equal(t, "INVALID_EMAIL", s.Reason())
		assert.Equal(t, "Email format check failed", s.Message())
		assert.Equal(t, metadata, s.Metadata())

		// Localized should be set correctly
		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "INVALID_EMAIL", localized.Key) // Uses reason as key immediately
		require.Len(t, localized.Args, 1)
		assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("test@invalid"))), localized.Args[0]))
	})

	t.Run("key is fixed at creation time", func(t *testing.T) {
		// Test that WithReason after New doesn't affect the already fixed key
		s := New(codes.InvalidArgument, "CREATION_REASON", "Initial message").
			WithReason("CHANGED_REASON") // Changes reason but key should remain fixed

		require.NotNil(t, s)
		assert.Equal(t, "CHANGED_REASON", s.Reason()) // Reason changed

		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "CREATION_REASON", localized.Key) // Key remains fixed to creation-time reason
		assert.Empty(t, localized.Args)
	})
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
		wrapped := Wrap(nil, codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "internal server error")
		assert.NotNil(t, wrapped)
		assert.Equal(t, codes.OK, wrapped.Code())
		assert.Equal(t, statusv1.ErrorReason_OK.String(), wrapped.Reason())
		assert.Equal(t, "", wrapped.Message())
	}

	{
		wrapped := Wrap(originalErr, codes.OK, statusv1.ErrorReason_OK.String(), "success")
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
		status := New(codes.NotFound, statusv1.ErrorReason_NOT_FOUND.String(), "resource not found")
		wrapped := Wrap(status.Err(), codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "internal server error")
		assert.Equal(t, wrapped, status)
		assert.Equal(t, codes.NotFound, wrapped.Code())
		assert.Equal(t, "NOT_FOUND", wrapped.Reason())
		assert.Equal(t, "resource not found", wrapped.Message())
	}

	{
		status, _ := status.New(codes.NotFound, "resource not found").WithDetails(&errdetails.ErrorInfo{
			Reason: "NOT_FOUND",
		})
		wrapped := Wrap(status.Err(), codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "internal server error")
		assert.Equal(t, codes.NotFound, wrapped.Code())
		assert.Equal(t, "NOT_FOUND", wrapped.Reason())
		assert.Equal(t, "resource not found", wrapped.Message())
	}

	// Test localized key consistency - this test would fail without the s.localized assignment
	t.Run("localized key should match reason after wrap", func(t *testing.T) {
		originalErr := errors.New("database connection failed")
		wrapped := Wrap(originalErr, codes.Internal, "DATABASE_CONNECTION_ERROR", "failed to connect to database")

		require.NotNil(t, wrapped)
		assert.Equal(t, "DATABASE_CONNECTION_ERROR", wrapped.Reason())

		// This is the critical test - localized key should match the reason
		localized := wrapped.Localized()
		require.NotNil(t, localized, "localized should not be nil")
		assert.Equal(t, "DATABASE_CONNECTION_ERROR", localized.Key,
			"localized key should match the reason for consistent i18n translation")
	})
}

func TestGRPCStatus(t *testing.T) {
	s := New(codes.PermissionDenied, statusv1.ErrorReason_PERMISSION_DENIED.String(), "permission denied").
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
	assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("access"))), localized.Args[0]))
	assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("user"))), localized.Args[1]))

	{
		s := New(codes.OK, statusv1.ErrorReason_OK.String(), "ok")
		st := s.GRPCStatus()
		require.NotNil(t, st)
		assert.Empty(t, st.Details())
		assert.Equal(t, codes.OK, st.Code())
		assert.Empty(t, st.Message())
	}
}

func TestFromError(t *testing.T) {
	s := New(codes.Unauthenticated, statusv1.ErrorReason_UNAUTHENTICATED.String(), "authentication required").
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

	err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "invalid input provided")
	assert.Equal(t, codes.InvalidArgument, Code(err.Err()))

	err = New(codes.OK, statusv1.ErrorReason_OK.String(), "success").WithCause(errors.New("original error"))
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

func TestWithLocalizedWellKnownTypes(t *testing.T) {
	now := time.Now()
	duration := 5 * time.Minute
	structData := map[string]any{
		"user": "john",
		"age":  30,
	}

	s := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
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
	assert.True(t, proto.Equal(lox.Must1(anypb.New(wrapperspb.String("string_arg"))), args[0]))
	// Note: We could add more specific extraction helpers for other types,
	// but the main point is that they're stored without errors
}

func TestWithLocalizedUnsupportedType(t *testing.T) {
	// Test that unsupported types cause panic
	assert.Panics(t, func() {
		New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").WithLocalized("test.key", make(chan int)) // channels are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").WithLocalized("test.key", func() {}) // functions are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").WithLocalized("test.key", []string{"a", "b"}) // arrays are not supported
	})

	assert.Panics(t, func() {
		New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").WithLocalized("test.key", []any{"a", "b"}) // slices are not supported
	})
}

func TestWithDetailsTypeValidation(t *testing.T) {
	// Valid usage should work
	t.Run("valid details", func(t *testing.T) {
		err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").
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
			New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").
				WithDetails(&statusv1.Localized{
					Key: "test.key",
				})
		}, "should panic when adding *statusv1.Localized")
	})

	t.Run("panic on errdetails.ErrorInfo", func(t *testing.T) {
		assert.Panics(t, func() {
			New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").
				WithDetails(&errdetails.ErrorInfo{
					Reason: "TEST_REASON",
				})
		}, "should panic when adding *errdetails.ErrorInfo")
	})
}

func TestWithFieldViolationsAPI(t *testing.T) {
	t.Run("single field via WithFieldViolations", func(t *testing.T) {
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
			WithFieldViolations(NewFieldViolation("email", "field.email.required", "Email is required"))

		assert.NotNil(t, status.badRequest)
		assert.Len(t, status.badRequest.FieldViolations, 1)

		field := status.badRequest.FieldViolations[0]
		assert.Equal(t, "email", field.Field)
		assert.Equal(t, "field.email.required", field.Reason)
	})

	t.Run("WithFieldViolations with args data", func(t *testing.T) {
		data := map[string]any{"minLength": 8}
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
			WithFieldViolations(NewFieldViolation("password", "field.password.min_length", "Password is too weak").WithLocalizedArgs(data))

		field := status.badRequest.FieldViolations[0]
		assert.Equal(t, "password", field.Field)
		assert.Equal(t, "field.password.min_length", field.Reason)
	})

	t.Run("batch WithFieldViolations", func(t *testing.T) {
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "multiple field errors").
			WithFieldViolations(
				NewFieldViolation("email", "field.email.required", "Email is required"),
				NewFieldViolation("password", "field.password.weak", "Password is weak").WithLocalizedArgs(map[string]any{"strength": "low"}),
				NewFieldViolation("age", "field.age.invalid_range", "Age is in invalid range").WithLocalizedArgs(18, 65),
			)

		assert.NotNil(t, status.badRequest)
		assert.Len(t, status.badRequest.FieldViolations, 3)
		assert.Equal(t, "email", status.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "password", status.badRequest.FieldViolations[1].Field)
		assert.Equal(t, "age", status.badRequest.FieldViolations[2].Field)
	})

	t.Run("combine WithLocalized and WithFieldViolations", func(t *testing.T) {
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
			WithLocalized("global.validation_failed").
			WithFieldViolations(
				NewFieldViolation("email", "field.email.invalid", "Email is invalid"),
				NewFieldViolation("phone", "field.phone.invalid", "Phone is invalid"),
			)

		assert.NotNil(t, status.Localized())
		assert.Equal(t, "global.validation_failed", status.Localized().GetKey())
		assert.Len(t, status.badRequest.FieldViolations, 2)
	})

	t.Run("multiple WithFieldViolations calls", func(t *testing.T) {
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
			WithFieldViolations(NewFieldViolation("email", "field.email.invalid", "Email is invalid")).
			WithFieldViolations(NewFieldViolation("password", "field.password.weak", "Password is weak"))

		assert.NotNil(t, status.badRequest)
		assert.Len(t, status.badRequest.FieldViolations, 2)

		assert.Equal(t, "email", status.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "password", status.badRequest.FieldViolations[1].Field)
	})

	t.Run("FieldViolation Proto mapping", func(t *testing.T) {
		fv := NewFieldViolation("email", "field.email.invalid", "Email is invalid").WithLocalizedArgs("test@example.com", 25)
		proto := fv.Proto()
		assert.Equal(t, "email", proto.Field)
		assert.Equal(t, "field.email.invalid", proto.Localized.Key)
		assert.Len(t, proto.Localized.Args, 2)
	})

	t.Run("WithFieldsLocalized append behavior", func(t *testing.T) {
		t.Run("duplicate fields are allowed and appended", func(t *testing.T) {
			// Multiple violations for the same field are now allowed
			status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithFieldViolations(
					NewFieldViolation("email", "REQUIRED", "Email is required"),
					NewFieldViolation("password", "WEAK", "Password is weak"),
					NewFieldViolation("email", "INVALID", "Email is invalid"),
				)

			// Should have all 3 violations
			assert.Len(t, status.badRequest.FieldViolations, 3)
			assert.Equal(t, "email", status.badRequest.FieldViolations[0].Field)
			assert.Equal(t, "REQUIRED", status.badRequest.FieldViolations[0].Reason)
			assert.Equal(t, "password", status.badRequest.FieldViolations[1].Field)
			assert.Equal(t, "WEAK", status.badRequest.FieldViolations[1].Reason)
			assert.Equal(t, "email", status.badRequest.FieldViolations[2].Field)
			assert.Equal(t, "INVALID", status.badRequest.FieldViolations[2].Reason)
		})

		t.Run("cross-call violations are appended in order", func(t *testing.T) {
			// First add some fields
			status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithFieldViolations(
					NewFieldViolation("email", "field.email.required", "Email is required"),
					NewFieldViolation("password", "field.password.weak", "Password is weak"),
					NewFieldViolation("age", "field.age.required", "Age is required"),
				)

				// Verify initial order
			assert.Len(t, status.badRequest.FieldViolations, 3)
			assert.Equal(t, "email", status.badRequest.FieldViolations[0].Field)
			assert.Equal(t, "password", status.badRequest.FieldViolations[1].Field)
			assert.Equal(t, "age", status.badRequest.FieldViolations[2].Field)

			// Add more fields including duplicate
			updatedStatus := status.WithFieldViolations(
				NewFieldViolation("phone", "field.phone.invalid", "Phone is invalid"),
				NewFieldViolation("password", "field.password.strong", "Password is strong"),
			)

			// Should have 5 fields total (original 3 + new 2)
			assert.Len(t, updatedStatus.badRequest.FieldViolations, 5)

			// Verify all violations are present in order
			fields := updatedStatus.badRequest.FieldViolations
			assert.Equal(t, "email", fields[0].Field)
			assert.Equal(t, "field.email.required", fields[0].Localized.Key)

			assert.Equal(t, "password", fields[1].Field)
			assert.Equal(t, "field.password.weak", fields[1].Localized.Key) // Original remains

			assert.Equal(t, "age", fields[2].Field)
			assert.Equal(t, "field.age.required", fields[2].Localized.Key)

			assert.Equal(t, "phone", fields[3].Field) // New field appended
			assert.Equal(t, "field.phone.invalid", fields[3].Localized.Key)

			assert.Equal(t, "password", fields[4].Field) // Duplicate password appended
			assert.Equal(t, "field.password.strong", fields[4].Localized.Key)
		})

		t.Run("WithFieldViolations appends rather than replaces", func(t *testing.T) {
			status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithFieldViolations(NewFieldViolation("email", "field.email.required", "Email is required"))

			// Add another violation for the same field
			updatedStatus := status.WithFieldViolations(NewFieldViolation("email", "field.email.format_error", "Email format error"))

			// Should have 2 violations total
			assert.Len(t, updatedStatus.badRequest.FieldViolations, 2)
			assert.Equal(t, "email", updatedStatus.badRequest.FieldViolations[0].Field)
			assert.Equal(t, "field.email.required", updatedStatus.badRequest.FieldViolations[0].Localized.Key)
			assert.Equal(t, "email", updatedStatus.badRequest.FieldViolations[1].Field)
			assert.Equal(t, "field.email.format_error", updatedStatus.badRequest.FieldViolations[1].Localized.Key)
		})

		t.Run("multiple same-field violations in single call", func(t *testing.T) {
			status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithFieldViolations(NewFieldViolation("email", "field.email.required", "Email is required"))

			// Add multiple violations for same field in single call
			updatedStatus := status.WithFieldViolations(
				NewFieldViolation("password", "WEAK", "Password is weak"),
				NewFieldViolation("phone", "INVALID", "Phone is invalid"),
				NewFieldViolation("password", "SHORT", "Password is short"),
			)

			// Should have 4 violations total (1 existing + 3 new)
			assert.Len(t, updatedStatus.badRequest.FieldViolations, 4)
			assert.Equal(t, "email", updatedStatus.badRequest.FieldViolations[0].Field)
			assert.Equal(t, "password", updatedStatus.badRequest.FieldViolations[1].Field)
			assert.Equal(t, "WEAK", updatedStatus.badRequest.FieldViolations[1].Reason)
			assert.Equal(t, "phone", updatedStatus.badRequest.FieldViolations[2].Field)
			assert.Equal(t, "password", updatedStatus.badRequest.FieldViolations[3].Field)
			assert.Equal(t, "SHORT", updatedStatus.badRequest.FieldViolations[3].Reason)
		})

		t.Run("new fields maintain input order", func(t *testing.T) {
			status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithFieldViolations(NewFieldViolation("email", "field.email.required", "Email is required"))

			// Add multiple new fields in specific order
			updatedStatus := status.WithFieldViolations(
				NewFieldViolation("phone", "field.phone.invalid", "Phone is invalid"),
				NewFieldViolation("address", "field.address.required", "Address is required"),
				NewFieldViolation("age", "field.age.invalid", "Age is invalid"),
				NewFieldViolation("country", "field.country.required", "Country is required"),
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
			fieldViolation := NewFieldViolation("email", "field.email.invalid_format", "Email address format is invalid").
				WithLocalizedArgs("user@example", "email format")

			proto := fieldViolation.Proto()
			assert.Equal(t, "email", proto.Field)
			assert.Equal(t, "field.email.invalid_format", proto.Localized.Key)
			assert.Len(t, proto.Localized.Args, 2)
			assert.Equal(t, "Email address format is invalid", proto.Description)
			assert.Equal(t, "field.email.invalid_format", proto.Reason)
		})

		t.Run("New WithFieldViolation API", func(t *testing.T) {
			status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithFieldViolations(NewFieldViolation("email", "field.email.invalid_format", "Email address format is invalid").WithLocalizedArgs("user@example"))

			assert.NotNil(t, status.badRequest)
			assert.Len(t, status.badRequest.FieldViolations, 1)

			fieldViol := status.badRequest.FieldViolations[0]
			assert.Equal(t, "email", fieldViol.Field)
			assert.Equal(t, "field.email.invalid_format", fieldViol.Reason)
			assert.Equal(t, "Email address format is invalid", fieldViol.Description)
			assert.Equal(t, "field.email.invalid_format", fieldViol.Localized.Key)
			assert.Len(t, fieldViol.Localized.Args, 1)
		})

		t.Run("New WithFieldViolations API", func(t *testing.T) {
			violations := []*FieldViolation{
				NewFieldViolation("email", "field.email.required", "Email is required"),
				NewFieldViolation("password", "field.password.weak", "Password is too weak").
					WithLocalizedArgs("minLength", 8),
			}

			status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
				WithFieldViolations(violations...)

			assert.NotNil(t, status.badRequest)
			assert.Len(t, status.badRequest.FieldViolations, 2)

			// Check first violation
			emailViol := status.badRequest.FieldViolations[0]
			assert.Equal(t, "email", emailViol.Field)
			assert.Equal(t, "field.email.required", emailViol.Reason)
			assert.Equal(t, "Email is required", emailViol.Description)

			// Check second violation
			passwordViol := status.badRequest.FieldViolations[1]
			assert.Equal(t, "password", passwordViol.Field)
			assert.Equal(t, "field.password.weak", passwordViol.Reason)
			assert.Equal(t, "Password is too weak", passwordViol.Description)
			assert.Len(t, passwordViol.Localized.Args, 2)
		})
	})
}

func TestWithFlattenFieldViolations(t *testing.T) {
	t.Run("mixed types in single call", func(t *testing.T) {
		// Create various inputs
		single := NewFieldViolation("email", "INVALID_FORMAT", "Email format is invalid")
		slice := []*FieldViolation{
			NewFieldViolation("name", "REQUIRED", "Name is required"),
			NewFieldViolation("age", "TOO_YOUNG", "Age is too young"),
		}

		// Create another status with violations
		otherStatus := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "nested validation").
			WithFieldViolations(NewFieldViolation("nested.field", "INVALID", "Nested field is invalid"))

		// Use WithFlattenFieldViolations
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
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
		statusErr := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
			WithFieldViolations(
				NewFieldViolation("field1", "INVALID", "Field1 is invalid"),
				NewFieldViolation("field2", "REQUIRED", "Field2 is required"),
			).Err()

		// Use error via ToFieldViolations in WithFlattenFieldViolations
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "main validation error").
			WithFlattenFieldViolations(ToFieldViolations(statusErr, "nested"), NewFieldViolation("direct", "INVALID", "Direct field is invalid"))

		// Verify results
		assert.NotNil(t, status.badRequest)
		// Should have: 1 main error + 2 nested errors + 1 direct = 4 total
		assert.Len(t, status.badRequest.FieldViolations, 4)
	})

	t.Run("delegation to existing method", func(t *testing.T) {
		// Test that the new method delegates properly to WithFieldViolations
		// and maintains the same append behavior for duplicate field handling

		existing := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "base").
			WithFieldViolations(NewFieldViolation("email", "ORIGINAL", "Original email error"))

		// Add same field via WithFlattenFieldViolations - should append
		updated := existing.WithFlattenFieldViolations(NewFieldViolation("email", "UPDATED", "Updated email error"))

		// Verify delegation behavior - should have both violations
		assert.Len(t, updated.badRequest.FieldViolations, 2)
		assert.Equal(t, "email", updated.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "ORIGINAL", updated.badRequest.FieldViolations[0].Reason)
		assert.Equal(t, "email", updated.badRequest.FieldViolations[1].Field)
		assert.Equal(t, "UPDATED", updated.badRequest.FieldViolations[1].Reason)
	})

	t.Run("nil input handling", func(t *testing.T) {
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").
			WithFlattenFieldViolations(nil, (*FieldViolation)(nil), []*FieldViolation(nil))

		// Should handle nil inputs gracefully
		if status.badRequest != nil {
			assert.Len(t, status.badRequest.FieldViolations, 0)
		}
	})
}

func TestWithFieldViolationsAppendStrategy(t *testing.T) {
	t.Run("allows multiple violations for same field", func(t *testing.T) {
		// Multiple violations for same field within same call are now allowed
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").WithFieldViolations(
			NewFieldViolation("email", "REQUIRED", "Email is required"),
			NewFieldViolation("email", "INVALID", "Email format is invalid"),
		)

		// Should have both violations for the same field
		assert.Len(t, status.badRequest.FieldViolations, 2)
		assert.Equal(t, "email", status.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "REQUIRED", status.badRequest.FieldViolations[0].Reason)
		assert.Equal(t, "email", status.badRequest.FieldViolations[1].Field)
		assert.Equal(t, "INVALID", status.badRequest.FieldViolations[1].Reason)

		// Cross-call duplicates are also appended
		updated := status.WithFieldViolations(
			NewFieldViolation("email", "TOO_LONG", "Email is too long"),
		)

		// Should have 3 violations total, all for email field
		assert.Len(t, updated.badRequest.FieldViolations, 3)
		assert.Equal(t, "email", updated.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "REQUIRED", updated.badRequest.FieldViolations[0].Reason)
		assert.Equal(t, "email", updated.badRequest.FieldViolations[1].Field)
		assert.Equal(t, "INVALID", updated.badRequest.FieldViolations[1].Reason)
		assert.Equal(t, "email", updated.badRequest.FieldViolations[2].Field)
		assert.Equal(t, "TOO_LONG", updated.badRequest.FieldViolations[2].Reason)
	})

	t.Run("append strategy preserves order and allows duplicates", func(t *testing.T) {
		// All violations are appended in order without deduplication
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").WithFieldViolations(
			NewFieldViolation("field1", "ERROR1", "Field1 error"),
			NewFieldViolation("field2", "ERROR2", "Field2 error"),
			NewFieldViolation("field3", "ERROR3", "Field3 error"),
		)

		// Add another violation for field2
		updated := status.WithFieldViolations(
			NewFieldViolation("field2", "ADDITIONAL_ERROR", "Additional field2 error"),
		)

		// Verify all violations are present, with the new one appended at the end
		assert.Len(t, updated.badRequest.FieldViolations, 4)
		assert.Equal(t, "field1", updated.badRequest.FieldViolations[0].Field)
		assert.Equal(t, "ERROR1", updated.badRequest.FieldViolations[0].Reason)
		assert.Equal(t, "field2", updated.badRequest.FieldViolations[1].Field)
		assert.Equal(t, "ERROR2", updated.badRequest.FieldViolations[1].Reason)
		assert.Equal(t, "field3", updated.badRequest.FieldViolations[2].Field)
		assert.Equal(t, "ERROR3", updated.badRequest.FieldViolations[2].Reason)
		assert.Equal(t, "field2", updated.badRequest.FieldViolations[3].Field)
		assert.Equal(t, "ADDITIONAL_ERROR", updated.badRequest.FieldViolations[3].Reason)
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
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "original message").
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
		status := New(codes.InvalidArgument, "VALIDATION_FAILED", "original message")

		translated := status.Translated(ib, language.MustParse("zh-CN"))

		details := translated.Details()
		localizedMsg := ExtractDetail[*errdetails.LocalizedMessage](details)
		assert.NotNil(t, localizedMsg)
		assert.Equal(t, "zh-CN", localizedMsg.Locale)
		assert.Equal(t, "验证失败", localizedMsg.Message)
	})

	t.Run("translate field violations with localized templates", func(t *testing.T) {
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
			WithFieldViolations(
				NewFieldViolation("email", "INVALID_FORMAT", "Email format is invalid").
					WithLocalized("email.invalid_format"),
				NewFieldViolation("password", "TOO_SHORT", "Password is too short").
					WithLocalized("field.password.min_length", 8),
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
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "original")

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
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
			WithLocalized("invalid_request").
			WithFieldViolations(
				NewFieldViolation("email", "REQUIRED", "Email is required").
					WithLocalized("email.invalid_format"),
			)

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

func TestWithMessage(t *testing.T) {
	t.Run("basic message replacement", func(t *testing.T) {
		originalStatus := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "original message")
		newStatus := originalStatus.WithMessage("new message")

		// Original status should remain unchanged
		assert.Equal(t, "original message", originalStatus.Message())

		// New status should have updated message
		assert.Equal(t, "new message", newStatus.Message())

		// Other fields should remain the same
		assert.Equal(t, codes.InvalidArgument, newStatus.Code())
		assert.Equal(t, "INVALID_ARGUMENT", newStatus.Reason())
	})

	t.Run("empty message", func(t *testing.T) {
		originalStatus := New(codes.NotFound, statusv1.ErrorReason_NOT_FOUND.String(), "resource not found")
		newStatus := originalStatus.WithMessage("")

		assert.Equal(t, "", newStatus.Message())
		assert.Equal(t, codes.NotFound, newStatus.Code())
	})

	t.Run("preserve other properties", func(t *testing.T) {
		originalStatus := New(codes.PermissionDenied, "CUSTOM_REASON", "permission denied").
			WithMetadata(map[string]string{"key": "value"})

		newStatus := originalStatus.WithMessage("access denied")

		assert.Equal(t, "access denied", newStatus.Message())
		assert.Equal(t, codes.PermissionDenied, newStatus.Code())
		assert.Equal(t, "CUSTOM_REASON", newStatus.Reason())
		assert.Equal(t, map[string]string{"key": "value"}, newStatus.Metadata())
	})
}

func TestWithMessagef(t *testing.T) {
	t.Run("basic formatted message", func(t *testing.T) {
		originalStatus := New(codes.NotFound, statusv1.ErrorReason_NOT_FOUND.String(), "original message")
		newStatus := originalStatus.WithMessagef("user %s not found", "john")

		// Original status should remain unchanged
		assert.Equal(t, "original message", originalStatus.Message())

		// New status should have formatted message
		assert.Equal(t, "user john not found", newStatus.Message())
		assert.Equal(t, codes.NotFound, newStatus.Code())
	})

	t.Run("multiple format arguments", func(t *testing.T) {
		originalStatus := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "original")
		newStatus := originalStatus.WithMessagef("validation failed: field %s requires %d characters", "email", 5)

		assert.Equal(t, "validation failed: field email requires 5 characters", newStatus.Message())
	})

	t.Run("no format arguments", func(t *testing.T) {
		originalStatus := New(codes.Internal, statusv1.ErrorReason_INTERNAL.String(), "original")
		newStatus := originalStatus.WithMessagef("internal error occurred")

		assert.Equal(t, "internal error occurred", newStatus.Message())
	})

	t.Run("preserve other properties", func(t *testing.T) {
		originalStatus := New(codes.Unauthenticated, "TOKEN_EXPIRED", "auth failed").
			WithMetadata(map[string]string{"token": "expired"})

		newStatus := originalStatus.WithMessagef("authentication failed: %s", "token expired")

		assert.Equal(t, "authentication failed: token expired", newStatus.Message())
		assert.Equal(t, codes.Unauthenticated, newStatus.Code())
		assert.Equal(t, "TOKEN_EXPIRED", newStatus.Reason())
		assert.Equal(t, map[string]string{"token": "expired"}, newStatus.Metadata())
	})
}

func TestStatusMethods(t *testing.T) {
	t.Run("Details method", func(t *testing.T) {
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test")

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
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test")
		assert.Nil(t, status.BadRequest())

		// Status with BadRequest
		status = status.WithFieldViolations(
			NewFieldViolation("email", "REQUIRED", "Email is required"),
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
		status := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "Validation failed").
			WithFieldViolations(
				NewFieldViolation("name", "REQUIRED", "Name required"),
			)
		violations := status.ToFieldViolations("user")

		assert.Len(t, violations, 2) // Main + nested
		assert.Equal(t, "user", violations[0].Field())
		assert.Equal(t, "user.name", violations[1].Field())
	})

	t.Run("WithCode method", func(t *testing.T) {
		originalStatus := New(codes.InvalidArgument, "TEST", "test message")
		newStatus := originalStatus.WithCode(codes.NotFound)

		// Original unchanged
		assert.Equal(t, codes.InvalidArgument, originalStatus.Code())

		// New status has updated code
		assert.Equal(t, codes.NotFound, newStatus.Code())
		assert.Equal(t, "TEST", newStatus.Reason())          // Reason unchanged
		assert.Equal(t, "test message", newStatus.Message()) // Message unchanged
	})

	t.Run("WithReason method", func(t *testing.T) {
		originalStatus := New(codes.InvalidArgument, "OLD_REASON", "test message")
		newStatus := originalStatus.WithReason("NEW_REASON")

		// Original unchanged
		assert.Equal(t, "OLD_REASON", originalStatus.Reason())

		// New status has updated reason
		assert.Equal(t, "NEW_REASON", newStatus.Reason())
		assert.Equal(t, codes.InvalidArgument, newStatus.Code()) // Code unchanged
		assert.Equal(t, "test message", newStatus.Message())     // Message unchanged
	})

	t.Run("Status method for StatusError", func(t *testing.T) {
		status := New(codes.NotFound, statusv1.ErrorReason_NOT_FOUND.String(), "Resource not found")
		err := status.Err().(*StatusError)

		grpcStatus := err.Status()
		assert.Equal(t, codes.NotFound, grpcStatus.Code())
		assert.Equal(t, "Resource not found", grpcStatus.Message())
	})
}

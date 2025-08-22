package statusx

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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
	assert.Equal(t, []string{"username", "required"}, localized.Args)
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
	require.Equal(t, []string{"access", "user"}, localizedMessage.Args)

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
		assert.Len(t, s.extraDetails, 1)
		assert.True(t, proto.Equal(s.extraDetails[0], &errdetails.LocalizedMessage{
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
		assert.Contains(t, formatted, "pkg/statusx.TestErrorFormat")
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

package httperrors

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("basic properties", func(t *testing.T) {
		s := New(http.StatusBadRequest, "INVALID_INPUT", "invalid input provided")
		require.NotNil(t, s)

		assert.Equal(t, http.StatusBadRequest, s.StatusCode())
		assert.Equal(t, "INVALID_INPUT", s.Reason())
		assert.Equal(t, "invalid input provided", s.Message())

		localized := s.Localized()
		assert.NotNil(t, localized)
		assert.Equal(t, "INVALID_INPUT", localized.Key())
		assert.Empty(t, localized.Args())
	})

	t.Run("stacktrace for non-2xx", func(t *testing.T) {
		s := New(http.StatusInternalServerError, ReasonInternal, "error")
		require.NotNil(t, s)
		require.NotNil(t, s.cause)
		assert.Contains(t, fmt.Sprintf("%+v", s.cause), "github.com/qor5/x/v3/httperrors.TestNew")
	})

	t.Run("no stacktrace for 2xx", func(t *testing.T) {
		s := New(http.StatusOK, ReasonOK, "success")
		require.NotNil(t, s)
		assert.Nil(t, s.cause)
	})

	t.Run("panic on empty reason", func(t *testing.T) {
		assert.Panics(t, func() {
			New(http.StatusBadRequest, "", "msg")
		})
	})
}

func TestNewf(t *testing.T) {
	s := Newf(http.StatusNotFound, ReasonNotFound, "resource %s not found", "user")
	require.NotNil(t, s)

	assert.Equal(t, http.StatusNotFound, s.StatusCode())
	assert.Equal(t, "NOT_FOUND", s.Reason())
	assert.Equal(t, "resource user not found", s.Message())
}

func TestNewStatus(t *testing.T) {
	s := NewStatus(http.StatusNotFound, "user not found")
	require.NotNil(t, s)

	assert.Equal(t, http.StatusNotFound, s.StatusCode())
	assert.Equal(t, ReasonNotFound, s.Reason())
	assert.Equal(t, "user not found", s.Message())
}

func TestNewStatusf(t *testing.T) {
	s := NewStatusf(http.StatusNotFound, "resource %s not found", "order")
	require.NotNil(t, s)

	assert.Equal(t, http.StatusNotFound, s.StatusCode())
	assert.Equal(t, ReasonNotFound, s.Reason())
	assert.Equal(t, "resource order not found", s.Message())
}

func TestError(t *testing.T) {
	err := Error(http.StatusForbidden, ReasonPermissionDenied, "permission denied")
	require.NotNil(t, err)

	s, ok := FromError(err)
	require.True(t, ok)
	require.NotNil(t, s)

	assert.Equal(t, http.StatusForbidden, s.StatusCode())
	assert.Equal(t, "PERMISSION_DENIED", s.Reason())
	assert.Equal(t, "permission denied", s.Message())

	require.NotNil(t, s.Err())
	assert.True(t, errors.Is(s.Err(), err))
}

func TestErrorf(t *testing.T) {
	err := Errorf(http.StatusBadRequest, ReasonInvalidArgument, "invalid input for %s", "username")
	require.NotNil(t, err)

	s, ok := FromError(err)
	require.True(t, ok)
	assert.Equal(t, http.StatusBadRequest, s.StatusCode())
	assert.Equal(t, "INVALID_ARGUMENT", s.Reason())
	assert.Equal(t, "invalid input for username", s.Message())
}

func TestWithMetadata(t *testing.T) {
	md := map[string]string{"field": "username", "error": "missing"}
	s := New(http.StatusBadRequest, "MISSING_FIELD", "required field is missing").WithMetadata(md)

	require.NotNil(t, s)
	assert.Equal(t, md, s.Metadata())

	// Verify immutability
	md["new_key"] = "new_value"
	assert.NotEqual(t, md, s.Metadata())
}

func TestWithLocalized(t *testing.T) {
	s := New(http.StatusNotFound, "NOT_FOUND", "not found").
		WithLocalized("custom.key", "arg1", "arg2")

	localized := s.Localized()
	require.NotNil(t, localized)
	assert.Equal(t, "custom.key", localized.Key())
	assert.Equal(t, []any{"arg1", "arg2"}, localized.Args())
}

func TestWithLocalizedArgs(t *testing.T) {
	s := New(http.StatusNotFound, "NOT_FOUND", "not found").
		WithLocalizedArgs("arg1", "arg2")

	localized := s.Localized()
	require.NotNil(t, localized)
	assert.Equal(t, "NOT_FOUND", localized.Key())
	assert.Equal(t, []any{"arg1", "arg2"}, localized.Args())
}

func TestWithLocalized_PanicOnEmptyKey(t *testing.T) {
	assert.Panics(t, func() {
		New(http.StatusBadRequest, "TEST", "test").WithLocalized("")
	})
}

func TestWithCause(t *testing.T) {
	original := errors.New("original error")
	s := New(http.StatusInternalServerError, ReasonInternal, "internal error").WithCause(original)

	require.NotNil(t, s.Cause())
	assert.Contains(t, s.Cause().Error(), "original error")
}

func TestWithMessage(t *testing.T) {
	s := New(http.StatusBadRequest, "TEST", "original").WithMessage("updated")
	assert.Equal(t, "updated", s.Message())
}

func TestWithMessagef(t *testing.T) {
	s := New(http.StatusBadRequest, "TEST", "original").WithMessagef("updated %s", "msg")
	assert.Equal(t, "updated msg", s.Message())
}

func TestWithReason(t *testing.T) {
	s := New(http.StatusBadRequest, "OLD_REASON", "msg").WithReason("NEW_REASON")
	assert.Equal(t, "NEW_REASON", s.Reason())
}

func TestWithStatusCode(t *testing.T) {
	s := New(http.StatusBadRequest, "TEST", "msg").WithStatusCode(http.StatusNotFound)
	assert.Equal(t, http.StatusNotFound, s.StatusCode())
}

func TestWithFieldViolations(t *testing.T) {
	fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
	fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")

	s := New(http.StatusBadRequest, ReasonInvalidArgument, "invalid argument").
		WithFieldViolations(fv1, fv2)

	fvs := s.FieldViolations()
	require.Len(t, fvs, 2)
	assert.Equal(t, "email", fvs[0].Field())
	assert.Equal(t, "REQUIRED", fvs[0].Reason())
	assert.Equal(t, "name", fvs[1].Field())
	assert.Equal(t, "TOO_SHORT", fvs[1].Reason())
}

func TestWithFieldViolations_PanicOnEmptyField(t *testing.T) {
	assert.Panics(t, func() {
		fv := &FieldViolation{field: "", reason: "REQUIRED"}
		New(http.StatusBadRequest, "TEST", "msg").WithFieldViolations(fv)
	})
}

func TestWithFieldViolations_PanicOnEmptyReason(t *testing.T) {
	assert.Panics(t, func() {
		fv := &FieldViolation{field: "email", reason: ""}
		New(http.StatusBadRequest, "TEST", "msg").WithFieldViolations(fv)
	})
}

func TestWithFlattenFieldViolations(t *testing.T) {
	fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
	fvs := FieldViolations{NewFieldViolation("name", "TOO_SHORT", "name is too short")}

	s := New(http.StatusBadRequest, ReasonInvalidArgument, "invalid argument").
		WithFlattenFieldViolations(fv1, fvs)

	result := s.FieldViolations()
	require.Len(t, result, 2)
}

func TestErr(t *testing.T) {
	t.Run("non-2xx returns error", func(t *testing.T) {
		s := New(http.StatusBadRequest, "TEST", "msg")
		require.NotNil(t, s.Err())
	})

	t.Run("2xx returns nil", func(t *testing.T) {
		s := New(http.StatusOK, ReasonOK, "ok")
		assert.Nil(t, s.Err())
	})
}

func TestStatusCode_OKWithCause(t *testing.T) {
	s := &Status{
		httpStatus: http.StatusOK,
		reason:     ReasonOK,
		cause:      errors.New("unexpected"),
	}
	assert.Equal(t, http.StatusInternalServerError, s.StatusCode())
	assert.Equal(t, ReasonUnknown, s.Reason())
}

func TestNilStatus(t *testing.T) {
	var s *Status
	assert.Equal(t, http.StatusOK, s.StatusCode())
	assert.Equal(t, ReasonOK, s.Reason())
	assert.Equal(t, "", s.Message())
	assert.Nil(t, s.Cause())
	assert.Nil(t, s.Metadata())
	assert.Nil(t, s.Localized())
	assert.Nil(t, s.FieldViolations())
	assert.Nil(t, s.ToFieldViolations("test"))
}

func TestClone(t *testing.T) {
	original := New(http.StatusBadRequest, "TEST", "msg").
		WithMetadata(map[string]string{"k": "v"}).
		WithLocalized("key", "arg").
		WithFieldViolations(NewFieldViolation("f", "R", "d"))

	cloned := Clone(original)

	assert.Equal(t, original.StatusCode(), cloned.StatusCode())
	assert.Equal(t, original.Reason(), cloned.Reason())
	assert.Equal(t, original.Message(), cloned.Message())
	assert.Equal(t, original.Metadata(), cloned.Metadata())
	assert.Equal(t, original.Localized().Key(), cloned.Localized().Key())
	require.Len(t, cloned.FieldViolations(), 1)
	assert.Equal(t, "f", cloned.FieldViolations()[0].Field())

	assert.Nil(t, Clone(nil))
}

func TestClone_Immutability(t *testing.T) {
	original := New(http.StatusBadRequest, "TEST", "msg").
		WithMetadata(map[string]string{"k": "v"})

	cloned := Clone(original)
	cloned.metadata["k"] = "modified"

	assert.Equal(t, "v", original.Metadata()["k"])
}

func TestStatusError_Error(t *testing.T) {
	err := New(http.StatusNotFound, "NOT_FOUND", "user not found").Err()
	require.NotNil(t, err)
	assert.Contains(t, err.Error(), "http error")
	assert.Contains(t, err.Error(), "NOT_FOUND")
	assert.Contains(t, err.Error(), "user not found")
}

func TestStatusError_Unwrap(t *testing.T) {
	original := errors.New("root cause")
	err := New(http.StatusInternalServerError, ReasonInternal, "failed").WithCause(original).Err()

	var se *StatusError
	require.True(t, errors.As(err, &se))
	assert.NotNil(t, se.Unwrap())
}

func TestStatusError_Is(t *testing.T) {
	err1 := Error(http.StatusNotFound, "NOT_FOUND", "user not found")
	err2 := Error(http.StatusNotFound, "NOT_FOUND", "order not found")
	err3 := Error(http.StatusBadRequest, "NOT_FOUND", "not found")
	err4 := Error(http.StatusNotFound, "OTHER_REASON", "not found")

	// Same httpStatus + reason → Is returns true (even with different messages)
	assert.True(t, errors.Is(err1, err2))

	// Different httpStatus → Is returns false
	assert.False(t, errors.Is(err1, err3))

	// Different reason → Is returns false
	assert.False(t, errors.Is(err1, err4))
}

func TestStatusError_Format(t *testing.T) {
	err := New(http.StatusInternalServerError, ReasonInternal, "internal error").Err()

	// %s format
	assert.Contains(t, fmt.Sprintf("%s", err), "http error")

	// %v format
	assert.Contains(t, fmt.Sprintf("%v", err), "http error")

	// %+v format includes stacktrace
	assert.Contains(t, fmt.Sprintf("%+v", err), "httperrors.TestStatusError_Format")

	// %q format
	assert.Contains(t, fmt.Sprintf("%q", err), "http error")
}

func TestStatusError_Status(t *testing.T) {
	s := New(http.StatusBadRequest, "TEST", "msg")
	err := s.Err()
	var se *StatusError
	require.True(t, errors.As(err, &se))
	assert.Equal(t, s.StatusCode(), se.Status().StatusCode())
	assert.Equal(t, s.Reason(), se.Status().Reason())
}

func TestFromError(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		s, ok := FromError(nil)
		assert.True(t, ok)
		assert.Equal(t, http.StatusOK, s.StatusCode())
		assert.Equal(t, ReasonOK, s.Reason())
	})

	t.Run("StatusError", func(t *testing.T) {
		err := Error(http.StatusNotFound, "NOT_FOUND", "not found")
		s, ok := FromError(err)
		assert.True(t, ok)
		assert.Equal(t, http.StatusNotFound, s.StatusCode())
		assert.Equal(t, "NOT_FOUND", s.Reason())
	})

	t.Run("context.DeadlineExceeded", func(t *testing.T) {
		s, ok := FromError(context.DeadlineExceeded)
		assert.True(t, ok)
		assert.Equal(t, http.StatusGatewayTimeout, s.StatusCode())
		assert.Equal(t, ReasonDeadlineExceeded, s.Reason())
	})

	t.Run("context.Canceled", func(t *testing.T) {
		s, ok := FromError(context.Canceled)
		assert.True(t, ok)
		assert.Equal(t, 499, s.StatusCode())
		assert.Equal(t, ReasonCanceled, s.Reason())
	})

	t.Run("wrapped context.DeadlineExceeded", func(t *testing.T) {
		err := fmt.Errorf("wrapper: %w", context.DeadlineExceeded)
		s, ok := FromError(err)
		assert.True(t, ok)
		assert.Equal(t, http.StatusGatewayTimeout, s.StatusCode())
	})

	t.Run("wrapped context.Canceled", func(t *testing.T) {
		err := fmt.Errorf("wrapper: %w", context.Canceled)
		s, ok := FromError(err)
		assert.True(t, ok)
		assert.Equal(t, 499, s.StatusCode())
	})

	t.Run("unknown error", func(t *testing.T) {
		err := errors.New("some random error")
		s, ok := FromError(err)
		assert.False(t, ok)
		assert.Equal(t, http.StatusInternalServerError, s.StatusCode())
		assert.Equal(t, ReasonUnknown, s.Reason())
	})
}

func TestStatusCode_Func(t *testing.T) {
	err := Error(http.StatusNotFound, "NOT_FOUND", "not found")
	assert.Equal(t, http.StatusNotFound, StatusCode(err))
}

func TestReason_Func(t *testing.T) {
	err := Error(http.StatusNotFound, "NOT_FOUND", "not found")
	assert.Equal(t, "NOT_FOUND", Reason(err))
}

func TestConvert(t *testing.T) {
	err := Error(http.StatusNotFound, "NOT_FOUND", "not found")
	s := Convert(err)
	assert.Equal(t, http.StatusNotFound, s.StatusCode())
	assert.Equal(t, "NOT_FOUND", s.Reason())
}

func TestWrap(t *testing.T) {
	t.Run("wraps unknown error", func(t *testing.T) {
		original := errors.New("db connection failed")
		s := Wrap(original, http.StatusInternalServerError, ReasonInternal, "database error")

		assert.Equal(t, http.StatusInternalServerError, s.StatusCode())
		assert.Equal(t, ReasonInternal, s.Reason())
		assert.Equal(t, "database error", s.Message())
		assert.NotNil(t, s.Cause())
	})

	t.Run("preserves StatusError", func(t *testing.T) {
		original := Error(http.StatusNotFound, "NOT_FOUND", "user not found")
		s := Wrap(original, http.StatusInternalServerError, ReasonInternal, "should not change")

		assert.Equal(t, http.StatusNotFound, s.StatusCode())
		assert.Equal(t, "NOT_FOUND", s.Reason())
		assert.Equal(t, "user not found", s.Message())
	})
}

func TestWrapf(t *testing.T) {
	original := errors.New("timeout")
	s := Wrapf(original, http.StatusGatewayTimeout, ReasonDeadlineExceeded, "request to %s timed out", "api.example.com")

	assert.Equal(t, http.StatusGatewayTimeout, s.StatusCode())
	assert.Equal(t, "request to api.example.com timed out", s.Message())
}

func TestWrapStatus(t *testing.T) {
	original := errors.New("db error")
	s := WrapStatus(original, http.StatusInternalServerError, "database error")

	assert.Equal(t, http.StatusInternalServerError, s.StatusCode())
	assert.Equal(t, ReasonInternal, s.Reason())
}

func TestWrapStatusf(t *testing.T) {
	original := errors.New("db error")
	s := WrapStatusf(original, http.StatusInternalServerError, "failed to query %s", "users")

	assert.Equal(t, http.StatusInternalServerError, s.StatusCode())
	assert.Equal(t, ReasonInternal, s.Reason())
	assert.Equal(t, "failed to query users", s.Message())
}

func TestString(t *testing.T) {
	s := New(http.StatusNotFound, "NOT_FOUND", "user not found")
	str := s.String()
	assert.Contains(t, str, "404")
	assert.Contains(t, str, "NOT_FOUND")
	assert.Contains(t, str, "user not found")
}

func TestImmutability(t *testing.T) {
	original := New(http.StatusBadRequest, "ORIGINAL", "original message")

	modified := original.WithMessage("modified message")
	assert.Equal(t, "original message", original.Message())
	assert.Equal(t, "modified message", modified.Message())

	withReason := original.WithReason("NEW_REASON")
	assert.Equal(t, "ORIGINAL", original.Reason())
	assert.Equal(t, "NEW_REASON", withReason.Reason())

	withStatus := original.WithStatusCode(http.StatusNotFound)
	assert.Equal(t, http.StatusBadRequest, original.StatusCode())
	assert.Equal(t, http.StatusNotFound, withStatus.StatusCode())
}

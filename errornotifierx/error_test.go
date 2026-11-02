package errornotifierx

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test helper functions
func errA() error {
	return errors.New("base error")
}

func errB() error {
	return errors.Wrap(errA(), "wrapped in B")
}

func errC() error {
	return errors.Wrap(errB(), "wrapped in C")
}

func errWithoutStack() error {
	return &simpleError{msg: "simple error without stack"} //nolint:errhandle
}

type simpleError struct {
	msg string
}

func (e *simpleError) Error() string {
	return e.msg
}

func TestNewTrackedError(t *testing.T) {
	t.Run("returns nil for nil error", func(t *testing.T) {
		result := newTrackedError(nil)
		assert.Nil(t, result)
	})

	t.Run("creates tracked error for error chain with stack", func(t *testing.T) {
		err := errC()
		result := newTrackedError(err)

		require.NotNil(t, result)
		trackedErr := &trackedError{}
		ok := errors.As(result, &trackedErr)
		require.True(t, ok)

		assert.Equal(t, err.Error(), trackedErr.Error())
		assert.NotNil(t, trackedErr.tracer)
		assert.NotEmpty(t, trackedErr.StackTrace())
	})

	t.Run("returns original error without stack", func(t *testing.T) {
		err := errWithoutStack()
		result := newTrackedError(err)
		assert.Equal(t, err, result)
	})
}

func TestFindDeepestStackTracer(t *testing.T) {
	t.Run("returns nil for nil error", func(t *testing.T) {
		result := findDeepestStackTracer(nil)
		assert.Nil(t, result)
	})

	t.Run("returns nil for error without stack tracer", func(t *testing.T) {
		err := errWithoutStack()
		result := findDeepestStackTracer(err)
		assert.Nil(t, result)
	})

	t.Run("finds deepest stack tracer in chain", func(t *testing.T) {
		err := errC()
		result := findDeepestStackTracer(err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.StackTrace())
		assert.Equal(t, "base error", result.(error).Error())
	})

	t.Run("handles single error with stack", func(t *testing.T) {
		err := errors.New("single error with stack")
		result := findDeepestStackTracer(err)
		assert.NotNil(t, result)
		assert.NotEmpty(t, result.StackTrace())
	})
}

func TestErrorChainIntegration(t *testing.T) {
	originalErr := errC()
	trackedErr := newTrackedError(originalErr)
	require.NotNil(t, trackedErr)

	assert.Equal(t, originalErr.Error(), trackedErr.Error())
	assert.Contains(t, trackedErr.Error(), "wrapped in C")
	assert.Contains(t, trackedErr.Error(), "base error")
}

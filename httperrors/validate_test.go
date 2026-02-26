package httperrors

import (
	"context"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type simpleValidator struct {
	err error
}

func (v *simpleValidator) Validate() error {
	return v.err
}

type contextValidator struct {
	err error
}

func (v *contextValidator) Validate(ctx context.Context) error {
	return v.err
}

type dualValidator struct {
	simpleErr  error
	contextErr error
}

func (v *dualValidator) Validate(ctx context.Context) error {
	return v.contextErr
}

type notValidator struct{}

func TestValidate(t *testing.T) {
	ctx := context.Background()

	t.Run("ContextValidator success", func(t *testing.T) {
		input := &contextValidator{err: nil}
		err := Validate(ctx, input)
		assert.NoError(t, err)
	})

	t.Run("ContextValidator failure", func(t *testing.T) {
		input := &contextValidator{err: Error(http.StatusBadRequest, "BAD", "bad")}
		err := Validate(ctx, input)
		require.Error(t, err)

		s := Convert(err)
		assert.Equal(t, http.StatusBadRequest, s.StatusCode())
	})

	t.Run("Validator success", func(t *testing.T) {
		input := &simpleValidator{err: nil}
		err := Validate(ctx, input)
		assert.NoError(t, err)
	})

	t.Run("Validator failure", func(t *testing.T) {
		input := &simpleValidator{err: errors.New("validation failed")}
		err := Validate(ctx, input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "validation failed")
	})

	t.Run("ContextValidator takes priority over Validator", func(t *testing.T) {
		// Type that explicitly implements both interfaces
		input := &dualValidator{
			simpleErr:  errors.New("simple error"),
			contextErr: errors.New("context error"),
		}
		err := Validate(ctx, input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "context error")
	})

	t.Run("not a validator", func(t *testing.T) {
		input := &notValidator{}
		err := Validate(ctx, input)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "does not implement")
	})
}

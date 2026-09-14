package httperrors

import (
	"context"

	"github.com/pkg/errors"
)

type Validator interface {
	Validate() error
}

type ContextValidator interface {
	Validate(ctx context.Context) error
}

func Validate(ctx context.Context, input any) error {
	// Try ContextValidator interface first (has context support)
	if val, ok := input.(ContextValidator); ok {
		return val.Validate(ctx)
	}

	// Fallback to simple Validator interface
	if val, ok := input.(Validator); ok {
		return val.Validate()
	}

	return errors.New("input does not implement Validator or ContextValidator interface")
}

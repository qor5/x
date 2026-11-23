package dynconfx

import (
	"context"

	clone "github.com/huandu/go-clone/generic"
	"github.com/pkg/errors"
)

// ErrNotFound is returned when no configuration is stored: no Document row,
// or a Collection row that does not exist or holds no configuration yet.
var ErrNotFound = errors.New("dynconfx: config not found")

// Source reads a configuration. A service that only reads its configuration
// should depend on Source rather than on where the configuration is kept, so
// that a Document or a value from the deployment's files (Static) can be
// handed to it.
type Source[T any] interface {
	Get(ctx context.Context) (T, error)
}

type staticSource[T any] struct{ value T }

func (s staticSource[T]) Get(context.Context) (T, error) {
	return clone.Clone(s.value), nil
}

// Static returns a Source that always returns a copy of value.
func Static[T any](value T) Source[T] {
	return staticSource[T]{value: clone.Clone(value)}
}

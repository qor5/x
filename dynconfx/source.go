package dynconfx

import (
	"context"

	"github.com/pkg/errors"
)

// ErrNotFound is returned when no configuration is stored: no Document row,
// or a Collection row that does not exist or holds no configuration yet.
var ErrNotFound = errors.New("dynconfx: config not found")

// Source reads a configuration. A service that only reads its configuration
// should depend on Source rather than on where the configuration is kept, so
// that a Document, a value from the deployment's files (Static) or a value
// built for one request can all be handed to it.
type Source[T any] interface {
	Get(ctx context.Context) (T, error)
}

// SourceFunc adapts a function to Source.
type SourceFunc[T any] func(ctx context.Context) (T, error)

// Get calls f.
func (f SourceFunc[T]) Get(ctx context.Context) (T, error) {
	return f(ctx)
}

// Static returns a Source that always returns value.
func Static[T any](value T) Source[T] {
	return SourceFunc[T](func(context.Context) (T, error) {
		return value, nil
	})
}

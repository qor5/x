package hook

import (
	"context"
	"slices"
)

type Hook[T any] func(next T) T

// Prepend returns a hook function that first runs the given previous hook and then
// runs the remaining hooks in the chain. When executed, previous hook will be
// the first one to process the value, despite being appended to the end of the slice
// internally (since hooks are executed in reverse order).
// Nil hooks are automatically filtered out.
func Prepend[T any](previous Hook[T], chain ...Hook[T]) Hook[T] {
	return Chain(append(chain, previous)...)
}

// Chain returns a single hook function that chains the given hooks together.
// Nil hooks are automatically filtered out.
// When the hook is called, each hook in the chain is called with the result of
// the previous hook as its argument. The last hook in the chain is called with
// the argument passed to the chained hook. If the chain is empty or all hooks
// are nil, this function returns nil.
//
// Example:
//
//	hooks := []Hook[int]{
//	    func(next int) int { return next * 2 },
//	    func(next int) int { return next + 1 },
//	}
//	chainedHook := Chain(hooks...)
//	result := chainedHook(0)
//	fmt.Println(result) // Output: 2
func Chain[T any](hooks ...Hook[T]) Hook[T] {
	hooks = slices.DeleteFunc(hooks, func(h Hook[T]) bool { return h == nil })
	if len(hooks) == 0 {
		return nil
	}
	if len(hooks) == 1 {
		return hooks[0]
	}
	return func(next T) T {
		for i := len(hooks); i > 0; i-- {
			next = hooks[i-1](next)
		}
		return next
	}
}

// ContextHook is a hook function that takes a context and returns an error.
// This is useful for hooks that need to perform I/O operations or other
// context-aware processing.
type ContextHook[T any] func(ctx context.Context, next T) (T, error)

// PrependContext returns a hook function that first runs the given previous hook and then
// runs the remaining hooks in the chain. When executed, previous hook will be
// the first one to process the value.
// Nil hooks are automatically filtered out.
func PrependContext[T any](previous ContextHook[T], chain ...ContextHook[T]) ContextHook[T] {
	return ChainContext(append(chain, previous)...)
}

// ChainContext returns a single hook function that chains the given context hooks together.
// Nil hooks are automatically filtered out.
// When the hook is called, each hook in the chain is called with the result of
// the previous hook as its argument. If any hook returns an error, the chain
// stops immediately and returns that error.
// If the chain is empty or all hooks are nil, this function returns nil.
func ChainContext[T any](hooks ...ContextHook[T]) ContextHook[T] {
	hooks = slices.DeleteFunc(hooks, func(h ContextHook[T]) bool { return h == nil })
	if len(hooks) == 0 {
		return nil
	}
	if len(hooks) == 1 {
		return hooks[0]
	}
	return func(ctx context.Context, next T) (T, error) {
		var err error
		for i := len(hooks); i > 0; i-- {
			next, err = hooks[i-1](ctx, next)
			if err != nil {
				return next, err
			}
		}
		return next, nil
	}
}

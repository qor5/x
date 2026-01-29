package hook

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrepend(t *testing.T) {
	type CreateFunc func(input string) string

	baseFunc := func(input string) string {
		return fmt.Sprintf("Base: %s", input)
	}

	hook1 := func(next CreateFunc) CreateFunc {
		return func(input string) string {
			return fmt.Sprintf("Hook1 -> %s", next(input))
		}
	}

	hook2 := func(next CreateFunc) CreateFunc {
		return func(input string) string {
			return fmt.Sprintf("Hook2 -> %s", next(input))
		}
	}

	hook3 := func(next CreateFunc) CreateFunc {
		return func(input string) string {
			return fmt.Sprintf("Hook3 -> %s", next(input))
		}
	}

	combinedHook := Prepend(nil, hook1, hook2, hook3)
	{
		finalFunc := combinedHook(baseFunc)
		result := finalFunc("input")

		expected := "Hook1 -> Hook2 -> Hook3 -> Base: input"
		assert.Equal(t, expected, result, "The hooks should execute in the correct order")
	}

	combinedHook = Prepend(combinedHook, hook1, hook2, hook3)
	{
		finalFunc := combinedHook(baseFunc)
		result := finalFunc("input")

		expected := "Hook1 -> Hook2 -> Hook3 -> Hook1 -> Hook2 -> Hook3 -> Base: input"
		assert.Equal(t, expected, result, "The hooks should execute in the correct order")
	}

	combinedHook = Prepend(combinedHook) // nothing to append
	{
		finalFunc := combinedHook(baseFunc)
		result := finalFunc("input")

		expected := "Hook1 -> Hook2 -> Hook3 -> Hook1 -> Hook2 -> Hook3 -> Base: input"
		assert.Equal(t, expected, result, "The hooks should execute in the correct order")
	}

	combinedHook = Prepend[CreateFunc](nil)
	assert.Nil(t, combinedHook, "The hook should be nil if no hooks are added")
}

func TestChain(t *testing.T) {
	hooks := []Hook[int]{
		func(next int) int { return next * 2 },
		func(next int) int { return next + 1 },
	}

	chainedHook := Chain(hooks...)
	result := chainedHook(0)
	require.Equal(t, 2, result)
}

func TestChain_NilFiltering(t *testing.T) {
	hook1 := func(next int) int { return next * 2 }
	hook2 := func(next int) int { return next + 1 }

	t.Run("filters_nil_hooks", func(t *testing.T) {
		chainedHook := Chain[int](nil, hook1, nil, hook2, nil)
		result := chainedHook(0)
		assert.Equal(t, 2, result)
	})

	t.Run("all_nil_returns_nil", func(t *testing.T) {
		chainedHook := Chain[int](nil, nil, nil)
		assert.Nil(t, chainedHook)
	})

	t.Run("single_non_nil_returns_that_hook", func(t *testing.T) {
		chainedHook := Chain[int](nil, hook1, nil)
		assert.NotNil(t, chainedHook)
		result := chainedHook(5)
		assert.Equal(t, 10, result)
	})
}

func TestChainContext(t *testing.T) {
	hook1 := func(ctx context.Context, next int) (int, error) {
		return next * 2, nil
	}
	hook2 := func(ctx context.Context, next int) (int, error) {
		return next + 1, nil
	}
	errHook := func(ctx context.Context, next int) (int, error) {
		return 0, errors.New("hook error")
	}

	t.Run("chains_hooks_successfully", func(t *testing.T) {
		chainedHook := ChainContext(hook1, hook2)
		result, err := chainedHook(context.Background(), 0)
		require.NoError(t, err)
		assert.Equal(t, 2, result)
	})

	t.Run("stops_on_error", func(t *testing.T) {
		chainedHook := ChainContext(hook1, errHook, hook2)
		_, err := chainedHook(context.Background(), 0)
		require.Error(t, err)
		assert.Equal(t, "hook error", err.Error())
	})

	t.Run("filters_nil_hooks", func(t *testing.T) {
		chainedHook := ChainContext[int](nil, hook1, nil, hook2, nil)
		result, err := chainedHook(context.Background(), 0)
		require.NoError(t, err)
		assert.Equal(t, 2, result)
	})

	t.Run("all_nil_returns_nil", func(t *testing.T) {
		chainedHook := ChainContext[int](nil, nil, nil)
		assert.Nil(t, chainedHook)
	})

	t.Run("single_non_nil_returns_that_hook", func(t *testing.T) {
		chainedHook := ChainContext[int](nil, hook1, nil)
		assert.NotNil(t, chainedHook)
		result, err := chainedHook(context.Background(), 5)
		require.NoError(t, err)
		assert.Equal(t, 10, result)
	})

	t.Run("respects_context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		ctxHook := func(ctx context.Context, next int) (int, error) {
			if ctx.Err() != nil {
				return 0, ctx.Err()
			}
			return next * 2, nil
		}

		chainedHook := ChainContext(ctxHook)
		_, err := chainedHook(ctx, 5)
		require.Error(t, err)
		assert.Equal(t, context.Canceled, err)
	})
}

func TestPrependContext(t *testing.T) {
	hook1 := func(ctx context.Context, next int) (int, error) {
		return next * 2, nil
	}
	hook2 := func(ctx context.Context, next int) (int, error) {
		return next + 1, nil
	}
	hook3 := func(ctx context.Context, next int) (int, error) {
		return next + 10, nil
	}

	t.Run("prepends_hook_to_chain", func(t *testing.T) {
		chainedHook := PrependContext(hook1, hook2, hook3)
		result, err := chainedHook(context.Background(), 0)
		require.NoError(t, err)
		// Execution order: hook1 -> hook2 -> hook3
		// hook1(0) = 0 * 2 = 0, hook2(0) = 0 + 1 = 1, hook3(1) = 1 + 10 = 11
		assert.Equal(t, 11, result)
	})

	t.Run("nil_previous_is_filtered", func(t *testing.T) {
		chainedHook := PrependContext[int](nil, hook2, hook3)
		result, err := chainedHook(context.Background(), 0)
		require.NoError(t, err)
		assert.Equal(t, 11, result)
	})

	t.Run("all_nil_returns_nil", func(t *testing.T) {
		chainedHook := PrependContext[int](nil)
		assert.Nil(t, chainedHook)
	})
}

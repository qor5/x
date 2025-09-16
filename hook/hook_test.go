package hook

import (
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
	hooks := []Hook[int]{ // Hooks to chain together
		func(next int) int { return next * 2 }, // Multiply by 2
		func(next int) int { return next + 1 }, // Add 1
	}

	chainedHook := Chain(hooks...)
	result := chainedHook(0) // Run the hook chain with argument 0
	require.Equal(t, 2, result)
}

package clonex

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ifaceWrap struct {
	V any
}

type sample struct {
	N int
	S []int
	M map[string]int
}

func TestCloneSlowly_InterfaceValue_DeepCopy(t *testing.T) {
	orig := &sample{N: 1, S: []int{1, 2, 3}, M: map[string]int{"a": 1}}
	var v any = orig

	cloned := CloneSlowly(v)

	got, ok := cloned.(*sample)
	require.True(t, ok, "cloned type mismatch: %T", cloned)
	require.NotSame(t, orig, got)

	// Mutate clone and ensure original not changed.
	got.N = 2
	got.S[0] = 9
	got.M["a"] = 7

	assert.Equal(t, 1, orig.N)
	assert.Equal(t, 1, orig.S[0])
	assert.Equal(t, 1, orig.M["a"])
}

func TestCloneSlowly_StructWithInterfaceField_DeepCopy(t *testing.T) {
	origInner := &sample{N: 1, S: []int{1, 2, 3}, M: map[string]int{"a": 1}}
	w := ifaceWrap{V: origInner}

	cloned := CloneSlowly(w)

	gotInner, ok := cloned.V.(*sample)
	require.True(t, ok, "cloned inner type mismatch: %T", cloned.V)
	require.NotSame(t, origInner, gotInner)

	gotInner.S[0] = 9
	assert.Equal(t, 1, origInner.S[0])
}

func TestCloneSlowly_NilInterfaceValue(t *testing.T) {
	var v any = nil
	assert.Panics(t, func() {
		_ = CloneSlowly(v)
	}, "should panic when cloning nil interface value")
}

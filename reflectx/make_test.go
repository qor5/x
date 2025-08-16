package reflectx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	// Test basic types
	t.Run("basic types", func(t *testing.T) {
		// Integer
		assert.Equal(t, 0, Make[int]())
		assert.Equal(t, int8(0), Make[int8]())
		assert.Equal(t, int16(0), Make[int16]())
		assert.Equal(t, int32(0), Make[int32]())
		assert.Equal(t, int64(0), Make[int64]())

		// Unsigned integer
		assert.Equal(t, uint(0), Make[uint]())
		assert.Equal(t, uint8(0), Make[uint8]())
		assert.Equal(t, uint16(0), Make[uint16]())
		assert.Equal(t, uint32(0), Make[uint32]())
		assert.Equal(t, uint64(0), Make[uint64]())

		// Float
		assert.Equal(t, float32(0), Make[float32]())
		assert.Equal(t, float64(0), Make[float64]())

		// Complex
		assert.Equal(t, complex64(0), Make[complex64]())
		assert.Equal(t, complex128(0), Make[complex128]())

		// Other basic types
		assert.Equal(t, false, Make[bool]())
		assert.Equal(t, "", Make[string]())
	})

	// Test pointer types
	t.Run("pointer types", func(t *testing.T) {
		// Single pointer
		intPtr := Make[*int]()
		assert.NotNil(t, intPtr)
		assert.Equal(t, 0, *intPtr)

		// Double pointer
		intPtrPtr := Make[**int]()
		assert.NotNil(t, intPtrPtr)
		assert.NotNil(t, *intPtrPtr)
		assert.Equal(t, 0, **intPtrPtr)

		// Triple pointer
		intPtrPtrPtr := Make[***int]()
		assert.NotNil(t, intPtrPtrPtr)
		assert.NotNil(t, *intPtrPtrPtr)
		assert.NotNil(t, **intPtrPtrPtr)
		assert.Equal(t, 0, ***intPtrPtrPtr)
	})

	// Test collection types
	t.Run("collection types", func(t *testing.T) {
		// Map
		m := Make[map[string]int]()
		assert.NotNil(t, m)
		assert.Equal(t, 0, len(m))
		m["test"] = 1
		assert.Equal(t, 1, m["test"])

		// Slice
		s := Make[[]int]()
		assert.NotNil(t, s)
		assert.Equal(t, 0, len(s))
		s = append(s, 1)
		assert.Equal(t, 1, s[0])

		// Channel
		ch := Make[chan int]()
		assert.NotNil(t, ch)
		// Verify it's a working channel
		go func() { ch <- 1 }()
		assert.Equal(t, 1, <-ch)

		// Pointer to collection types
		mapPtr := Make[*map[string]int]()
		assert.NotNil(t, mapPtr)
		assert.NotNil(t, *mapPtr)
		(*mapPtr)["test"] = 1
		assert.Equal(t, 1, (*mapPtr)["test"])

		slicePtr := Make[*[]int]()
		assert.NotNil(t, slicePtr)
		assert.NotNil(t, *slicePtr)
		*slicePtr = append(*slicePtr, 1)
		assert.Equal(t, 1, (*slicePtr)[0])

		chanPtr := Make[*chan int]()
		assert.NotNil(t, chanPtr)
		assert.NotNil(t, *chanPtr)
		go func() { *chanPtr <- 1 }()
		assert.Equal(t, 1, <-*chanPtr)

		// Double pointer to collection types
		mapPtrPtr := Make[**map[string]int]()
		assert.NotNil(t, mapPtrPtr)
		assert.NotNil(t, *mapPtrPtr)
		assert.NotNil(t, **mapPtrPtr)
		(**mapPtrPtr)["test"] = 1
		assert.Equal(t, 1, (**mapPtrPtr)["test"])

		slicePtrPtr := Make[**[]int]()
		assert.NotNil(t, slicePtrPtr)
		assert.NotNil(t, *slicePtrPtr)
		assert.NotNil(t, **slicePtrPtr)
		**slicePtrPtr = append(**slicePtrPtr, 1)
		assert.Equal(t, 1, (**slicePtrPtr)[0])

		chanPtrPtr := Make[**chan int]()
		assert.NotNil(t, chanPtrPtr)
		assert.NotNil(t, *chanPtrPtr)
		assert.NotNil(t, **chanPtrPtr)
		go func() { **chanPtrPtr <- 1 }()
		assert.Equal(t, 1, <-**chanPtrPtr)
	})

	// Test struct types
	t.Run("struct types", func(t *testing.T) {
		type Foo struct {
			A int
			B string
			C *float64
		}

		// Non-pointer struct
		foo := Make[Foo]()
		assert.Equal(t, 0, foo.A)
		assert.Equal(t, "", foo.B)
		assert.Nil(t, foo.C)

		// Pointer to struct
		fooPtr := Make[*Foo]()
		assert.NotNil(t, fooPtr)
		assert.Equal(t, 0, fooPtr.A)
		assert.Equal(t, "", fooPtr.B)
		assert.Nil(t, fooPtr.C)

		// Double pointer to struct
		fooPtrPtr := Make[**Foo]()
		assert.NotNil(t, fooPtrPtr)
		assert.NotNil(t, *fooPtrPtr)
		assert.Equal(t, 0, (*fooPtrPtr).A)
		assert.Equal(t, "", (*fooPtrPtr).B)
		assert.Nil(t, (*fooPtrPtr).C)
	})

	// Test error cases
	t.Run("error cases", func(t *testing.T) {
		// Function type
		assert.Panics(t, func() {
			Make[func()]()
		}, "should panic for function type")

		// Interface type
		assert.Panics(t, func() {
			_ = Make[any]()
		}, "should panic for nil interface type")
	})
}

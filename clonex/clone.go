package clonex

import (
	"reflect"

	clone "github.com/huandu/go-clone/generic"
)

var cloner clone.Cloner

func init() {
	allocator := clone.NewAllocator(nil, &clone.AllocatorMethods{
		IsScalar: func(t reflect.Kind) bool {
			// ensure string is not treated as scalar (deepcopy)
			return t != reflect.String && clone.IsScalar(t)
		},
	})
	cloner = clone.MakeCloner(allocator)
}

func Clone[T any](t T) T {
	return cloner.Clone(t).(T)
}

func CloneSlowly[T any](t T) T {
	return cloner.CloneSlowly(t).(T)
}

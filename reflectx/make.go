package reflectx

import "reflect"

// Make creates a fully initialized instance of type T.
// For pointer types, it recursively initializes each level.
// For map, slice, and channel types, it returns initialized empty instances.
// For other types, it returns their zero values.
//
// It panics if:
// - T is a nil interface (type cannot be determined)
// - T is a function type (not supported)
func Make[T any]() T {
	var zero T
	// Optimize for simple kinds: return zero directly without MakeValue
	t := reflect.TypeOf(zero)
	if t == nil {
		panic("Make: cannot determine type from nil interface")
	}
	if t.Kind() == reflect.Func {
		panic("Make: function type is not supported")
	}
	switch t.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan:
		v := MakeValue(t)
		return v.Interface().(T)
	default:
		return zero
	}
}

// MakeValue recursively creates a new instance based on the given reflect.Type.
// The returned Value is always addressable (can be used with Set method).
// For map, slice, and channel types, it returns initialized empty instances.
// For pointer types, it recursively creates the element instance and returns a pointer to it.
// For other types, it returns their zero values.
func MakeValue(t reflect.Type) reflect.Value {
	if t == nil {
		panic("Make: cannot determine type from nil interface")
	}
	if t.Kind() == reflect.Func {
		panic("Make: function type is not supported")
	}

	if t.Kind() != reflect.Ptr {
		val := reflect.New(t).Elem()
		switch t.Kind() {
		case reflect.Map:
			val.Set(reflect.MakeMap(t))
		case reflect.Slice:
			val.Set(reflect.MakeSlice(t, 0, 0))
		case reflect.Chan:
			val.Set(reflect.MakeChan(t, 0))
		}
		return val
	}

	elemType := t.Elem()
	elemValue := MakeValue(elemType)
	ptrValue := reflect.New(elemType)
	ptrValue.Elem().Set(elemValue)
	return ptrValue
}

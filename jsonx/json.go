package jsonx

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Marshal converts a value to JSON bytes, with special handling for proto.Message types.
// Supports marshaling from proto.Message through multiple levels of pointers.
func Marshal(v any) ([]byte, error) {
	// Special handling for nil values
	if v == nil {
		return []byte("null"), nil
	}

	// Special handling for nil pointers
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr && rv.IsNil() {
		return []byte("null"), nil
	}

	pb, ok := v.(proto.Message)
	if !ok {
		// Try to find proto.Message in pointer chain
		pb = findProtoMessage(rv)
	}
	if pb != nil {
		b, err := (protojson.MarshalOptions{
			EmitUnpopulated: true,
		}).Marshal(pb)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal to protojson")
		}
		return b, nil
	}

	// Fallback to standard json marshal
	b, err := json.Marshal(v)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal to json")
	}
	return b, nil
}

// findProtoMessage recursively searches for a proto.Message in a reflect.Value.
// This is an optimized version that works directly with reflect.Value to avoid
// creating new reflect.Value instances during recursion.
func findProtoMessage(rv reflect.Value) proto.Message {
	kind := rv.Kind()

	if kind == reflect.Ptr {
		if rv.IsNil() {
			return nil
		}

		if rv.CanInterface() {
			if pb, ok := rv.Interface().(proto.Message); ok {
				return pb
			}
		}

		return findProtoMessage(rv.Elem())
	}

	if kind == reflect.Struct {
		newVal := reflect.New(rv.Type())
		newVal.Elem().Set(rv)
		if pb, ok := newVal.Interface().(proto.Message); ok {
			return pb
		}
	}

	return nil
}

var nullBytes = []byte("null")

// Unmarshal parses JSON bytes into a value, with special handling for proto.Message types.
// Supports unmarshaling into proto.Message through multiple levels of pointers.
func Unmarshal(data []byte, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return errors.New("unmarshal target must be a pointer")
	}

	if rv.IsNil() {
		// Maintain consistent behavior with json.Unmarshal
		return errors.New("cannot unmarshal into nil pointer")
	}

	if bytes.Equal(bytes.TrimSpace(data), nullBytes) {
		return nil
	}

	pb, setter := resolveProtoWithSetter(rv)
	if pb != nil {
		if err := (protojson.UnmarshalOptions{
			DiscardUnknown: true,
		}).Unmarshal(data, pb); err != nil {
			return errors.Wrap(err, "failed to unmarshal protojson")
		}

		setter()
		return nil
	}

	if err := json.Unmarshal(data, v); err != nil {
		return errors.Wrap(err, "failed to unmarshal json")
	}
	return nil
}

// resolveProtoWithSetter searches for a proto.Message in a pointer chain.
// Returns the found proto.Message and a function to set it back to the original destination.
func resolveProtoWithSetter(rv reflect.Value) (proto.Message, func()) {
	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		if pb, ok := rv.Interface().(proto.Message); ok {
			return pb, func() {}
		}

		elem := rv.Elem()

		if elem.Kind() == reflect.Ptr {
			if elem.IsNil() {
				newVal := reflect.New(elem.Type().Elem())

				if pb, ok := newVal.Interface().(proto.Message); ok {
					return pb, func() {
						elem.Set(newVal)
					}
				}

				deeperPb, deeperSetter := resolveProtoWithSetter(newVal)
				if deeperPb != nil {
					return deeperPb, func() {
						elem.Set(newVal)
						deeperSetter()
					}
				}
			} else {
				deeperPb, deeperSetter := resolveProtoWithSetter(elem)
				if deeperPb != nil {
					return deeperPb, deeperSetter
				}
			}
		}
	}

	return nil, nil
}

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	raw, err := Marshal(v)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := Indent(&buf, raw, prefix, indent); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func MustMarshal(v any) []byte {
	b, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func MarshalX[T ~string | ~[]byte](v any) (T, error) {
	var zero T
	b, err := Marshal(v)
	if err != nil {
		return zero, err
	}
	return T(b), nil
}

func MustMarshalX[T ~string | ~[]byte](v any) T {
	b, err := MarshalX[T](v)
	if err != nil {
		panic(err)
	}
	return b
}

func Beautify[T ~string | ~[]byte](rawJSON T) (T, error) {
	var zero T
	var buf bytes.Buffer
	if err := Indent(&buf, []byte(rawJSON), "", "  "); err != nil {
		return zero, err
	}
	return T(buf.Bytes()), nil
}

func MustBeautify[T ~string | ~[]byte](rawJSON T) T {
	v, err := Beautify(rawJSON)
	if err != nil {
		panic(err)
	}
	return v
}

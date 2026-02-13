package reflectx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ptr[T any](v T) *T { return &v }

type Inner struct {
	X string
	Y int
}

type AllTypes struct {
	// bool
	Bool bool

	// integer types
	Int   int
	Int8  int8
	Int16 int16
	Int32 int32
	Int64 int64

	// unsigned integer types
	Uint   uint
	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64

	// float types
	Float32 float32
	Float64 float64

	// complex types
	Complex64  complex64
	Complex128 complex128

	// string
	String string

	// byte slice ([]byte)
	Bytes []byte

	// slice
	Slice []string

	// map
	Map map[string]string

	// pointer
	Ptr *string

	// interface
	Iface any

	// chan
	Chan chan int

	// func
	Fn func()

	// array
	Array [2]int

	// time.Time (struct with IsZero)
	Time time.Time

	// embedded struct (anonymous)
	Inner

	// named struct
	Named Inner
}

func TestPatch(t *testing.T) {
	ch := make(chan int, 1)
	fn := func() {}
	now := time.Now()

	tests := []struct {
		name     string
		patch    *AllTypes
		dest     *AllTypes
		expected *AllTypes
	}{
		{
			name:  "zero patch does not overwrite dest",
			patch: &AllTypes{},
			dest: &AllTypes{
				Bool:       true,
				Int:        1,
				Int8:       2,
				Int16:      3,
				Int32:      4,
				Int64:      5,
				Uint:       6,
				Uint8:      7,
				Uint16:     8,
				Uint32:     9,
				Uint64:     10,
				Float32:    1.1,
				Float64:    2.2,
				Complex64:  1 + 2i,
				Complex128: 3 + 4i,
				String:     "dest",
				Bytes:      []byte("dest"),
				Slice:      []string{"a"},
				Map:        map[string]string{"k": "v"},
				Ptr:        ptr("dest"),
				Iface:      "dest",
				Chan:       ch,
				Fn:         fn,
				Array:      [2]int{1, 2},
				Time:       now,
				Inner:      Inner{X: "dx", Y: 1},
				Named:      Inner{X: "nx", Y: 2},
			},
			expected: &AllTypes{
				Bool:       true,
				Int:        1,
				Int8:       2,
				Int16:      3,
				Int32:      4,
				Int64:      5,
				Uint:       6,
				Uint8:      7,
				Uint16:     8,
				Uint32:     9,
				Uint64:     10,
				Float32:    1.1,
				Float64:    2.2,
				Complex64:  1 + 2i,
				Complex128: 3 + 4i,
				String:     "dest",
				Bytes:      []byte("dest"),
				Slice:      []string{"a"},
				Map:        map[string]string{"k": "v"},
				Ptr:        ptr("dest"),
				Iface:      "dest",
				Chan:       ch,
				Fn:         fn,
				Array:      [2]int{1, 2},
				Time:       now,
				Inner:      Inner{X: "dx", Y: 1},
				Named:      Inner{X: "nx", Y: 2},
			},
		},
		{
			name: "non-zero patch overwrites dest",
			patch: &AllTypes{
				Bool:       true,
				Int:        100,
				Int8:       101,
				Int16:      102,
				Int32:      103,
				Int64:      104,
				Uint:       200,
				Uint8:      201,
				Uint16:     202,
				Uint32:     203,
				Uint64:     204,
				Float32:    9.9,
				Float64:    8.8,
				Complex64:  5 + 6i,
				Complex128: 7 + 8i,
				String:     "patch",
				Bytes:      []byte("patch"),
				Slice:      []string{"x"},
				Map:        map[string]string{"pk": "pv"},
				Ptr:        ptr("patch"),
				Iface:      "patch",
				Chan:       ch,
				Fn:         fn,
				Array:      [2]int{9, 8},
				Time:       now,
				Inner:      Inner{X: "px", Y: 99},
				Named:      Inner{X: "pnx", Y: 88},
			},
			dest: &AllTypes{
				Bool:       false,
				Int:        1,
				Int8:       2,
				Int16:      3,
				Int32:      4,
				Int64:      5,
				Uint:       6,
				Uint8:      7,
				Uint16:     8,
				Uint32:     9,
				Uint64:     10,
				Float32:    1.1,
				Float64:    2.2,
				Complex64:  1 + 2i,
				Complex128: 3 + 4i,
				String:     "dest",
				Bytes:      []byte("dest"),
				Slice:      []string{"a"},
				Map:        map[string]string{"k": "v"},
				Ptr:        ptr("dest"),
				Iface:      "dest",
				Chan:       nil,
				Fn:         nil,
				Array:      [2]int{1, 2},
				Time:       time.Time{},
				Inner:      Inner{X: "dx", Y: 1},
				Named:      Inner{X: "nx", Y: 2},
			},
			expected: &AllTypes{
				Bool:       true,
				Int:        100,
				Int8:       101,
				Int16:      102,
				Int32:      103,
				Int64:      104,
				Uint:       200,
				Uint8:      201,
				Uint16:     202,
				Uint32:     203,
				Uint64:     204,
				Float32:    9.9,
				Float64:    8.8,
				Complex64:  5 + 6i,
				Complex128: 7 + 8i,
				String:     "patch",
				Bytes:      []byte("patch"),
				Slice:      []string{"x"},
				Map:        map[string]string{"pk": "pv"},
				Ptr:        ptr("patch"),
				Iface:      "patch",
				Chan:       ch,
				Fn:         fn,
				Array:      [2]int{9, 8},
				Time:       now,
				Inner:      Inner{X: "px", Y: 99},
				Named:      Inner{X: "pnx", Y: 88},
			},
		},
		{
			name: "partial patch: only non-zero fields overwrite",
			patch: &AllTypes{
				String: "patched",
				Int:    42,
				Ptr:    ptr("patched"),
				Slice:  []string{"new"},
				Inner:  Inner{X: "px"},
				Named:  Inner{Y: 77},
			},
			dest: &AllTypes{
				Bool:    true,
				Int:     1,
				String:  "dest",
				Float64: 3.14,
				Bytes:   []byte("dest"),
				Slice:   []string{"a", "b"},
				Map:     map[string]string{"k": "v"},
				Ptr:     ptr("dest"),
				Iface:   "dest",
				Array:   [2]int{1, 2},
				Time:    now,
				Inner:   Inner{X: "dx", Y: 10},
				Named:   Inner{X: "nx", Y: 20},
			},
			expected: &AllTypes{
				Bool:    true,
				Int:     42,
				String:  "patched",
				Float64: 3.14,
				Bytes:   []byte("dest"),
				Slice:   []string{"new"},
				Map:     map[string]string{"k": "v"},
				Ptr:     ptr("patched"),
				Iface:   "dest",
				Array:   [2]int{1, 2},
				Time:    now,
				Inner:   Inner{X: "px", Y: 10},
				Named:   Inner{X: "nx", Y: 77},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, Patch(tt.patch, tt.dest))
			// compare non-func/chan fields via assert.Equal
			// func and chan are not comparable with Equal, check them separately
			assert.Equal(t, tt.expected.Bool, tt.dest.Bool)
			assert.Equal(t, tt.expected.Int, tt.dest.Int)
			assert.Equal(t, tt.expected.Int8, tt.dest.Int8)
			assert.Equal(t, tt.expected.Int16, tt.dest.Int16)
			assert.Equal(t, tt.expected.Int32, tt.dest.Int32)
			assert.Equal(t, tt.expected.Int64, tt.dest.Int64)
			assert.Equal(t, tt.expected.Uint, tt.dest.Uint)
			assert.Equal(t, tt.expected.Uint8, tt.dest.Uint8)
			assert.Equal(t, tt.expected.Uint16, tt.dest.Uint16)
			assert.Equal(t, tt.expected.Uint32, tt.dest.Uint32)
			assert.Equal(t, tt.expected.Uint64, tt.dest.Uint64)
			assert.Equal(t, tt.expected.Float32, tt.dest.Float32)
			assert.Equal(t, tt.expected.Float64, tt.dest.Float64)
			assert.Equal(t, tt.expected.Complex64, tt.dest.Complex64)
			assert.Equal(t, tt.expected.Complex128, tt.dest.Complex128)
			assert.Equal(t, tt.expected.String, tt.dest.String)
			assert.Equal(t, tt.expected.Bytes, tt.dest.Bytes)
			assert.Equal(t, tt.expected.Slice, tt.dest.Slice)
			assert.Equal(t, tt.expected.Map, tt.dest.Map)
			assert.Equal(t, tt.expected.Ptr, tt.dest.Ptr)
			assert.Equal(t, tt.expected.Iface, tt.dest.Iface)
			assert.Equal(t, tt.expected.Array, tt.dest.Array)
			assert.Equal(t, tt.expected.Time, tt.dest.Time)
			assert.Equal(t, tt.expected.Inner, tt.dest.Inner)
			assert.Equal(t, tt.expected.Named, tt.dest.Named)
			if tt.expected.Chan != nil {
				assert.Equal(t, tt.expected.Chan, tt.dest.Chan)
			}
			if tt.expected.Fn != nil {
				assert.NotNil(t, tt.dest.Fn)
			}
		})
	}
}

func TestPatch_DeeplyNested(t *testing.T) {
	type L3 struct {
		A string
		B int
	}
	type L2 struct {
		L3
		C string
	}
	type L1 struct {
		L2
		D string
	}
	type Top struct {
		L1
		E string
	}

	tests := []struct {
		name     string
		patch    *Top
		dest     *Top
		expected *Top
	}{
		{
			name: "patch L3.A only, keep all others",
			patch: &Top{
				L1: L1{L2: L2{L3: L3{A: "patched"}}},
			},
			dest: &Top{
				L1: L1{
					L2: L2{
						L3: L3{A: "old-a", B: 1},
						C:  "old-c",
					},
					D: "old-d",
				},
				E: "old-e",
			},
			expected: &Top{
				L1: L1{
					L2: L2{
						L3: L3{A: "patched", B: 1},
						C:  "old-c",
					},
					D: "old-d",
				},
				E: "old-e",
			},
		},
		{
			name: "patch L3.B and L1.D, keep L3.A, L2.C, Top.E",
			patch: &Top{
				L1: L1{
					L2: L2{L3: L3{B: 99}},
					D:  "new-d",
				},
			},
			dest: &Top{
				L1: L1{
					L2: L2{
						L3: L3{A: "keep-a", B: 1},
						C:  "keep-c",
					},
					D: "old-d",
				},
				E: "keep-e",
			},
			expected: &Top{
				L1: L1{
					L2: L2{
						L3: L3{A: "keep-a", B: 99},
						C:  "keep-c",
					},
					D: "new-d",
				},
				E: "keep-e",
			},
		},
		{
			name:  "zero patch does not overwrite any level",
			patch: &Top{},
			dest: &Top{
				L1: L1{
					L2: L2{
						L3: L3{A: "a", B: 2},
						C:  "c",
					},
					D: "d",
				},
				E: "e",
			},
			expected: &Top{
				L1: L1{
					L2: L2{
						L3: L3{A: "a", B: 2},
						C:  "c",
					},
					D: "d",
				},
				E: "e",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, Patch(tt.patch, tt.dest))
			assert.Equal(t, tt.expected, tt.dest)
		})
	}
}

func TestPatch_EdgeCases(t *testing.T) {
	type Foo struct{ Name string }
	type Bar struct{ Name string }

	tests := []struct {
		name     string
		patch    any
		dest     any
		wantErr  bool
		wantDest any
	}{
		{
			name:     "nil patch: no error, dest unchanged",
			patch:    (*Foo)(nil),
			dest:     &Foo{Name: "old"},
			wantErr:  false,
			wantDest: &Foo{Name: "old"},
		},
		{
			name:    "nil dest",
			patch:   &Foo{Name: "new"},
			dest:    (*Foo)(nil),
			wantErr: true,
		},
		{
			name:    "patch non-pointer",
			patch:   Foo{},
			dest:    &Foo{},
			wantErr: true,
		},
		{
			name:    "dest non-pointer",
			patch:   &Foo{},
			dest:    Foo{},
			wantErr: true,
		},
		{
			name:    "type mismatch",
			patch:   &Foo{},
			dest:    &Bar{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Patch(tt.patch, tt.dest)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			if tt.wantDest != nil {
				assert.Equal(t, tt.wantDest, tt.dest)
			}
		})
	}
}

package jsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/qor5/x/v3/jsonx/testdata/gen"
)

func TestEqual(t *testing.T) {
	t.Run("error cases", func(t *testing.T) {
		assert.False(t, Equal(func() {}, func() {}))
		assert.False(t, Equal(make(chan int), make(chan int)))
		assert.False(t, Equal(make(chan int), make(chan string)))
	})

	t.Run("equal simple objects", func(t *testing.T) {
		a := map[string]any{"name": "John", "age": 30}
		b := map[string]any{"name": "John", "age": 30}

		assert.True(t, Equal(a, b))
	})

	t.Run("different simple objects", func(t *testing.T) {
		a := map[string]any{"name": "John", "age": 30}
		b := map[string]any{"name": "Jane", "age": 30}

		assert.False(t, Equal(a, b))
	})

	t.Run("different field types but same JSON representation", func(t *testing.T) {
		a := map[string]any{"value": 42}
		b := map[string]any{"value": 42.0}

		// These should be equal since JSON doesn't distinguish between int and float
		assert.True(t, Equal(a, b))
	})

	t.Run("different field order but same content", func(t *testing.T) {
		a := map[string]any{"name": "John", "age": 30}
		b := map[string]any{"age": 30, "name": "John"}

		assert.True(t, Equal(a, b))
	})

	t.Run("nested objects with same content", func(t *testing.T) {
		a := map[string]any{
			"person": map[string]any{
				"name": "John",
				"address": map[string]any{
					"city": "Shanghai",
					"zip":  "200000",
				},
			},
		}

		b := map[string]any{
			"person": map[string]any{
				"name": "John",
				"address": map[string]any{
					"city": "Shanghai",
					"zip":  "200000",
				},
			},
		}

		assert.True(t, Equal(a, b))
	})

	t.Run("nested objects with different content", func(t *testing.T) {
		a := map[string]any{
			"person": map[string]any{
				"name": "John",
				"address": map[string]any{
					"city": "Shanghai",
					"zip":  "200000",
				},
			},
		}

		b := map[string]any{
			"person": map[string]any{
				"name": "John",
				"address": map[string]any{
					"city": "Beijing", // Different city
					"zip":  "200000",
				},
			},
		}

		assert.False(t, Equal(a, b))
	})

	t.Run("arrays with same elements", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{1, 2, 3, 4}

		assert.True(t, Equal(a, b))
	})

	t.Run("arrays with different elements", func(t *testing.T) {
		a := []int{1, 2, 3, 4}
		b := []int{1, 2, 4, 3}

		assert.False(t, Equal(a, b))
	})

	t.Run("arrays with different types but same JSON representation", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []float64{1.0, 2.0, 3.0}

		assert.True(t, Equal(a, b))
	})

	t.Run("nil values", func(t *testing.T) {
		assert.True(t, Equal(nil, nil))
		assert.False(t, Equal(nil, ""))
		assert.False(t, Equal(nil, 0))
		assert.False(t, Equal(nil, false))
	})

	t.Run("zero values", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		a := Person{}
		b := Person{Name: "", Age: 0}

		assert.True(t, Equal(a, b))
	})

	t.Run("struct pointers", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		a := &Person{Name: "John", Age: 30}
		b := &Person{Name: "John", Age: 30}

		assert.True(t, Equal(a, b))
	})

	t.Run("proto messages", func(t *testing.T) {
		a := &gen.Identity{
			Id:        "user-123",
			GivenName: "John",
			Email:     "john@example.com",
		}

		b := &gen.Identity{
			Id:        "user-123",
			GivenName: "John",
			Email:     "john@example.com",
		}

		assert.True(t, Equal(a, b))

		// Different values
		c := &gen.Identity{
			Id:        "user-123",
			GivenName: "Jane", // Different name
			Email:     "john@example.com",
		}

		assert.False(t, Equal(a, c))
	})

	t.Run("complex protobuf objects with structpb", func(t *testing.T) {
		// Create two identical complex proto objects
		metadataValues1, _ := structpb.NewStruct(map[string]any{
			"preferences": map[string]any{
				"theme":         "dark",
				"notifications": true,
			},
			"devices": []any{"mobile", "tablet"},
		})

		metadataValues2, _ := structpb.NewStruct(map[string]any{
			"preferences": map[string]any{
				"theme":         "dark",
				"notifications": true,
			},
			"devices": []any{"mobile", "tablet"},
		})

		// Create matching timestamps to test time comparison
		timestamp1 := timestamppb.Now()
		timestamp2 := proto.Clone(timestamp1).(*timestamppb.Timestamp)

		a := &gen.UpdateIdentityInput{
			GivenName:      proto.String("John"),
			FamilyName:     proto.String("Doe"),
			DataConsent:    proto.Bool(true),
			MetadataPublic: metadataValues1,
			Birthday:       timestamp1,
		}

		b := &gen.UpdateIdentityInput{
			GivenName:      proto.String("John"),
			FamilyName:     proto.String("Doe"),
			DataConsent:    proto.Bool(true),
			MetadataPublic: metadataValues2,
			Birthday:       timestamp2,
		}

		assert.True(t, Equal(a, b))

		// Different structpb content
		metadataValues3, _ := structpb.NewStruct(map[string]any{
			"preferences": map[string]any{
				"theme":         "light", // Different value
				"notifications": true,
			},
			"devices": []any{"mobile", "tablet"},
		})

		c := &gen.UpdateIdentityInput{
			GivenName:      proto.String("John"),
			FamilyName:     proto.String("Doe"),
			DataConsent:    proto.Bool(true),
			MetadataPublic: metadataValues3,
		}

		assert.False(t, Equal(a, c))
	})

	t.Run("different types with same JSON representation", func(t *testing.T) {
		type Person1 struct {
			AgeX int    `json:"age"`
			Name string `json:"name"`
		}

		type Person2 struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		a := Person1{Name: "John", AgeX: 30}
		b := Person2{Name: "John", Age: 30}

		// These are different types but serialize to the same JSON
		assert.True(t, Equal(a, b))
	})

	t.Run("json number precision", func(t *testing.T) {
		a := map[string]any{"value": 123.456}
		b := map[string]any{"value": 123.456}

		assert.True(t, Equal(a, b))

		// Different precision
		c := map[string]any{"value": 123.45600000001}

		assert.False(t, Equal(a, c))
	})
}

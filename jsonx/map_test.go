package jsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToMap(t *testing.T) {
	t.Run("nil value", func(t *testing.T) {
		result, err := ToMap(nil)
		require.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("simple struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		p := Person{Name: "John", Age: 25}
		result, err := ToMap(p)
		require.NoError(t, err)

		expected := map[string]any{
			"Name": "John",
			"Age":  float64(25),
		}
		assert.Equal(t, expected, result)
	})

	t.Run("struct with json tags", func(t *testing.T) {
		type Product struct {
			Name        string  `json:"name"`
			Price       float64 `json:"price"`
			Category    string  `json:"category"`
			IsAvailable bool    `json:"isAvailable"`
		}

		p := Product{
			Name:        "Laptop",
			Price:       999.99,
			Category:    "Electronics",
			IsAvailable: true,
		}

		result, err := ToMap(p)
		require.NoError(t, err)

		expected := map[string]any{
			"name":        "Laptop",
			"price":       999.99,
			"category":    "Electronics",
			"isAvailable": true,
		}
		assert.Equal(t, expected, result)
	})

	t.Run("nested struct", func(t *testing.T) {
		type Address struct {
			Street string `json:"street"`
			City   string `json:"city"`
		}

		type User struct {
			Name    string  `json:"name"`
			Address Address `json:"address"`
		}

		u := User{
			Name: "Alice",
			Address: Address{
				Street: "123 Main St",
				City:   "New York",
			},
		}

		result, err := ToMap(u)
		require.NoError(t, err)

		expected := map[string]any{
			"name": "Alice",
			"address": map[string]any{
				"street": "123 Main St",
				"city":   "New York",
			},
		}
		assert.Equal(t, expected, result)
	})

	t.Run("map input", func(t *testing.T) {
		input := map[string]any{
			"key1": "value1",
			"key2": float64(42),
			"key3": true,
		}

		result, err := ToMap(input)
		require.NoError(t, err)
		assert.Equal(t, input, result)
	})

	t.Run("slice input", func(t *testing.T) {
		input := []string{"a", "b", "c"}

		result, err := ToMap(input)
		// This should fail because slices can't be converted to map[string]any
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("primitive types", func(t *testing.T) {
		// Test various primitive types that should fail
		testCases := []any{
			"string",
			42,
			3.14,
			true,
		}

		for _, tc := range testCases {
			result, err := ToMap(tc)
			assert.Error(t, err, "primitive type %T should return error", tc)
			assert.Nil(t, result)
		}
	})

	t.Run("pointer to struct", func(t *testing.T) {
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		p := &Person{Name: "Bob", Age: 30}
		result, err := ToMap(p)
		require.NoError(t, err)

		expected := map[string]any{
			"name": "Bob",
			"age":  float64(30),
		}
		assert.Equal(t, expected, result)
	})

	t.Run("struct with omitempty and zero values", func(t *testing.T) {
		type Config struct {
			Host    string `json:"host,omitempty"`
			Port    int    `json:"port,omitempty"`
			Enabled bool   `json:"enabled,omitempty"`
			Timeout int    `json:"timeout"`
		}

		c := Config{
			Host:    "",    // zero value with omitempty
			Port:    0,     // zero value with omitempty
			Enabled: false, // zero value with omitempty
			Timeout: 0,     // zero value without omitempty
		}

		result, err := ToMap(c)
		require.NoError(t, err)

		// Only timeout should be present due to omitempty tags
		expected := map[string]any{
			"timeout": float64(0),
		}
		assert.Equal(t, expected, result)
	})
}

func TestMustToMap(t *testing.T) {
	t.Run("successful conversion", func(t *testing.T) {
		type Person struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		p := Person{Name: "Charlie", Age: 35}
		result := MustToMap(p)

		expected := map[string]any{
			"name": "Charlie",
			"age":  float64(35),
		}
		assert.Equal(t, expected, result)
	})

	t.Run("nil input", func(t *testing.T) {
		result := MustToMap(nil)
		assert.Nil(t, result)
	})

	t.Run("panic on error", func(t *testing.T) {
		// Test that MustToMap panics when ToMap would return an error
		assert.Panics(t, func() {
			MustToMap("invalid input")
		}, "MustToMap should panic on invalid input")
	})

	t.Run("panic on unsupported type", func(t *testing.T) {
		// Function type should cause panic
		assert.Panics(t, func() {
			MustToMap(func() {})
		}, "MustToMap should panic on function type")
	})

	t.Run("panic on channel type", func(t *testing.T) {
		// Channel type should cause panic
		ch := make(chan int)
		assert.Panics(t, func() {
			MustToMap(ch)
		}, "MustToMap should panic on channel type")
	})
}

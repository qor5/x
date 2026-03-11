package httperrors

import (
	"net/http"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBadRequest(t *testing.T) {
	t.Run("with violations", func(t *testing.T) {
		fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
		fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")

		s := BadRequest(fv1, fv2)
		assert.Equal(t, http.StatusBadRequest, s.StatusCode())
		assert.Equal(t, ReasonInvalidArgument, s.Reason())
		assert.Equal(t, "invalid argument", s.Message())

		fvs := s.FieldViolations()
		require.Len(t, fvs, 2)
		assert.Equal(t, "email", fvs[0].Field())
		assert.Equal(t, "name", fvs[1].Field())
	})

	t.Run("empty violations returns OK", func(t *testing.T) {
		s := BadRequest()
		assert.Equal(t, http.StatusOK, s.StatusCode())
		assert.Equal(t, ReasonOK, s.Reason())
	})

	t.Run("nil input skipped", func(t *testing.T) {
		s := BadRequest(nil)
		assert.Equal(t, http.StatusOK, s.StatusCode())
	})
}

func TestUnprocessableEntity(t *testing.T) {
	t.Run("with violations", func(t *testing.T) {
		fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
		fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")

		s := UnprocessableEntity(fv1, fv2)
		assert.Equal(t, http.StatusUnprocessableEntity, s.StatusCode())
		assert.Equal(t, ReasonInvalidArgument, s.Reason())
		assert.Equal(t, "invalid argument", s.Message())

		fvs := s.FieldViolations()
		require.Len(t, fvs, 2)
		assert.Equal(t, "email", fvs[0].Field())
		assert.Equal(t, "name", fvs[1].Field())
	})

	t.Run("empty violations returns OK", func(t *testing.T) {
		s := UnprocessableEntity()
		assert.Equal(t, http.StatusOK, s.StatusCode())
		assert.Equal(t, ReasonOK, s.Reason())
	})

	t.Run("nil input skipped", func(t *testing.T) {
		s := UnprocessableEntity(nil)
		assert.Equal(t, http.StatusOK, s.StatusCode())
	})
}

func TestFormatField(t *testing.T) {
	tests := []struct {
		name       string
		field      string
		formatFunc func(string) string
		expected   string
	}{
		{
			name:       "simple field",
			field:      "user_name",
			formatFunc: lo.CamelCase,
			expected:   "userName",
		},
		{
			name:       "dotted path",
			field:      "user_info.street_name",
			formatFunc: lo.CamelCase,
			expected:   "userInfo.streetName",
		},
		{
			name:       "with array index",
			field:      "user_info.addresses[0].street_name",
			formatFunc: lo.CamelCase,
			expected:   "userInfo.addresses[0].streetName",
		},
		{
			name:       "multiple array indices",
			field:      "matrix_data[1][2]",
			formatFunc: lo.CamelCase,
			expected:   "matrixData[1][2]",
		},
		{
			name:       "identity function",
			field:      "user.addresses[0].street",
			formatFunc: func(s string) string { return s },
			expected:   "user.addresses[0].street",
		},
		{
			name:       "to upper",
			field:      "user.name",
			formatFunc: strings.ToUpper,
			expected:   "USER.NAME",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatField(tt.field, tt.formatFunc)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPrependField(t *testing.T) {
	fv1 := NewFieldViolation("email", "REQUIRED", "email is required")
	fv2 := NewFieldViolation("name", "TOO_SHORT", "name is too short")

	result := PrependField("user", fv1, fv2)
	require.Len(t, result, 2)
	assert.Equal(t, "user.email", result[0].Field())
	assert.Equal(t, "user.name", result[1].Field())
	assert.Equal(t, "email", fv1.Field())
	assert.Equal(t, "name", fv2.Field())
}

func TestFieldViolations_PrependField(t *testing.T) {
	fvs := FieldViolations{
		NewFieldViolation("street", "REQUIRED", "street is required"),
		NewFieldViolation("city", "REQUIRED", "city is required"),
	}

	result := fvs.PrependField("address")
	require.Len(t, result, 2)
	assert.Equal(t, "address.street", result[0].Field())
	assert.Equal(t, "address.city", result[1].Field())
	assert.Equal(t, "street", fvs[0].Field())
	assert.Equal(t, "city", fvs[1].Field())
}

func TestNewFieldViolation(t *testing.T) {
	fv := NewFieldViolation("email", "INVALID_FORMAT", "email format is invalid")

	assert.Equal(t, "email", fv.Field())
	assert.Equal(t, "INVALID_FORMAT", fv.Reason())
	assert.Equal(t, "email format is invalid", fv.Description())

	localized := fv.GetLocalized()
	require.NotNil(t, localized)
	assert.Equal(t, "INVALID_FORMAT", localized.Key())
	assert.Nil(t, localized.Args())
}

func TestNewFieldViolation_PanicOnEmptyField(t *testing.T) {
	assert.Panics(t, func() {
		NewFieldViolation("", "REQUIRED", "field is required")
	})
}

func TestNewFieldViolation_PanicOnEmptyReason(t *testing.T) {
	assert.Panics(t, func() {
		NewFieldViolation("email", "", "some description")
	})
}

func TestNewFieldViolationf(t *testing.T) {
	fv := NewFieldViolationf("age", "OUT_OF_RANGE", "age must be between %d and %d", 0, 120)
	assert.Equal(t, "age", fv.Field())
	assert.Equal(t, "OUT_OF_RANGE", fv.Reason())
	assert.Equal(t, "age must be between 0 and 120", fv.Description())
}

func TestFieldViolation_WithLocalized(t *testing.T) {
	fv := NewFieldViolation("email", "REQUIRED", "email is required").
		WithLocalized("custom.email.required", "arg1")

	localized := fv.GetLocalized()
	require.NotNil(t, localized)
	assert.Equal(t, "custom.email.required", localized.Key())
	assert.Equal(t, []any{"arg1"}, localized.Args())

	// Original field/reason/description preserved
	assert.Equal(t, "email", fv.Field())
	assert.Equal(t, "REQUIRED", fv.Reason())
	assert.Equal(t, "email is required", fv.Description())
}

func TestFieldViolation_WithLocalized_PanicOnEmptyKey(t *testing.T) {
	assert.Panics(t, func() {
		NewFieldViolation("email", "REQUIRED", "required").WithLocalized("")
	})
}

func TestFieldViolation_WithLocalizedArgs(t *testing.T) {
	fv := NewFieldViolation("name", "TOO_LONG", "name is too long").
		WithLocalizedArgs(100)

	localized := fv.GetLocalized()
	require.NotNil(t, localized)
	assert.Equal(t, "TOO_LONG", localized.Key())
	assert.Equal(t, []any{100}, localized.Args())
}

func TestFieldViolation_NilSafety(t *testing.T) {
	var fv *FieldViolation
	assert.Equal(t, "", fv.Field())
	assert.Equal(t, "", fv.Reason())
	assert.Equal(t, "", fv.Description())
	assert.Nil(t, fv.GetLocalized())
	assert.Nil(t, fv.GetLocalizedMessage())
	assert.Nil(t, fv.Clone())
}

func TestFieldViolation_Clone(t *testing.T) {
	fv := NewFieldViolation("email", "REQUIRED", "required").
		WithLocalized("key", "arg")

	cloned := fv.Clone()
	assert.Equal(t, fv.Field(), cloned.Field())
	assert.Equal(t, fv.Reason(), cloned.Reason())
	assert.Equal(t, fv.Description(), cloned.Description())
	assert.Equal(t, fv.GetLocalized().Key(), cloned.GetLocalized().Key())
}

func TestToFieldViolations(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := ToFieldViolations(nil, "field")
		assert.Nil(t, result)
	})

	t.Run("with field name", func(t *testing.T) {
		err := New(http.StatusNotFound, "NOT_FOUND", "user not found").Err()
		result := ToFieldViolations(err, "user")

		require.Len(t, result, 1)
		assert.Equal(t, "user", result[0].Field())
		assert.Equal(t, "NOT_FOUND", result[0].Reason())
		assert.Equal(t, "user not found", result[0].Description())
	})

	t.Run("empty field name skips main error", func(t *testing.T) {
		err := New(http.StatusNotFound, "NOT_FOUND", "not found").Err()
		result := ToFieldViolations(err, "")

		assert.Len(t, result, 0)
	})

	t.Run("with nested field violations", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "email is required")
		err := New(http.StatusBadRequest, ReasonInvalidArgument, "invalid").
			WithFieldViolations(fv).Err()

		result := ToFieldViolations(err, "user")
		require.Len(t, result, 2) // main error + nested
		assert.Equal(t, "user", result[0].Field())
		assert.Equal(t, "user.email", result[1].Field())
	})

	t.Run("nested violations without field prefix", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "required")
		err := New(http.StatusBadRequest, ReasonInvalidArgument, "invalid").
			WithFieldViolations(fv).Err()

		result := ToFieldViolations(err, "")
		require.Len(t, result, 1)
		assert.Equal(t, "email", result[0].Field())
	})
}

func TestFlattenFieldViolations(t *testing.T) {
	t.Run("single FieldViolation", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "required")
		result, err := FlattenFieldViolations(fv)
		require.NoError(t, err)
		require.Len(t, result, 1)
		assert.Equal(t, "email", result[0].Field())
	})

	t.Run("slice of FieldViolation", func(t *testing.T) {
		fvs := []*FieldViolation{
			NewFieldViolation("email", "REQUIRED", "required"),
			NewFieldViolation("name", "TOO_SHORT", "too short"),
		}
		result, err := FlattenFieldViolations(fvs)
		require.NoError(t, err)
		require.Len(t, result, 2)
	})

	t.Run("FieldViolations type", func(t *testing.T) {
		fvs := FieldViolations{
			NewFieldViolation("email", "REQUIRED", "required"),
		}
		result, err := FlattenFieldViolations(fvs)
		require.NoError(t, err)
		require.Len(t, result, 1)
	})

	t.Run("mixed types", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "required")
		fvs := FieldViolations{NewFieldViolation("name", "TOO_SHORT", "too short")}

		result, err := FlattenFieldViolations(fv, fvs)
		require.NoError(t, err)
		require.Len(t, result, 2)
	})

	t.Run("nil inputs skipped", func(t *testing.T) {
		result, err := FlattenFieldViolations(nil, nil)
		require.NoError(t, err)
		assert.Len(t, result, 0)
	})

	t.Run("unsupported type returns error", func(t *testing.T) {
		_, err := FlattenFieldViolations("not a field violation")
		require.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type")
	})
}

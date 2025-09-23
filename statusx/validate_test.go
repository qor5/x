package statusx

import (
	"context"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	testdatav1 "github.com/qor5/x/v3/statusx/gen/testdata/v1"
)

// Mock Validator interface implementation for testing (no context)
type mockSimpleValidator struct {
	violations []*FieldViolation
}

var _ Validator = (*mockSimpleValidator)(nil)

func (m *mockSimpleValidator) Validate() error {
	return BadRequest(m.violations).Err()
}

// Mock ContextValidator interface implementation for testing (with context)
type mockContextValidator struct {
	violations []*FieldViolation
}

var _ ContextValidator = (*mockContextValidator)(nil)

func (m *mockContextValidator) Validate(_ context.Context) error {
	return BadRequest(m.violations).Err()
}

func TestValidate(t *testing.T) {
	t.Run("validates ContextValidator interface", func(t *testing.T) {
		validator := &mockContextValidator{
			violations: []*FieldViolation{
				NewFieldViolation("name", "name.required", "Name is required"),
				NewFieldViolation("email", "email.invalid", "Invalid email"),
			},
		}

		err := Validate(context.Background(), validator)
		violations := ToFieldViolations(err, "")

		assert.Len(t, violations, 2)
		assert.Equal(t, "name", violations[0].Field())
		assert.Equal(t, "email", violations[1].Field())
	})

	t.Run("validates simple Validator interface", func(t *testing.T) {
		validator := &mockSimpleValidator{
			violations: []*FieldViolation{
				NewFieldViolation("username", "username.required", "Username is required"),
			},
		}

		err := Validate(context.Background(), validator)
		violations := ToFieldViolations(err, "")

		assert.Len(t, violations, 1)
		assert.Equal(t, "username", violations[0].Field())
	})

	t.Run("applies field prefix to ContextValidator results", func(t *testing.T) {
		validator := &mockContextValidator{
			violations: []*FieldViolation{
				NewFieldViolation("name", "name.required", "Name is required"),
			},
		}

		err := Validate(context.Background(), validator)
		violations := ToFieldViolations(err, "user")

		assert.Len(t, violations, 2)
		assert.Equal(t, "user", violations[0].Field())      // main error
		assert.Equal(t, "user.name", violations[1].Field()) // nested violation
	})

	t.Run("returns empty for valid ContextValidator", func(t *testing.T) {
		validator := &mockContextValidator{violations: nil}

		err := Validate(context.Background(), validator)
		violations := ToFieldViolations(err, "")

		assert.Empty(t, violations)
	})

	t.Run("returns empty for valid simple Validator", func(t *testing.T) {
		validator := &mockSimpleValidator{violations: nil}

		err := Validate(context.Background(), validator)
		violations := ToFieldViolations(err, "")

		assert.Empty(t, violations)
	})

	t.Run("returns empty for non-validator input", func(t *testing.T) {
		err := Validate(context.Background(), "not a validator")
		violations := ToFieldViolations(err, "")

		assert.Empty(t, violations)
	})
}

// Mock implementation of proto-gen-validate error interface
type mockPgvErr struct {
	field  string
	reason string
	name   string
}

func (m *mockPgvErr) Field() string     { return m.field }
func (m *mockPgvErr) Reason() string    { return m.reason }
func (m *mockPgvErr) Key() bool         { return false }
func (m *mockPgvErr) Cause() error      { return nil }
func (m *mockPgvErr) ErrorName() string { return m.name }

func TestFormatField(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		formatFunc func(string) string
		expected   string
	}{
		{
			name:       "simple field with camelCase",
			input:      "user_name",
			formatFunc: lo.CamelCase,
			expected:   "userName",
		},
		{
			name:       "nested field with camelCase",
			input:      "user_info.first_name",
			formatFunc: lo.CamelCase,
			expected:   "userInfo.firstName",
		},
		{
			name:       "array index preservation with camelCase",
			input:      "user_addresses[0]",
			formatFunc: lo.CamelCase,
			expected:   "userAddresses[0]",
		},
		{
			name:       "complex nested array with camelCase",
			input:      "user_info.addresses[0].street_name",
			formatFunc: lo.CamelCase,
			expected:   "userInfo.addresses[0].streetName",
		},
		{
			name:       "multiple array indices with camelCase",
			input:      "matrix_data[1][2].cell_value",
			formatFunc: lo.CamelCase,
			expected:   "matrixData[1][2].cellValue",
		},
		{
			name:       "mixed case preservation",
			input:      "user_ID.address_info[0].ZIP_code",
			formatFunc: lo.CamelCase,
			expected:   "userId.addressInfo[0].zipCode",
		},
		{
			name:       "snake_case transformation",
			input:      "userInfo.addressList[0].cityName",
			formatFunc: lo.SnakeCase,
			expected:   "user_info.address_list[0].city_name",
		},
		{
			name:       "kebab-case transformation",
			input:      "user_info.address_data[0].street_name",
			formatFunc: lo.KebabCase,
			expected:   "user-info.address-data[0].street-name",
		},
		{
			name:       "PascalCase transformation",
			input:      "user_name.address_list[0].city_code",
			formatFunc: lo.PascalCase,
			expected:   "UserName.AddressList[0].CityCode",
		},
		{
			name:       "SCREAMING_SNAKE_CASE transformation",
			input:      "user_info.address[0].zip_code",
			formatFunc: func(s string) string { return strings.ToUpper(lo.SnakeCase(s)) },
			expected:   "USER_INFO.ADDRESS[0].ZIP_CODE",
		},
		{
			name:       "uppercase transformation",
			input:      "user_name.address[0].city",
			formatFunc: strings.ToUpper,
			expected:   "USER_NAME.ADDRESS[0].CITY",
		},
		{
			name:       "empty field",
			input:      "",
			formatFunc: lo.CamelCase,
			expected:   "",
		},
		{
			name:       "field with no dots",
			input:      "simple_field",
			formatFunc: lo.CamelCase,
			expected:   "simpleField",
		},
		{
			name:       "field ending with array index",
			input:      "items[10]",
			formatFunc: lo.CamelCase,
			expected:   "items[10]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatField(tt.input, tt.formatFunc)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestConvertProtoGenErrToFV(t *testing.T) {
	tests := []struct {
		name          string
		inputField    string
		expectedField string
	}{
		{
			name:          "simple field",
			inputField:    "AAA",
			expectedField: "aaa",
		},
		{
			name:          "nested field",
			inputField:    "Aaa.BBB",
			expectedField: "aaa.bbb",
		},
		{
			name:          "array index",
			inputField:    "Aaa[0].ID",
			expectedField: "aaa[0].id",
		},
		{
			name:          "mixed case",
			inputField:    "ParentField.childField[1].GrandChild",
			expectedField: "parentField.childField[1].grandChild",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockErr := &mockPgvErr{
				field:  tt.inputField,
				reason: "test reason",
				name:   "test_name",
			}

			result := convertProtoGenErrToFV(mockErr)
			assert.Equal(t, tt.expectedField, result.Field)
			assert.Equal(t, "test reason", result.Description)
			assert.Equal(t, "PROTO_GEN_VALIDATE", result.Reason)
		})
	}
}

func TestLoCamelCase(t *testing.T) {
	got := lo.CamelCase("person[0]")
	assert.Equal(t, "person0", got)
}

func TestProtoGenValidate(t *testing.T) {
	type testCase struct {
		name        string
		input       any
		parentField string
		wantField   []string
	}
	cases := []testCase{
		{
			name: "one error validate all",
			input: &testdatav1.TestValidateError{
				GivenName:  lo.ToPtr("Terry"),
				FamilyName: lo.ToPtr("X"),
			},
			parentField: "",
			wantField:   []string{"familyName"},
		},
		{
			name: "one more error validate all",
			input: &testdatav1.TestValidateError{
				GivenName:  lo.ToPtr("T"),
				FamilyName: lo.ToPtr("X"),
			},
			parentField: "",
			wantField:   []string{"givenName", "familyName"},
		},
		{
			name: "error validate with parent field",
			input: &testdatav1.TestValidateError{
				GivenName:  lo.ToPtr("Terry"),
				FamilyName: lo.ToPtr("X"),
			},
			parentField: "UpdateInput[0]",
			wantField:   []string{"UpdateInput[0]", "UpdateInput[0].familyName"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			errs := Validate(context.Background(), c.input)
			violations := ToFieldViolations(errs, c.parentField)
			require.Len(t, violations, len(c.wantField))
			for i, e := range violations {
				assert.Equal(t, c.wantField[i], e.Field())
				// First violation is main error (BadRequest container), rest are proto-gen-validate violations
				if i == 0 && c.parentField != "" {
					assert.Equal(t, "INVALID_ARGUMENT", e.Reason()) // Main error from BadRequest
				} else {
					assert.Equal(t, "PROTO_GEN_VALIDATE", e.Reason()) // Nested violations
				}
			}
		})
	}
}

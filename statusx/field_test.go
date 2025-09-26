package statusx

import (
	"strings"
	"testing"

	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

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

func TestFlattenFieldViolations(t *testing.T) {
	t.Run("flatten mixed single and slice", func(t *testing.T) {
		// ✅ mixed types in a single call
		single1 := NewFieldViolation("email", "INVALID", "Email is invalid")
		slice1 := []*FieldViolation{
			NewFieldViolation("name", "REQUIRED", "Name is required"),
			NewFieldViolation("age", "TOO_YOUNG", "Age is too young"),
		}
		single2 := NewFieldViolation("phone", "INVALID_FORMAT", "Phone has invalid format")
		slice2 := FieldViolations{
			NewFieldViolation("email", "INVALID", "Email is invalid"),
			NewFieldViolation("name", "REQUIRED", "Name is required"),
			NewFieldViolation("age", "TOO_YOUNG", "Age is too young"),
		}

		// Mixed usage: single, slice, single
		result, err := FlattenFieldViolations(single1, slice1, single2, slice2)
		require.NoError(t, err)

		// Verify results
		assert.Len(t, result, 7)
		assert.Equal(t, "email", result[0].Field())
		assert.Equal(t, "name", result[1].Field())
		assert.Equal(t, "age", result[2].Field())
		assert.Equal(t, "phone", result[3].Field())
		assert.Equal(t, "email", result[4].Field())
		assert.Equal(t, "name", result[5].Field())
		assert.Equal(t, "age", result[6].Field())
	})

	t.Run("flatten handles nil", func(t *testing.T) {
		var nilSingle *FieldViolation
		slice := []*FieldViolation{
			NewFieldViolation("valid", "OK", "Valid field"),
		}

		result, err := FlattenFieldViolations(nilSingle, slice, nilSingle)
		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "valid", result[0].Field())
	})

	t.Run("flatten with ToFieldViolations from error", func(t *testing.T) {
		// Create a status error with field violations
		statusErr := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "validation failed").
			WithFieldViolations(
				NewFieldViolation("nested.field1", "INVALID", "Field1 is invalid"),
				NewFieldViolation("nested.field2", "REQUIRED", "Field2 is required"),
			).Err()

		// Convert error to field violations using ToFieldViolations
		errorViolations := ToFieldViolations(statusErr, "user")

		// Mix with other violations
		directViolations := []*FieldViolation{
			NewFieldViolation("email", "INVALID_FORMAT", "Email format is invalid"),
		}

		result, err := FlattenFieldViolations(errorViolations, directViolations)
		require.NoError(t, err)

		// Verify results: main error + 2 nested violations + 1 direct violation = 4 total
		assert.Len(t, result, 4)
		assert.Equal(t, "user", result[0].Field())               // main error field
		assert.Equal(t, "user.nested.field1", result[1].Field()) // prefixed nested field
		assert.Equal(t, "user.nested.field2", result[2].Field()) // prefixed nested field
		assert.Equal(t, "email", result[3].Field())              // direct violation
	})

	t.Run("flatten with Status.ToFieldViolations method", func(t *testing.T) {
		// Create a status with field violations
		status := New(codes.InvalidArgument, "USER_INVALID", "user data invalid").
			WithFieldViolations(
				NewFieldViolation("profile.bio", "TOO_LONG", "Bio is too long"),
				NewFieldViolation("profile.avatar", "INVALID_FORMAT", "Avatar format is invalid"),
			)

		// Use Status.ToFieldViolations method
		statusViolations := status.ToFieldViolations("request")

		// Mix with single violations
		singleViolation := NewFieldViolation("timestamp", "EXPIRED", "Timestamp has expired")

		result, err := FlattenFieldViolations(statusViolations, singleViolation)
		require.NoError(t, err)

		// Verify results: main status error + 2 nested violations + 1 single violation = 4 total
		assert.Len(t, result, 4)
		assert.Equal(t, "request", result[0].Field())                // main status field
		assert.Equal(t, "request.profile.bio", result[1].Field())    // prefixed nested field
		assert.Equal(t, "request.profile.avatar", result[2].Field()) // prefixed nested field
		assert.Equal(t, "timestamp", result[3].Field())              // single violation
		assert.Equal(t, "USER_INVALID", result[0].Reason())          // main status reason
	})

	t.Run("flatten protobuf field violations", func(t *testing.T) {
		// Test direct *errdetails.BadRequest_FieldViolation support
		pbSingle := &errdetails.BadRequest_FieldViolation{
			Field:       "pb_email",
			Reason:      "INVALID_FORMAT",
			Description: "Protobuf email invalid",
		}

		pbSlice := []*errdetails.BadRequest_FieldViolation{
			{Field: "pb_name", Reason: "REQUIRED", Description: "Protobuf name required"},
			{Field: "pb_age", Reason: "TOO_YOUNG", Description: "Protobuf age too young"},
		}

		// Mix protobuf types with internal types
		internal := NewFieldViolation("internal", "CUSTOM", "Custom internal validation")

		result, err := FlattenFieldViolations(pbSingle, pbSlice, internal)
		require.NoError(t, err)

		// Verify results: 1 single pb + 2 slice pb + 1 internal = 4 total
		assert.Len(t, result, 4)
		assert.Equal(t, "pb_email", result[0].Field())
		assert.Equal(t, "INVALID_FORMAT", result[0].Reason())
		assert.Equal(t, "Protobuf email invalid", result[0].Description())
		assert.Equal(t, "pb_name", result[1].Field())
		assert.Equal(t, "pb_age", result[2].Field())
		assert.Equal(t, "internal", result[3].Field())
		assert.Equal(t, "CUSTOM", result[3].Reason())
	})

	t.Run("flatten nil protobuf violations", func(t *testing.T) {
		var nilPbSingle *errdetails.BadRequest_FieldViolation
		var nilPbSlice []*errdetails.BadRequest_FieldViolation
		validInternal := NewFieldViolation("valid", "OK", "Valid field")

		result, err := FlattenFieldViolations(nilPbSingle, nilPbSlice, validInternal)
		require.NoError(t, err)

		// Should handle nil protobuf inputs gracefully
		assert.Len(t, result, 1)
		assert.Equal(t, "valid", result[0].Field())
	})
}

func TestToFieldViolations(t *testing.T) {
	t.Run("basic cases", func(t *testing.T) {
		// Nil input
		violations := ToFieldViolations(nil, "user")
		assert.Nil(t, violations)

		// Simple error
		err := New(codes.InvalidArgument, "INVALID_EMAIL", "Email format is invalid").Err()
		violations = ToFieldViolations(err, "email")
		assert.Len(t, violations, 1)
		assert.Equal(t, "email", violations[0].Field())
		assert.Equal(t, "INVALID_EMAIL", violations[0].Reason())
	})

	t.Run("localization priority", func(t *testing.T) {
		// LocalizedMessage takes priority over Localized
		localizedMsg := &errdetails.LocalizedMessage{Locale: "zh-CN", Message: "预置本地化消息"}
		err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "test").
			WithLocalized("template.key", "arg1").
			WithDetails(localizedMsg).Err()
		violations := ToFieldViolations(err, "test")

		assert.Len(t, violations, 1)
		assert.Nil(t, violations[0].Localized())
		assert.Equal(t, localizedMsg, violations[0].LocalizedMessage())
	})

	t.Run("nested violations", func(t *testing.T) {
		// Custom BadRequest
		err := New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "Form validation failed").
			WithFieldViolations(
				NewFieldViolation("email", "field.email.required", "Email is required"),
			).Err()
		violations := ToFieldViolations(err, "form")

		assert.Len(t, violations, 2)
		assert.Equal(t, "form", violations[0].Field())
		assert.Equal(t, "form.email", violations[1].Field())

		// Standard BadRequest
		standardBadRequest := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{Field: "username", Reason: "EXISTS", Description: "Username exists"},
			},
		}
		err = New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "User validation failed").
			WithDetails(standardBadRequest).Err()
		violations = ToFieldViolations(err, "user")

		assert.Len(t, violations, 2)
		assert.Equal(t, "user", violations[0].Field())
		assert.Equal(t, "user.username", violations[1].Field())
	})
}

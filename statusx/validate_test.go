package statusx

import (
	"context"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"

	testdatav1 "github.com/qor5/x/v3/statusx/gen/testdata/v1"
)

// Mock ValidatorX implementation for testing
type mockValidator struct {
	violations []*errdetails.BadRequest_FieldViolation
}

func (m *mockValidator) ValidateX(ctx context.Context) []*errdetails.BadRequest_FieldViolation {
	return m.violations
}

func TestValidateX(t *testing.T) {
	t.Run("validates ValidatorX interface", func(t *testing.T) {
		validator := &mockValidator{
			violations: []*errdetails.BadRequest_FieldViolation{
				{Field: "name", Description: "Name is required", Reason: "REQUIRED"},
				{Field: "email", Description: "Invalid email", Reason: "INVALID"},
			},
		}

		violations := ValidateX(context.Background(), validator)

		assert.Len(t, violations, 2)
		assert.Equal(t, "name", violations[0].Field)
		assert.Equal(t, "email", violations[1].Field)
	})

	t.Run("applies field prefix to ValidatorX results", func(t *testing.T) {
		validator := &mockValidator{
			violations: []*errdetails.BadRequest_FieldViolation{
				{Field: "name", Description: "Name is required", Reason: "REQUIRED"},
			},
		}

		violations := ValidateX(context.Background(), validator, WithFieldPrefix("user."))

		assert.Len(t, violations, 1)
		assert.Equal(t, "user.name", violations[0].Field)
	})

	t.Run("returns empty for valid ValidatorX", func(t *testing.T) {
		validator := &mockValidator{violations: nil}

		violations := ValidateX(context.Background(), validator)

		assert.Empty(t, violations)
	})

	t.Run("returns empty for non-validator input", func(t *testing.T) {
		violations := ValidateX(context.Background(), "not a validator")

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

			result := convertProtoGenErrToFV(mockErr, "")
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
		fieldPrefix string
		wantField   []string
	}
	cases := []testCase{
		{
			name: "one error validate all",
			input: &testdatav1.TestValidateError{
				GivenName:  lo.ToPtr("Terry"),
				FamilyName: lo.ToPtr("X"),
			},
			fieldPrefix: "",
			wantField:   []string{"familyName"},
		},
		{
			name: "one more error validate all",
			input: &testdatav1.TestValidateError{
				GivenName:  lo.ToPtr("T"),
				FamilyName: lo.ToPtr("X"),
			},
			fieldPrefix: "",
			wantField:   []string{"givenName", "familyName"},
		},
		{
			name: "error validate with field prefix",
			input: &testdatav1.TestValidateError{
				GivenName:  lo.ToPtr("Terry"),
				FamilyName: lo.ToPtr("X"),
			},
			fieldPrefix: "UpdateInput[0].",
			wantField:   []string{"updateInput[0].familyName"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			errs := ValidateX(context.Background(), c.input, WithFieldPrefix(c.fieldPrefix))
			require.Len(t, errs, len(c.wantField))
			for i, e := range errs {
				assert.Equal(t, c.wantField[i], e.GetField())
				assert.Equal(t, "PROTO_GEN_VALIDATE", e.GetReason())
			}
		})
	}
}

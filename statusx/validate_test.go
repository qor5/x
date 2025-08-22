package statusx

import (
	"context"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/qor5/x/v3/statusx/testdata"
)

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
		// expectedReason string
	}{
		{
			name:          "simple field",
			inputField:    "AAA",
			expectedField: "aaa",
			// expectedReason: "test_name:aaa",
		},
		{
			name:          "nested field",
			inputField:    "Aaa.BBB",
			expectedField: "aaa.bbb",
			// expectedReason: "test_name:aaa.bbb",
		},
		{
			name:          "array index",
			inputField:    "Aaa[0].ID",
			expectedField: "aaa[0].id",
			// expectedReason: "test_name:aaa[0].id",
		},
		{
			name:          "mixed case",
			inputField:    "ParentField.childField[1].GrandChild",
			expectedField: "parentField.childField[1].grandChild",
			// expectedReason: "test_name:parentField.childField[1].grandChild",
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
			// assert.Equal(t, tt.expectedReason, result.Reason)
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
		// wantReason  []string
	}
	cases := []testCase{
		{
			name: "one error validate all",
			input: &testdata.TestValidateError{
				GivenName:  lo.ToPtr("Terry"),
				FamilyName: lo.ToPtr("X"),
			},
			fieldPrefix: "",
			wantField:   []string{"familyName"},
			// wantReason:  []string{"TestValidateErrorValidationError:familyName"},
		},
		{
			name: "one more error validate all",
			input: &testdata.TestValidateError{
				GivenName:  lo.ToPtr("T"),
				FamilyName: lo.ToPtr("X"),
			},
			fieldPrefix: "",
			wantField:   []string{"givenName", "familyName"},
			// wantReason:  []string{"TestValidateErrorValidationError:givenName", "TestValidateErrorValidationError:familyName"},
		},
		{
			name: "error validate with field prefix",
			input: &testdata.TestValidateError{
				GivenName:  lo.ToPtr("Terry"),
				FamilyName: lo.ToPtr("X"),
			},
			fieldPrefix: "UpdateInput[0].",
			wantField:   []string{"updateInput[0].familyName"},
			// wantReason:  []string{"TestValidateErrorValidationError:updateInput[0].familyName"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			errs := ValidateX(context.Background(), c.input, WithFieldPrefix(c.fieldPrefix))
			require.Len(t, errs, len(c.wantField))
			for i, e := range errs {
				assert.Equal(t, c.wantField[i], e.GetField())
				// assert.Equal(t, c.wantReason[i], e.GetReason())
				assert.Equal(t, "PROTO_GEN_VALIDATE", e.GetReason())
			}
		})
	}
}

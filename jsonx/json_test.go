package jsonx

import (
	"encoding/json"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/qor5/x/v3/jsonx/testdata/gen"
)

func TestMarshal(t *testing.T) {
	t.Run("marshal *modelspb.Identity", func(t *testing.T) {
		idPtr := &gen.Identity{
			Id:          "test-id",
			Status:      "ACTIVE",
			GivenName:   "John",
			FamilyName:  "Doe",
			Email:       "test@example.com",
			DataConsent: true,
		}

		var a any = idPtr
		b0, err := Marshal(a)
		assert.NoError(t, err)

		b1, err := Marshal(idPtr)
		assert.NoError(t, err)

		b2, err := Marshal(&idPtr)
		assert.NoError(t, err)

		b3, err := Marshal(lo.ToPtr(&idPtr))
		assert.NoError(t, err)

		expectedJSON := `{
			"id": "test-id",
			"createdAt": null,
			"updatedAt": null,
			"status": "ACTIVE",
			"metadataPublic": null,
			"givenName": "John",
			"familyName": "Doe",
			"email": "test@example.com",
			"dataConsent": true,
			"totalSessionsCreated": 0
		}`

		assert.JSONEq(t, expectedJSON, string(b0))
		assert.JSONEq(t, expectedJSON, string(b1))
		assert.JSONEq(t, expectedJSON, string(b2))
		assert.JSONEq(t, expectedJSON, string(b3))
	})

	t.Run("marshal with nil pointer", func(t *testing.T) {
		var idPtr *gen.Identity = nil

		b1, err := Marshal(idPtr)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(b1))

		b2, err := Marshal(&idPtr)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(b2))

		b3, err := Marshal(lo.ToPtr(&idPtr))
		assert.NoError(t, err)
		assert.Equal(t, "null", string(b3))

		b4, err := Marshal(nil)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(b4))
	})

	t.Run("marshal regular json", func(t *testing.T) {
		data := map[string]any{"key": "value"}
		b, err := Marshal(data)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"key":"value"}`, string(b))
	})

	t.Run("marshal int", func(t *testing.T) {
		data := lo.ToPtr(123)
		b, err := Marshal(*data)
		assert.NoError(t, err)
		assert.JSONEq(t, `123`, string(b))
	})

	t.Run("marshal CreateSessionRequest with oneof field", func(t *testing.T) {
		req := &gen.RecoverIdentityRequest{
			Recovery: &gen.RecoverIdentityRequest_EmailRecovery{
				EmailRecovery: &gen.EmailRecovery{
					Email:       "test@example.com",
					NewPassword: "secure-password",
				},
			},
		}

		data, err := Marshal(req)
		assert.NoError(t, err)

		expectedJSON := `{
			"emailRecovery": {
				"email": "test@example.com",
				"newPassword": "secure-password"
			}
		}`
		assert.JSONEq(t, expectedJSON, string(data))
	})
}

func TestUnmarshal(t *testing.T) {
	t.Run("unmarshal regular json", func(t *testing.T) {
		data := `{"key":"value"}`
		m := map[string]any{}

		err := Unmarshal([]byte(data), &m)
		assert.NoError(t, err)

		expectedMap := map[string]any{"key": "value"}
		assert.Equal(t, expectedMap, m)
	})

	t.Run("unmarshal with error", func(t *testing.T) {
		// Invalid JSON syntax - missing quotes around string value
		data := `{"id":123, "status": invalid-json}`
		target := &gen.Identity{}

		err := Unmarshal([]byte(data), target)
		assert.Error(t, err)
	})

	t.Run("unmarshal RecoverIdentityRequest with oneof field using camelCase", func(t *testing.T) {
		data := `{"emailRecovery":{"email":"test@example.com","newPassword":"secure-password"}}`
		target := &gen.RecoverIdentityRequest{}

		err := Unmarshal([]byte(data), target)
		assert.NoError(t, err)

		expected := &gen.RecoverIdentityRequest{
			Recovery: &gen.RecoverIdentityRequest_EmailRecovery{
				EmailRecovery: &gen.EmailRecovery{
					Email:       "test@example.com",
					NewPassword: "secure-password",
				},
			},
		}
		assert.True(t, proto.Equal(target, expected), "proto message not equal with camelCase")
	})

	t.Run("unmarshal RecoverIdentityRequest with oneof field using snake_case", func(t *testing.T) {
		data := `{"email_recovery":{"email":"test@example.com","new_password":"secure-password"}}`
		target := &gen.RecoverIdentityRequest{}

		err := Unmarshal([]byte(data), target)
		assert.NoError(t, err)

		expected := &gen.RecoverIdentityRequest{
			Recovery: &gen.RecoverIdentityRequest_EmailRecovery{
				EmailRecovery: &gen.EmailRecovery{
					Email:       "test@example.com",
					NewPassword: "secure-password",
				},
			},
		}
		assert.True(t, proto.Equal(target, expected), "proto message not equal with snake_case")
	})

	t.Run("unmarshal RecoverIdentityRequest with empty object", func(t *testing.T) {
		data := `{"emailRecovery":{}}`
		target := &gen.RecoverIdentityRequest{}

		err := Unmarshal([]byte(data), target)
		assert.NoError(t, err)

		expected := &gen.RecoverIdentityRequest{
			Recovery: &gen.RecoverIdentityRequest_EmailRecovery{
				EmailRecovery: &gen.EmailRecovery{},
			},
		}
		assert.True(t, proto.Equal(target, expected), "proto message not equal with empty object")
	})

	t.Run("unmarshal RecoverIdentityRequest with unknown field", func(t *testing.T) {
		// Include metadata field along with standard fields
		data := `{
			"emailRecovery": {
				"email": "test@example.com", 
				"newPassword": "complex-password"
			},
			"unknownField": "value"
		}`

		target := &gen.RecoverIdentityRequest{}

		err := Unmarshal([]byte(data), target)
		assert.NoError(t, err)

		expected := &gen.RecoverIdentityRequest{
			Recovery: &gen.RecoverIdentityRequest_EmailRecovery{
				EmailRecovery: &gen.EmailRecovery{
					Email:       "test@example.com",
					NewPassword: "complex-password",
				},
			},
		}
		assert.True(t, proto.Equal(target, expected), "proto message not equal with complex structure")
	})

	t.Run("unmarshal **RecoverIdentityRequest", func(t *testing.T) {
		data := `{"emailRecovery":{"email":"test@example.com","newPassword":"secure-password"}}`
		target := &gen.RecoverIdentityRequest{}

		err := Unmarshal([]byte(data), &target)
		assert.NoError(t, err)

		expected := &gen.RecoverIdentityRequest{
			Recovery: &gen.RecoverIdentityRequest_EmailRecovery{
				EmailRecovery: &gen.EmailRecovery{
					Email:       "test@example.com",
					NewPassword: "secure-password",
				},
			},
		}
		assert.True(t, proto.Equal(target, expected), "proto message not equal with double pointer")
	})

	t.Run("unmarshal to nil", func(t *testing.T) {
		data := `{"emailRecovery":{"email":"test@example.com","newPassword":"secure-password"}}`

		var target *gen.RecoverIdentityRequest
		err := Unmarshal([]byte(data), target) // nil pointer
		assert.Error(t, err)

		var targetPtr **gen.RecoverIdentityRequest
		err = Unmarshal([]byte(data), targetPtr) // nil pointer
		assert.Error(t, err)

		err = Unmarshal([]byte(data), &target) // nil pointer ref
		assert.NoError(t, err)
		assert.NotNil(t, target)

		err = Unmarshal([]byte(data), &targetPtr) // ref of nil pointer's pointer
		assert.NoError(t, err)
		assert.NotNil(t, *targetPtr)

		expected := &gen.RecoverIdentityRequest{
			Recovery: &gen.RecoverIdentityRequest_EmailRecovery{
				EmailRecovery: &gen.EmailRecovery{
					Email:       "test@example.com",
					NewPassword: "secure-password",
				},
			},
		}
		assert.True(t, proto.Equal(target, expected), "proto message not equal after nil initialization")
		assert.True(t, proto.Equal(*targetPtr, expected), "proto message not equal after nil initialization")
	})
}

func TestMarshalNilBehavior(t *testing.T) {
	t.Run("json.Marshal nil handling", func(t *testing.T) {
		// Direct nil value
		jsonResult, err := json.Marshal(nil)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(jsonResult))

		// Nil pointer to proto Message
		var nilProto *gen.Identity = nil
		jsonResult, err = json.Marshal(nilProto)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(jsonResult))

		// Double nil pointer to proto Message
		var nilProtoPtr **gen.Identity = nil
		jsonResult, err = json.Marshal(nilProtoPtr)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(jsonResult))

		// Pointer to nil proto Message
		var nilProtoVal *gen.Identity = nil
		jsonResult, err = json.Marshal(&nilProtoVal)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(jsonResult))
	})

	t.Run("protojson.Marshal nil handling", func(t *testing.T) {
		// protojson.Marshal cannot handle direct nil

		// Nil pointer to proto Message
		var nilProto *gen.Identity = nil
		protojsonResult, err := protojson.Marshal(nilProto)
		assert.NoError(t, err)
		// TIPS: protojson.Marshal returns an empty message with default values
		// This behaves differently from json.Marshal, goes against intuition, I don't like it.
		assert.Equal(t, "{}", string(protojsonResult))
	})

	t.Run("our Marshal nil handling", func(t *testing.T) {
		// Direct nil
		mResult, err := Marshal(nil)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(mResult))

		// Nil proto.Message pointer (should match json.Marshal behavior)
		var nilProto *gen.Identity = nil
		mResult, err = Marshal(nilProto)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(mResult))

		// Pointer to nil proto.Message (should match json.Marshal behavior)
		var nilProtoVal *gen.Identity = nil
		mResult, err = Marshal(&nilProtoVal)
		assert.NoError(t, err)
		assert.Equal(t, "null", string(mResult))

		// Non-nil empty proto message (should use protojson)
		emptyProto := &gen.Identity{}
		mResult, err = Marshal(emptyProto)
		assert.NoError(t, err)
		assert.Contains(t, string(mResult), `"id":""`)
		assert.Contains(t, string(mResult), `"givenName":""`)
	})
}

func TestUnmarshalNilBehavior(t *testing.T) {
	t.Run("unmarshal null to Protocol Buffer messages", func(t *testing.T) {
		// JSON null value
		data := "null"

		// Test with pointers to non-nil Protocol Buffer messages
		var nilTarget *gen.Identity
		jsonTarget := &gen.Identity{Id: "original"}
		protojsonTarget := &gen.Identity{Id: "original"}
		jsonxTarget := &gen.Identity{Id: "original"}

		// Standard json.Unmarshal keeps the target message unmodified
		err := json.Unmarshal([]byte(data), nilTarget)
		assert.ErrorContains(t, err, "json: Unmarshal(nil")

		err = json.Unmarshal([]byte(data), jsonTarget)
		assert.NoError(t, err)
		assert.Equal(t, "original", jsonTarget.Id, "json.Unmarshal did not modify the message when unmarshaling null")

		// protojson.Unmarshal returns an error
		assert.Panics(t, func() {
			_ = protojson.Unmarshal([]byte(data), nilTarget)
		})

		err = protojson.Unmarshal([]byte(data), protojsonTarget)
		assert.ErrorContains(t, err, "unexpected token null")

		// jsonx.Unmarshal follows the same behavior of json.Unmarshal
		err = Unmarshal([]byte(data), nilTarget)
		assert.Error(t, err)

		err = Unmarshal([]byte(data), jsonxTarget)
		assert.NoError(t, err)
		assert.Equal(t, "original", jsonxTarget.Id, "jsonx.Unmarshal did not modify the message when unmarshaling null")
	})
}

func TestBeautify(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "beautify simple object",
			input: `{"name":"John","age":30}`,
			want: `{
  "name": "John",
  "age": 30
}`,
		},
		{
			name:  "beautify nested object",
			input: `{"user":{"name":"John","age":30}}`,
			want: `{
  "user": {
    "name": "John",
    "age": 30
  }
}`,
		},
		{
			name:    "invalid json",
			input:   `{"name":"John"`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Beautify(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMustBeautify(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      string
		wantPanic bool
	}{
		{
			name:  "beautify simple object",
			input: `{"name":"John","age":30}`,
			want: `{
  "name": "John",
  "age": 30
}`,
		},
		{
			name:      "invalid json (panic)",
			input:     `{"name":"John"`,
			wantPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				require.Panics(t, func() {
					MustBeautify(tt.input)
				})
				return
			}
			got := MustBeautify(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestMarshalX(t *testing.T) {
	t.Run("marshal struct", func(t *testing.T) {
		got := MustMarshalX[string](struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{Name: "John", Age: 30})
		require.Equal(t, `{"name":"John","age":30}`, got)
	})

	t.Run("marshal struct with []byte", func(t *testing.T) {
		got := MustMarshalX[[]byte](struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{Name: "John", Age: 30})
		require.Equal(t, `{"name":"John","age":30}`, string(got))
	})

	t.Run("marshal struct with string type", func(t *testing.T) {
		type FooString string
		got := MustMarshalX[FooString](struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{Name: "John", Age: 30})
		require.Equal(t, `{"name":"John","age":30}`, string(got))
		t.Logf("got: %s", MustBeautify(got))
	})

	t.Run("marshal struct with []byte type", func(t *testing.T) {
		type FooBytes []byte
		got := MustMarshalX[FooBytes](struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{Name: "John", Age: 30})
		require.Equal(t, `{"name":"John","age":30}`, string(got))
		t.Logf("got: %s", MustBeautify(got))
	})
}

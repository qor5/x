package jsonx

import (
	"testing"

	"github.com/huandu/go-clone"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/qor5/x/v3/jsonx/testdata/gen"
)

func TestApplyPatch(t *testing.T) {
	t.Run("successfully applies patch to struct", func(t *testing.T) {
		type Person struct {
			Name    string `json:"name"`
			Age     int    `json:"age"`
			Address string `json:"address,omitempty"`
		}

		// Initial data
		person := Person{
			Name: "张三",
			Age:  25,
		}

		// Patch to apply
		patch := []byte(`{"age": 30, "address": "北京市"}`)

		// Apply patch
		err := Patch(patch, &person)
		assert.NoError(t, err)

		// Verify result
		assert.Equal(t, "张三", person.Name)
		assert.Equal(t, 30, person.Age)
		assert.Equal(t, "北京市", person.Address)
	})

	t.Run("handles empty patch", func(t *testing.T) {
		type Config struct {
			Enabled bool   `json:"enabled"`
			Count   int    `json:"count"`
			Key     string `json:"key"`
		}

		// Initial data
		config := Config{
			Enabled: true,
			Count:   5,
			Key:     "test-key",
		}

		// Empty patch
		patch := []byte(`{}`)

		// Apply patch
		err := Patch(patch, &config)
		assert.NoError(t, err)

		// Verify no changes
		assert.Equal(t, true, config.Enabled)
		assert.Equal(t, 5, config.Count)
		assert.Equal(t, "test-key", config.Key)
	})

	t.Run("handles nested objects", func(t *testing.T) {
		type Address struct {
			Street string `json:"street"`
			City   string `json:"city"`
		}

		type User struct {
			Name    string  `json:"name"`
			Address Address `json:"address"`
		}

		// Initial data
		user := User{
			Name: "李四",
			Address: Address{
				Street: "文化路",
				City:   "上海",
			},
		}

		// Patch to apply
		patch := []byte(`{"address": {"city": "广州"}}`)

		// Apply patch
		err := Patch(patch, &user)
		assert.NoError(t, err)

		// Verify result
		assert.Equal(t, "李四", user.Name)
		assert.Equal(t, "文化路", user.Address.Street)
		assert.Equal(t, "广州", user.Address.City)
	})

	t.Run("handles array modification", func(t *testing.T) {
		type Project struct {
			Name  string   `json:"name"`
			Tags  []string `json:"tags"`
			Users []string `json:"users"`
		}

		// Initial data
		project := Project{
			Name:  "JSONX",
			Tags:  []string{"go", "api"},
			Users: []string{"user1", "user2"},
		}

		// Patch to apply (replace entire arrays)
		patch := []byte(`{"tags": ["go", "api", "cloud"], "users": ["user3"]}`)

		// Apply patch
		err := Patch(patch, &project)
		assert.NoError(t, err)

		// Verify result
		assert.Equal(t, "JSONX", project.Name)
		assert.Equal(t, []string{"go", "api", "cloud"}, project.Tags)
		assert.Equal(t, []string{"user3"}, project.Users)
	})

	t.Run("handles invalid JSON in patch", func(t *testing.T) {
		type Simple struct {
			Value string `json:"value"`
		}

		simple := Simple{Value: "original"}

		// Invalid JSON patch
		patch := []byte(`{"value": "new"`)

		// Apply patch should return error
		err := Patch(patch, &simple)
		assert.Error(t, err)

		// Original should be unchanged
		assert.Equal(t, "original", simple.Value)
	})

	t.Run("handles nil destination", func(t *testing.T) {
		var nilDest *map[string]any
		patch := []byte(`{"key": "value"}`)

		err := Patch(patch, nilDest)
		assert.Error(t, err)
	})

	t.Run("handles map type", func(t *testing.T) {
		// Map destination
		dest := map[string]any{
			"name": "王五",
			"data": map[string]any{
				"id":    123,
				"valid": true,
			},
		}

		// Patch to apply
		patch := []byte(`{"name": "赵六", "data": {"id": 456}}`)

		// Apply patch
		err := Patch(patch, &dest)
		assert.NoError(t, err)

		// Verify result
		assert.Equal(t, "赵六", dest["name"])
		data, ok := dest["data"].(map[string]any)
		assert.True(t, ok)
		assert.Equal(t, float64(456), data["id"])
		assert.Equal(t, true, data["valid"])
	})

	t.Run("correctly replaces fields with null", func(t *testing.T) {
		// Original test derived from comment in the ApplyPatch function
		// which says null values indicate field addition or overriding
		type TestStruct struct {
			Field1 *string `json:"field1"`
			Field2 *int    `json:"field2"`
		}

		str := "value"
		num := 42
		dest := TestStruct{
			Field1: &str,
			Field2: &num,
		}

		// Patch with null value
		patch := []byte(`{"field1": null}`)

		// Apply patch
		err := Patch(patch, &dest)
		assert.NoError(t, err)

		// Verify Field1 is set to null/nil
		assert.Nil(t, dest.Field1)
		// Field2 remains unchanged
		assert.NotNil(t, dest.Field2)
		assert.Equal(t, 42, *dest.Field2)
	})

	t.Run("handles deeply nested protobuf structures with identitypb types", func(t *testing.T) {
		// Create the UpdateIdentityInput protobuf object
		metadataValues, err := structpb.NewStruct(map[string]any{
			"language": "en",
			"region":   "US",
			"preferences": map[string]any{
				"theme":   "dark",
				"timeout": 30,
			},
		})
		assert.NoError(t, err)

		// Create real protobuf objects from identity package
		updateInput := &gen.UpdateIdentityInput{
			GivenName:      proto.String("John"),
			FamilyName:     proto.String("Doe"),
			DataConsent:    proto.Bool(true),
			MetadataPublic: metadataValues,
		}

		emailCreation := &gen.EmailCreation{
			Email:    "test@example.com",
			Password: "securepassword",
			Input:    updateInput,
		}

		createRequest := &gen.CreateIdentityRequest{
			Creation: &gen.CreateIdentityRequest_EmailCreation{
				EmailCreation: emailCreation,
			},
		}

		// Create patch with deeply nested changes
		patch := []byte(`{
			"emailCreation": {
				"input": {
					"givenName": "Jane",
					"locationConsent": true,
					"metadataPublic": {
						"region": "UK",
						"preferences": {
							"theme": "light",
							"notifications": true
						},
						"devices": ["mobile", "tablet"]
					},
					"emailMarketingConsent": true
				}
			}
		}`)

		// Apply patch
		err = Patch(patch, createRequest)
		assert.NoError(t, err)

		// Extract the email creation from the oneof field
		resultEmailCreation := createRequest.GetEmailCreation()
		assert.NotNil(t, resultEmailCreation)

		// Verify top-level fields remain unchanged
		assert.Equal(t, "test@example.com", resultEmailCreation.Email)
		assert.Equal(t, "securepassword", resultEmailCreation.Password)

		// Verify updated input fields
		input := resultEmailCreation.Input
		assert.Equal(t, "Jane", input.GetGivenName())
		assert.Equal(t, "Doe", input.GetFamilyName()) // Unchanged

		// Verify nested metadata structure was properly merged
		fields := input.MetadataPublic.GetFields()
		assert.Equal(t, "en", fields["language"].GetStringValue())
		assert.Equal(t, "UK", fields["region"].GetStringValue())

		// Verify preferences were merged
		preferencesValue := fields["preferences"].GetStructValue()
		assert.NotNil(t, preferencesValue)
		preferences := preferencesValue.GetFields()
		assert.Equal(t, "light", preferences["theme"].GetStringValue())
		assert.Equal(t, float64(30), preferences["timeout"].GetNumberValue())
		assert.Equal(t, true, preferences["notifications"].GetBoolValue())

		// Verify new array field was added
		devices := fields["devices"].GetListValue().Values
		assert.Equal(t, 2, len(devices))
		assert.Equal(t, "mobile", devices[0].GetStringValue())
		assert.Equal(t, "tablet", devices[1].GetStringValue())
	})

	t.Run("regular structs preserve original values when patch has type mismatch, unlike protobuf", func(t *testing.T) {
		// Create a regular struct with various field types
		type UpdateIdentityInput struct {
			GivenName   string `json:"givenName"`
			FamilyName  string `json:"familyName"`
			DataConsent bool   `json:"dataConsent"`
		}

		// Create the initial object
		updateInput := &UpdateIdentityInput{
			GivenName:   "John",
			FamilyName:  "Doe",
			DataConsent: true,
		}

		// Make a deep copy of the original for comparison
		originalInput := clone.Clone(updateInput)

		// Patch with type mismatches
		patch := []byte(`{
			"givenName": "123",
			"familyName": 456,
			"dataConsent": "yes",
			"metadataPublic": {"invalid": {"nested": [1, 2, "a"]}}
		}`)

		// Apply patch should return error due to type mismatch
		err := Patch(patch, updateInput)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal")

		assert.NotEqual(t, originalInput, updateInput)

		assert.Equal(t, "123", updateInput.GivenName)
		assert.Equal(t, "Doe", updateInput.FamilyName) // Unchanged
		assert.True(t, updateInput.DataConsent)        // Unchanged
	})

	t.Run("returns error and doesn't preserve original values when patch has type mismatch for protobuf fields", func(t *testing.T) {
		// Create a simple protobuf message
		updateInput := &gen.UpdateIdentityInput{
			GivenName:   proto.String("John"),
			FamilyName:  proto.String("Doe"),
			DataConsent: proto.Bool(true),
		}

		// Clone the original object to compare later
		originalInput := proto.Clone(updateInput).(*gen.UpdateIdentityInput)

		// Patch with type mismatches - this is an invalid patch for protobuf since
		// protojson requires types to match the schema
		patch := []byte(`{
			"givenName": "123",
			"familyName": 456,
			"dataConsent": "yes",
			"metadataPublic": {"invalid": {"nested": [1, 2, "a"]}}
		}`)

		// Apply patch should return error due to type mismatch
		err := Patch(patch, updateInput)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to unmarshal")

		// Note: ApplyPatch implementation doesn't guarantee original values are preserved
		// on error. This test documents the actual behavior that values may be reset or changed
		// during a failed patch operation.
		assert.False(t, proto.Equal(updateInput, originalInput),
			"Expected object to be modified even when patch fails due to type mismatch")

		// The current implementation will likely reset fields to zero values
		assert.Equal(t, "123", updateInput.GetGivenName())
		assert.Empty(t, updateInput.GetFamilyName())  // Reset to empty string
		assert.False(t, updateInput.GetDataConsent()) // Reset to false

		// TIPS: Not like regular structs, protobuf objects will reset fields to zero values first then apply the patch.
	})
}

func TestCopy(t *testing.T) {
	tests := []struct {
		name    string
		src     any
		dst     any
		check   func(t *testing.T, dst any)
		wantErr bool
	}{
		{
			name: "copy map to map",
			src:  map[string]string{"name": "John"},
			dst:  &map[string]string{},
			check: func(t *testing.T, dst any) {
				actual := dst.(*map[string]string)
				expected := map[string]string{"name": "John"}
				require.Equal(t, expected, *actual)
			},
		},
		{
			name: "copy struct to struct",
			src: struct {
				Name string `json:"name"`
			}{Name: "John"},
			dst: &struct {
				Name string `json:"name"`
			}{},
			check: func(t *testing.T, dst any) {
				actual := dst.(*struct {
					Name string `json:"name"`
				})
				expected := struct {
					Name string `json:"name"`
				}{Name: "John"}
				require.Equal(t, expected, *actual)
			},
		},
		{
			name: "merge metadata scenario",
			src: &structpb.Struct{
				Fields: map[string]*structpb.Value{
					"region": structpb.NewStringValue("US"),
					"preferences": structpb.NewStructValue(&structpb.Struct{
						Fields: map[string]*structpb.Value{
							"theme": structpb.NewStringValue("dark"),
						},
					}),
				},
			},
			dst: &struct {
				Region      string `json:"region"`
				Language    string `json:"language"`
				Preferences struct {
					Theme    string `json:"theme"`
					FontSize int    `json:"fontSize"`
				} `json:"preferences"`
			}{
				Language: "en",
				Preferences: struct {
					Theme    string `json:"theme"`
					FontSize int    `json:"fontSize"`
				}{
					FontSize: 14,
				},
			},
			check: func(t *testing.T, dst any) {
				actual := dst.(*struct {
					Region      string `json:"region"`
					Language    string `json:"language"`
					Preferences struct {
						Theme    string `json:"theme"`
						FontSize int    `json:"fontSize"`
					} `json:"preferences"`
				})

				// Check that values were correctly merged
				assert.Equal(t, "US", actual.Region)
				assert.Equal(t, "en", actual.Language) // Original value preserved
				assert.Equal(t, "dark", actual.Preferences.Theme)
				assert.Equal(t, 14, actual.Preferences.FontSize) // Original value preserved
			},
		},
		{
			name:    "invalid source",
			src:     make(chan int),
			dst:     &map[string]string{},
			wantErr: true,
		},
		{
			name:    "non-pointer destination",
			src:     map[string]string{"name": "John"},
			dst:     map[string]string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Copy(tt.dst, tt.src)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			tt.check(t, tt.dst)
		})
	}
}

func TestCopyFromStructToProto(t *testing.T) {
	initialJSON := []byte(`{
		"emailCreation": {
			"email": "user@example.com",
			"password": "initial-password",
			"input": {
				"givenName": "John",
				"familyName": "Doe",
				"dataConsent": true
			}
		}
	}`)

	var createRequest *gen.CreateIdentityRequest
	err := Unmarshal(initialJSON, &createRequest)
	require.NoError(t, err)

	require.NotNil(t, createRequest.GetEmailCreation())
	emailCreation := createRequest.GetEmailCreation()
	require.Equal(t, "user@example.com", emailCreation.Email)
	require.Equal(t, "initial-password", emailCreation.Password)
	require.Equal(t, "John", emailCreation.Input.GetGivenName())
	require.Equal(t, "Doe", emailCreation.Input.GetFamilyName())
	require.Equal(t, true, emailCreation.Input.GetDataConsent())

	patchJSON := []byte(`{
		"emailCreation": {
			"email": "changed@example.com",
			"password": "initial-password",
			"input": {
				"familyName": "Changed",
				"unsupportedField": "unsupportedValue"
			}
		}
	}`)

	var patch *structpb.Struct
	err = Unmarshal(patchJSON, &patch)
	require.NoError(t, err)

	err = Copy(createRequest, patch)
	require.NoError(t, err)

	require.Equal(t, "changed@example.com", createRequest.GetEmailCreation().Email)
	require.Equal(t, "initial-password", createRequest.GetEmailCreation().Password)
	require.Equal(t, "John", createRequest.GetEmailCreation().Input.GetGivenName())
	require.Equal(t, "Changed", createRequest.GetEmailCreation().Input.GetFamilyName())
	require.Equal(t, true, createRequest.GetEmailCreation().Input.GetDataConsent())
}

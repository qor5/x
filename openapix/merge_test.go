package openapix

import (
	"context"
	"embed"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

// Test data
var (
	iamServiceSpec = []byte(`
openapi: 3.0.4
info:
  title: IAM Service API
  version: 1.0.0
  x-iam-service:
    code: iam
    name: IAM Service
    description: Identity and Access Management

paths:
  /users:
    get:
      operationId: listUsers
      summary: List users
      responses:
        "200":
          description: Success
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/UserList'
  /users/{id}:
    parameters:
      - name: id
        in: path
        required: true
        schema:
          type: string
    get:
      operationId: getUser
      summary: Get user by ID
      responses:
        "200":
          description: Success

components:
  schemas:
    User:
      type: object
      properties:
        id:
          type: string
        email:
          type: string
    UserList:
      type: object
      properties:
        items:
          type: array
          items:
            $ref: '#/components/schemas/User'
`)

	mediaLibrarySpec = []byte(`
openapi: 3.0.4
info:
  title: Media Library API
  version: 1.0.0

paths:
  /media:
    get:
      operationId: listMedia
      summary: List media
      responses:
        "200":
          description: Success
    post:
      operationId: uploadMedia
      summary: Upload media
      responses:
        "201":
          description: Created
  /media/{id}:
    parameters:
      - name: id
        in: path
        required: true
        schema:
          type: string
    get:
      operationId: getMedia
      summary: Get media by ID
      responses:
        "200":
          description: Success
    delete:
      operationId: deleteMedia
      summary: Delete media
      responses:
        "204":
          description: Deleted

components:
  schemas:
    MediaItem:
      type: object
      properties:
        id:
          type: string
        url:
          type: string
        mimeType:
          type: string
    MediaList:
      type: object
      properties:
        items:
          type: array
          items:
            $ref: '#/components/schemas/MediaItem'
`)

	notificationSpec = []byte(`
openapi: 3.0.4
info:
  title: Notification API
  version: 1.0.0

paths:
  /notifications:
    get:
      operationId: listNotifications
      summary: List notifications
      responses:
        "200":
          description: Success
  /notifications/{id}/read:
    parameters:
      - name: id
        in: path
        required: true
        schema:
          type: string
    post:
      operationId: markAsRead
      summary: Mark notification as read
      responses:
        "200":
          description: Success

components:
  schemas:
    Notification:
      type: object
      properties:
        id:
          type: string
        message:
          type: string
        read:
          type: boolean
`)
)

// =============================================================================
// Scenario 1: Basic merge - merge all paths and components from multiple sources
// =============================================================================

func TestMergeOpenAPISpecs_BasicMergeAll(t *testing.T) {
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{Data: mediaLibrarySpec},
		{Data: notificationSpec},
	})
	require.NoError(t, err)

	// First source's info is preserved
	require.Equal(t, "IAM Service API", spec.Info.Title)

	// All paths from all sources
	require.NotNil(t, spec.Paths.Find("/users"))
	require.NotNil(t, spec.Paths.Find("/users/{id}"))
	require.NotNil(t, spec.Paths.Find("/media"))
	require.NotNil(t, spec.Paths.Find("/media/{id}"))
	require.NotNil(t, spec.Paths.Find("/notifications"))
	require.NotNil(t, spec.Paths.Find("/notifications/{id}/read"))

	// All schemas from all sources
	require.NotNil(t, spec.Components.Schemas["User"])
	require.NotNil(t, spec.Components.Schemas["UserList"])
	require.NotNil(t, spec.Components.Schemas["MediaItem"])
	require.NotNil(t, spec.Components.Schemas["MediaList"])
	require.NotNil(t, spec.Components.Schemas["Notification"])
}

// =============================================================================
// Scenario 2: PathHandler - skip specific paths
// =============================================================================

func TestMergeOpenAPISpecs_PathHandler(t *testing.T) {
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{
			Data: mediaLibrarySpec,
			PathHandler: func(ctx context.Context, info MergePathInfo) error {
				if info.Name == "/media" {
					return nil // Include /media
				}
				return ErrMergeSkip // Skip others
			},
		},
	})
	require.NoError(t, err)

	// IAM paths preserved
	require.NotNil(t, spec.Paths.Find("/users"))
	require.NotNil(t, spec.Paths.Find("/users/{id}"))

	// Only /media included, not /media/{id}
	require.NotNil(t, spec.Paths.Find("/media"))
	require.Nil(t, spec.Paths.Find("/media/{id}"))

	// All schemas still included (no component handler)
	require.NotNil(t, spec.Components.Schemas["MediaItem"])
	require.NotNil(t, spec.Components.Schemas["MediaList"])
}

// =============================================================================
// Scenario 3: ComponentHandler - skip specific components
// =============================================================================

func TestMergeOpenAPISpecs_ComponentHandler(t *testing.T) {
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{
			Data: mediaLibrarySpec,
			ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
				if info.Name == "MediaItem" {
					return nil // Include MediaItem
				}
				return ErrMergeSkip // Skip others
			},
		},
	})
	require.NoError(t, err)

	// All paths included
	require.NotNil(t, spec.Paths.Find("/media"))
	require.NotNil(t, spec.Paths.Find("/media/{id}"))

	// Only MediaItem included
	require.NotNil(t, spec.Components.Schemas["MediaItem"])
	require.Nil(t, spec.Components.Schemas["MediaList"])
}

// =============================================================================
// Scenario 4: Combined handlers - both path and component handlers
// =============================================================================

func TestMergeOpenAPISpecs_CombinedHandlers(t *testing.T) {
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{
			Data: mediaLibrarySpec,
			PathHandler: func(ctx context.Context, info MergePathInfo) error {
				if info.Name == "/media" {
					return nil
				}
				return ErrMergeSkip
			},
			ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
				if info.Name == "MediaItem" {
					return nil
				}
				return ErrMergeSkip
			},
		},
	})
	require.NoError(t, err)

	// Only /media path
	require.NotNil(t, spec.Paths.Find("/media"))
	require.Nil(t, spec.Paths.Find("/media/{id}"))

	// Only MediaItem schema
	require.NotNil(t, spec.Components.Schemas["MediaItem"])
	require.Nil(t, spec.Components.Schemas["MediaList"])
}

// =============================================================================
// Scenario 5: Custom handler with prefix matching
// =============================================================================

func TestMergeOpenAPISpecs_CustomPrefixHandler(t *testing.T) {
	mediaPathHandler := func(ctx context.Context, info MergePathInfo) error {
		if strings.HasPrefix(info.Name, "/media") {
			return nil
		}
		return ErrMergeSkip
	}
	mediaComponentHandler := func(ctx context.Context, info MergeComponentInfo) error {
		if strings.Contains(info.Name, "Media") {
			return nil
		}
		return ErrMergeSkip
	}

	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{
			Data:             mediaLibrarySpec,
			PathHandler:      mediaPathHandler,
			ComponentHandler: mediaComponentHandler,
		},
		{
			Data:             notificationSpec,
			PathHandler:      mediaPathHandler,      // Won't match
			ComponentHandler: mediaComponentHandler, // Won't match
		},
	})
	require.NoError(t, err)

	// Media paths included
	require.NotNil(t, spec.Paths.Find("/media"))
	require.NotNil(t, spec.Paths.Find("/media/{id}"))

	// Notification paths excluded
	require.Nil(t, spec.Paths.Find("/notifications"))

	// Media schemas included
	require.NotNil(t, spec.Components.Schemas["MediaItem"])
	require.NotNil(t, spec.Components.Schemas["MediaList"])

	// Notification schema excluded
	require.Nil(t, spec.Components.Schemas["Notification"])
}

// =============================================================================
// Scenario 6: Default behavior - later source overwrites earlier (no handler)
// =============================================================================

func TestMergeOpenAPISpecs_DefaultOverwrite(t *testing.T) {
	specA := []byte(`
openapi: 3.0.4
info:
  title: Service A
  version: 1.0.0
paths: {}
components:
  schemas:
    Item:
      type: object
      properties:
        name:
          type: string
          description: "From Service A"
`)

	specB := []byte(`
openapi: 3.0.4
info:
  title: Service B
  version: 1.0.0
paths: {}
components:
  schemas:
    Item:
      type: object
      properties:
        name:
          type: string
          description: "From Service B"
`)

	spec, err := Merge(context.Background(), []MergeSource{
		{Data: specA},
		{Data: specB},
	})
	require.NoError(t, err)

	// Service B's Item overwrites Service A's Item
	item := spec.Components.Schemas["Item"]
	require.NotNil(t, item)
	require.Equal(t, "From Service B", item.Value.Properties["name"].Value.Description)
}

// =============================================================================
// Scenario 7: Handler returns error on conflict
// =============================================================================

func TestMergeOpenAPISpecs_ConflictError(t *testing.T) {
	specA := []byte(`
openapi: 3.0.4
info:
  title: Service A
  version: 1.0.0
components:
  schemas:
    Item:
      type: object
`)

	specB := []byte(`
openapi: 3.0.4
info:
  title: Service B
  version: 1.0.0
components:
  schemas:
    Item:
      type: object
`)

	_, err := Merge(context.Background(), []MergeSource{
		{Data: specA},
		{
			Data: specB,
			ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
				if info.Target != nil {
					return errors.Errorf("conflict: %s %s already exists", info.Type, info.Name)
				}
				return nil
			},
		},
	})

	require.Error(t, err)
	require.Contains(t, err.Error(), "conflict: SCHEMA Item already exists")
}

// =============================================================================
// Scenario 8: Handler returns error on path conflict
// =============================================================================

func TestMergeOpenAPISpecs_PathConflictError(t *testing.T) {
	specA := []byte(`
openapi: 3.0.4
info:
  title: Service A
  version: 1.0.0
paths:
  /items:
    get:
      operationId: listItemsA
      responses:
        "200":
          description: OK
`)

	specB := []byte(`
openapi: 3.0.4
info:
  title: Service B
  version: 1.0.0
paths:
  /items:
    get:
      operationId: listItemsB
      responses:
        "200":
          description: OK
`)

	_, err := Merge(context.Background(), []MergeSource{
		{Data: specA},
		{
			Data: specB,
			PathHandler: func(ctx context.Context, info MergePathInfo) error {
				if info.Target != nil {
					return errors.Errorf("path conflict: %s already exists", info.Name)
				}
				return nil
			},
		},
	})

	require.Error(t, err)
	require.Contains(t, err.Error(), "path conflict: /items already exists")
}

// =============================================================================
// Scenario 9: Handler on first source (base spec)
// =============================================================================

func TestMergeOpenAPISpecs_HandlerFirstSource(t *testing.T) {
	spec, err := Merge(context.Background(), []MergeSource{
		{
			Data: iamServiceSpec,
			PathHandler: func(ctx context.Context, info MergePathInfo) error {
				if info.Name == "/users" {
					return nil // Only include /users from base
				}
				return ErrMergeSkip
			},
		},
		{Data: mediaLibrarySpec},
	})
	require.NoError(t, err)

	// Only /users from first source
	require.NotNil(t, spec.Paths.Find("/users"))
	require.Nil(t, spec.Paths.Find("/users/{id}"))

	// All paths from second source
	require.NotNil(t, spec.Paths.Find("/media"))
	require.NotNil(t, spec.Paths.Find("/media/{id}"))
}

// =============================================================================
// Scenario 10: Empty sources
// =============================================================================

func TestMergeOpenAPISpecs_EmptySources(t *testing.T) {
	_, err := Merge(context.Background(), []MergeSource{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "at least one source is required")
}

// =============================================================================
// Scenario 11: Single source (no merge needed)
// =============================================================================

func TestMergeOpenAPISpecs_SingleSource(t *testing.T) {
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
	})
	require.NoError(t, err)

	require.Equal(t, "IAM Service API", spec.Info.Title)
	require.NotNil(t, spec.Paths.Find("/users"))
	require.NotNil(t, spec.Components.Schemas["User"])
}

// =============================================================================
// Scenario 12: Using custom Loader with embed.FS
// =============================================================================

//go:embed testdata/*
var testEmbedFS embed.FS

func TestMergeOpenAPISpecs_WithCustomLoader(t *testing.T) {
	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, uri *url.URL) ([]byte, error) {
		filePath := path.Clean(uri.Path)
		if len(filePath) > 0 && filePath[0] == '/' {
			filePath = filePath[1:]
		}
		return testEmbedFS.ReadFile("testdata/" + filePath)
	}

	mainData, err := testEmbedFS.ReadFile("testdata/openapi_with_refs.yaml")
	require.NoError(t, err)

	spec, err := Merge(context.Background(), []MergeSource{
		{
			Data:   mainData,
			Loader: loader,
		},
	})
	require.NoError(t, err)

	// Verify the $ref was resolved
	require.NotNil(t, spec.Paths.Find("/media"))
	require.NotNil(t, spec.Paths.Find("/media/{id}"))
}

// =============================================================================
// Scenario 13: Main spec references source's component, then merge source
// =============================================================================

func TestMergeOpenAPISpecs_MainReferencesSourceComponent(t *testing.T) {
	mainWithRef := []byte(`
openapi: 3.0.4
info:
  title: Main API
  version: 1.0.0

paths:
  /gallery:
    get:
      operationId: listGallery
      summary: List gallery items
      responses:
        "200":
          description: Success
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/MediaItem'

components:
  schemas:
    MediaItem:
      type: object
      description: "Placeholder - will be overwritten by media library"
`)

	spec, err := Merge(context.Background(), []MergeSource{
		{Data: mainWithRef},
		{Data: mediaLibrarySpec}, // MediaItem from here overwrites placeholder
	})
	require.NoError(t, err)

	// Main's /gallery path preserved
	require.NotNil(t, spec.Paths.Find("/gallery"))

	// Media library's paths merged
	require.NotNil(t, spec.Paths.Find("/media"))
	require.NotNil(t, spec.Paths.Find("/media/{id}"))

	// MediaItem from media library (has url and mimeType properties)
	mediaItem := spec.Components.Schemas["MediaItem"]
	require.NotNil(t, mediaItem)
	require.NotNil(t, mediaItem.Value.Properties["url"])
	require.NotNil(t, mediaItem.Value.Properties["mimeType"])
}

// =============================================================================
// Scenario 14: Marshal merged spec to YAML
// =============================================================================

func TestMarshal(t *testing.T) {
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{Data: mediaLibrarySpec},
	})
	require.NoError(t, err)

	data, err := Marshal(spec)
	require.NoError(t, err)
	require.NotEmpty(t, data)

	yamlStr := string(data)
	require.Contains(t, yamlStr, "openapi:")
	require.Contains(t, yamlStr, "/users")
	require.Contains(t, yamlStr, "/media")
	require.Contains(t, yamlStr, "MediaItem")
	require.Contains(t, yamlStr, "User")

	// Optional: write to file for inspection
	if os.Getenv("WRITE_TEST_OUTPUT") == "1" {
		err = os.WriteFile("testdata/merged_spec.yaml", data, 0o644)
		require.NoError(t, err)
	}
}

// =============================================================================
// Scenario 15: Invalid YAML source
// =============================================================================

func TestMergeOpenAPISpecs_InvalidYAML(t *testing.T) {
	invalidYAML := []byte(`
this is not valid yaml: [
`)

	_, err := Merge(context.Background(), []MergeSource{
		{Data: invalidYAML},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "failed to load source 0")
}

// =============================================================================
// Scenario 16: All component types - merge and handler verification
// =============================================================================

func TestMergeOpenAPISpecs_AllComponentTypes(t *testing.T) {
	specA := []byte(`
openapi: 3.0.4
info:
  title: Service A
  version: 1.0.0
paths: {}
components:
  schemas:
    SchemaA:
      type: object
  responses:
    ResponseA:
      description: Response A
  parameters:
    ParamA:
      name: paramA
      in: query
      schema:
        type: string
  examples:
    ExampleA:
      value: "example A"
  requestBodies:
    BodyA:
      content:
        application/json:
          schema:
            type: object
  headers:
    HeaderA:
      schema:
        type: string
  securitySchemes:
    AuthA:
      type: http
      scheme: bearer
  links:
    LinkA:
      operationId: opA
  callbacks:
    CallbackA:
      '{$request.body#/url}':
        post:
          responses:
            "200":
              description: OK
`)

	specB := []byte(`
openapi: 3.0.4
info:
  title: Service B
  version: 1.0.0
paths: {}
components:
  schemas:
    SchemaB:
      type: object
  responses:
    ResponseB:
      description: Response B
  parameters:
    ParamB:
      name: paramB
      in: query
      schema:
        type: string
  examples:
    ExampleB:
      value: "example B"
  requestBodies:
    BodyB:
      content:
        application/json:
          schema:
            type: object
  headers:
    HeaderB:
      schema:
        type: string
  securitySchemes:
    AuthB:
      type: apiKey
      in: header
      name: X-API-Key
  links:
    LinkB:
      operationId: opB
  callbacks:
    CallbackB:
      '{$request.body#/url}':
        post:
          responses:
            "200":
              description: OK
`)

	// Test merge of all component types
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: specA},
		{Data: specB},
	})
	require.NoError(t, err)

	// Verify all components from both sources are merged
	require.NotNil(t, spec.Components.Schemas["SchemaA"])
	require.NotNil(t, spec.Components.Schemas["SchemaB"])
	require.NotNil(t, spec.Components.Responses["ResponseA"])
	require.NotNil(t, spec.Components.Responses["ResponseB"])
	require.NotNil(t, spec.Components.Parameters["ParamA"])
	require.NotNil(t, spec.Components.Parameters["ParamB"])
	require.NotNil(t, spec.Components.Examples["ExampleA"])
	require.NotNil(t, spec.Components.Examples["ExampleB"])
	require.NotNil(t, spec.Components.RequestBodies["BodyA"])
	require.NotNil(t, spec.Components.RequestBodies["BodyB"])
	require.NotNil(t, spec.Components.Headers["HeaderA"])
	require.NotNil(t, spec.Components.Headers["HeaderB"])
	require.NotNil(t, spec.Components.SecuritySchemes["AuthA"])
	require.NotNil(t, spec.Components.SecuritySchemes["AuthB"])
	require.NotNil(t, spec.Components.Links["LinkA"])
	require.NotNil(t, spec.Components.Links["LinkB"])
	require.NotNil(t, spec.Components.Callbacks["CallbackA"])
	require.NotNil(t, spec.Components.Callbacks["CallbackB"])

	// Test ComponentType enum values via handler
	collectedTypes := make(map[ComponentType][]string)
	_, err = Merge(context.Background(), []MergeSource{
		{
			Data: specA,
			ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
				collectedTypes[info.Type] = append(collectedTypes[info.Type], info.Name)
				return nil
			},
		},
	})
	require.NoError(t, err)

	require.Contains(t, collectedTypes[ComponentTypeSchema], "SchemaA")
	require.Contains(t, collectedTypes[ComponentTypeResponse], "ResponseA")
	require.Contains(t, collectedTypes[ComponentTypeParameter], "ParamA")
	require.Contains(t, collectedTypes[ComponentTypeExample], "ExampleA")
	require.Contains(t, collectedTypes[ComponentTypeRequestBody], "BodyA")
	require.Contains(t, collectedTypes[ComponentTypeHeader], "HeaderA")
	require.Contains(t, collectedTypes[ComponentTypeSecurityScheme], "AuthA")
	require.Contains(t, collectedTypes[ComponentTypeLink], "LinkA")
	require.Contains(t, collectedTypes[ComponentTypeCallback], "CallbackA")
}

// =============================================================================
// Scenario 17: Context propagation and first source handlers
// =============================================================================

func TestMergeOpenAPISpecs_ContextAndFirstSourceHandlers(t *testing.T) {
	type ctxKey string
	const testKey ctxKey = "test-key"

	ctx := context.WithValue(context.Background(), testKey, "test-value")

	var pathCtxValue, componentCtxValue any

	spec, err := Merge(ctx, []MergeSource{
		{
			Data: iamServiceSpec,
			PathHandler: func(ctx context.Context, info MergePathInfo) error {
				pathCtxValue = ctx.Value(testKey)
				if info.Name == "/users" {
					return nil
				}
				return ErrMergeSkip
			},
			ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
				componentCtxValue = ctx.Value(testKey)
				if info.Name == "User" {
					return nil
				}
				return ErrMergeSkip
			},
		},
		{Data: mediaLibrarySpec},
	})
	require.NoError(t, err)

	// Context propagation
	require.Equal(t, "test-value", pathCtxValue)
	require.Equal(t, "test-value", componentCtxValue)

	// First source filtering
	require.NotNil(t, spec.Paths.Find("/users"))
	require.Nil(t, spec.Paths.Find("/users/{id}"))
	require.NotNil(t, spec.Components.Schemas["User"])
	require.Nil(t, spec.Components.Schemas["UserList"])

	// Second source merged normally
	require.NotNil(t, spec.Paths.Find("/media"))
	require.NotNil(t, spec.Components.Schemas["MediaItem"])
}

// =============================================================================
// Scenario 18: Edge cases - empty paths/components
// =============================================================================

func TestMergeOpenAPISpecs_EdgeCases(t *testing.T) {
	specNoPaths := []byte(`
openapi: 3.0.4
info:
  title: No Paths API
  version: 1.0.0
paths: {}
components:
  schemas:
    SharedModel:
      type: object
`)

	specNoComponents := []byte(`
openapi: 3.0.4
info:
  title: No Components API
  version: 1.0.0
paths:
  /health:
    get:
      operationId: healthCheck
      responses:
        "200":
          description: OK
`)

	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{Data: specNoPaths},
		{Data: specNoComponents},
	})
	require.NoError(t, err)

	require.NotNil(t, spec.Paths.Find("/users"))
	require.NotNil(t, spec.Paths.Find("/health"))
	require.NotNil(t, spec.Components.Schemas["User"])
	require.NotNil(t, spec.Components.Schemas["SharedModel"])
}

// =============================================================================
// Scenario 22: Target verification in handlers
// =============================================================================

func TestMergeOpenAPISpecs_TargetInHandler(t *testing.T) {
	specA := []byte(`
openapi: 3.0.4
info:
  title: Service A
  version: 1.0.0
paths:
  /items:
    get:
      operationId: listItems
      responses:
        "200":
          description: OK
components:
  schemas:
    Item:
      type: object
      description: "Item from A"
`)

	specB := []byte(`
openapi: 3.0.4
info:
  title: Service B
  version: 1.0.0
paths:
  /items:
    post:
      operationId: createItem
      responses:
        "201":
          description: Created
  /new-path:
    get:
      operationId: newPath
      responses:
        "200":
          description: OK
components:
  schemas:
    Item:
      type: object
      description: "Item from B"
    NewSchema:
      type: object
`)

	var pathTargetItems []string
	var pathNewItems []string
	var componentTargetItems []string
	var componentNewItems []string

	_, err := Merge(context.Background(), []MergeSource{
		{Data: specA},
		{
			Data: specB,
			PathHandler: func(ctx context.Context, info MergePathInfo) error {
				if info.Target != nil {
					pathTargetItems = append(pathTargetItems, info.Name)
				} else {
					pathNewItems = append(pathNewItems, info.Name)
				}
				return nil
			},
			ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
				if info.Target != nil {
					componentTargetItems = append(componentTargetItems, info.Name)
				} else {
					componentNewItems = append(componentNewItems, info.Name)
				}
				return nil
			},
		},
	})
	require.NoError(t, err)

	// /items exists in specA, /new-path is new
	require.Contains(t, pathTargetItems, "/items")
	require.Contains(t, pathNewItems, "/new-path")

	// Item exists in specA, NewSchema is new
	require.Contains(t, componentTargetItems, "Item")
	require.NotContains(t, componentTargetItems, "NewSchema")
	require.Contains(t, componentNewItems, "NewSchema")
	require.NotContains(t, componentNewItems, "Item")
}

// =============================================================================
// Scenario 19: Error handling - validation failure and handler errors
// =============================================================================

func TestMergeOpenAPISpecs_ErrorHandling(t *testing.T) {
	// Validation failure
	invalidSpec := []byte(`
openapi: 3.0.4
info:
  title: Invalid API
  version: 1.0.0
paths:
  /items/{id}:
    get:
      operationId: getItem
      responses:
        "200":
          description: OK
`)
	_, err := Merge(context.Background(), []MergeSource{{Data: invalidSpec}})
	require.Error(t, err)
	require.Contains(t, err.Error(), "merged spec validation failed")

	// Path handler error
	_, err = Merge(context.Background(), []MergeSource{{
		Data: iamServiceSpec,
		PathHandler: func(ctx context.Context, info MergePathInfo) error {
			return errors.New("path handler error")
		},
	}})
	require.Error(t, err)
	require.Contains(t, err.Error(), "path handler error")

	// Component handler error
	_, err = Merge(context.Background(), []MergeSource{{
		Data: iamServiceSpec,
		ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
			return errors.New("component handler error")
		},
	}})
	require.Error(t, err)
	require.Contains(t, err.Error(), "component handler error")
}

// =============================================================================
// Scenario 20: Skip all paths/components from a source
// =============================================================================

func TestMergeOpenAPISpecs_SkipAll(t *testing.T) {
	// Skip all paths
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{
			Data: mediaLibrarySpec,
			PathHandler: func(ctx context.Context, info MergePathInfo) error {
				return ErrMergeSkip
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, spec.Paths.Find("/users"))
	require.Nil(t, spec.Paths.Find("/media"))
	require.NotNil(t, spec.Components.Schemas["MediaItem"]) // Components still merged

	// Skip all components
	spec, err = Merge(context.Background(), []MergeSource{
		{Data: iamServiceSpec},
		{
			Data: mediaLibrarySpec,
			ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
				return ErrMergeSkip
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, spec.Components.Schemas["User"])
	require.Nil(t, spec.Components.Schemas["MediaItem"])
	require.NotNil(t, spec.Paths.Find("/media")) // Paths still merged
}

// =============================================================================
// Scenario 21: Filter by ComponentType (first source delete logic)
// =============================================================================

func TestMergeOpenAPISpecs_FirstSourceDeleteAllComponentTypes(t *testing.T) {
	specWithAllComponents := []byte(`
openapi: 3.0.4
info:
  title: All Components API
  version: 1.0.0
paths: {}
components:
  schemas:
    TestSchema:
      type: object
  responses:
    TestResponse:
      description: Test
  parameters:
    TestParam:
      name: test
      in: query
      schema:
        type: string
  examples:
    TestExample:
      value: "test"
  requestBodies:
    TestBody:
      content:
        application/json:
          schema:
            type: object
  headers:
    TestHeader:
      schema:
        type: string
  securitySchemes:
    TestAuth:
      type: http
      scheme: bearer
  links:
    TestLink:
      operationId: test
  callbacks:
    TestCallback:
      '{$request.body#/url}':
        post:
          responses:
            "200":
              description: OK
`)

	// Test that all component types can be deleted via first source handler
	spec, err := Merge(context.Background(), []MergeSource{
		{
			Data: specWithAllComponents,
			ComponentHandler: func(ctx context.Context, info MergeComponentInfo) error {
				// Skip all components - they should be deleted
				return ErrMergeSkip
			},
		},
	})
	require.NoError(t, err)

	// All components should be deleted
	require.Nil(t, spec.Components.Schemas["TestSchema"])
	require.Nil(t, spec.Components.Responses["TestResponse"])
	require.Nil(t, spec.Components.Parameters["TestParam"])
	require.Nil(t, spec.Components.Examples["TestExample"])
	require.Nil(t, spec.Components.RequestBodies["TestBody"])
	require.Nil(t, spec.Components.Headers["TestHeader"])
	require.Nil(t, spec.Components.SecuritySchemes["TestAuth"])
	require.Nil(t, spec.Components.Links["TestLink"])
	require.Nil(t, spec.Components.Callbacks["TestCallback"])
}

// =============================================================================
// Scenario 22: Extensions merger for x-iam-permissions
// =============================================================================

func TestMergeOpenAPISpecs_ExtensionsHandler(t *testing.T) {
	// Main spec (e.g., pim.yaml)
	pimSpec := []byte(`
openapi: 3.0.4
info:
  title: PIM Service API
  version: 1.0.0
  x-iam-service:
    code: pim
    name: PIM Service

x-iam-permissions:
  - code: "pim:product:read"
    name: "View Products"
    type: feature
  - code: "pim:product:write"
    name: "Edit Products"
    type: feature

paths:
  /products:
    get:
      operationId: listProducts
      x-iam-access: authorized
      x-required-permission: pim:product:read
      responses:
        "200":
          description: OK
`)

	// Shared module spec (e.g., media_library.yaml) - already prefixed
	mediaLibrarySpec := []byte(`
openapi: 3.0.4
info:
  title: Media Library API
  version: 1.0.0

x-iam-permissions:
  - code: "pim:media:read"
    name: "View Media"
    type: feature
  - code: "pim:media:write"
    name: "Upload Media"
    type: feature

paths:
  /media:
    get:
      operationId: listMedia
      x-iam-access: authorized
      x-required-permission: pim:media:read
      responses:
        "200":
          description: OK
`)

	// Custom extensions handler that appends x-iam-permissions
	permissionsHandler := func(ctx context.Context, info MergeExtensionsInfo) error {
		sourcePerms, ok := info.Source["x-iam-permissions"]
		if !ok {
			return nil
		}
		sourceList, ok := sourcePerms.([]any)
		if !ok {
			return nil
		}

		targetPerms, exists := info.Target["x-iam-permissions"]
		if !exists {
			info.Target["x-iam-permissions"] = sourceList
			return nil
		}

		targetList, ok := targetPerms.([]any)
		if !ok {
			info.Target["x-iam-permissions"] = sourceList
			return nil
		}

		info.Target["x-iam-permissions"] = append(targetList, sourceList...)
		return nil
	}

	spec, err := Merge(context.Background(), []MergeSource{
		{Data: pimSpec},
		{
			Data:              mediaLibrarySpec,
			ExtensionsHandler: permissionsHandler,
		},
	})
	require.NoError(t, err)

	// Verify paths merged
	require.NotNil(t, spec.Paths.Find("/products"))
	require.NotNil(t, spec.Paths.Find("/media"))

	// Verify extensions merged
	perms, ok := spec.Extensions["x-iam-permissions"].([]any)
	require.True(t, ok)
	require.Len(t, perms, 4) // 2 from pim + 2 from media_library

	// Verify permission codes
	codes := make([]string, 0, len(perms))
	for _, p := range perms {
		if pm, ok := p.(map[string]any); ok {
			if code, ok := pm["code"].(string); ok {
				codes = append(codes, code)
			}
		}
	}
	require.Contains(t, codes, "pim:product:read")
	require.Contains(t, codes, "pim:product:write")
	require.Contains(t, codes, "pim:media:read")
	require.Contains(t, codes, "pim:media:write")
}

// =============================================================================
// Scenario 23: Extensions merger with prefix transformation
// =============================================================================

func TestMergeOpenAPISpecs_ExtensionsHandlerWithPrefixTransform(t *testing.T) {
	// This test demonstrates the recommended workflow:
	// 1. Load media_library.yaml
	// 2. Transform permission codes and path extensions with prefix
	// 3. Merge into main spec

	// Main spec (cms.yaml)
	cmsSpec := []byte(`
openapi: 3.0.4
info:
  title: CMS Service API
  version: 1.0.0
  x-iam-service:
    code: cms
    name: CMS Service

x-iam-permissions:
  - code: "cms:page:read"
    name: "View Pages"
    type: feature

paths:
  /pages:
    get:
      operationId: listPages
      x-iam-access: authorized
      x-required-permission: cms:page:read
      responses:
        "200":
          description: OK
`)

	// Shared module (media_library.yaml) with generic codes
	mediaLibraryGeneric := []byte(`
openapi: 3.0.4
info:
  title: Media Library API
  version: 1.0.0

x-iam-permissions:
  - code: "media:read"
    name: "View Media"
    type: feature
  - code: "media:write"
    name: "Upload Media"
    type: feature

paths:
  /media:
    get:
      operationId: listMedia
      x-iam-access: authorized
      x-required-permission: media:read
      responses:
        "200":
          description: OK
    post:
      operationId: uploadMedia
      x-iam-access: authorized
      x-required-permission: media:write
      responses:
        "201":
          description: Created
`)

	// Helper to add prefix to permission codes
	addPrefix := func(prefix string) MergeExtensionsHandler {
		return func(ctx context.Context, info MergeExtensionsInfo) error {
			sourcePerms, ok := info.Source["x-iam-permissions"]
			if !ok {
				return nil
			}
			sourceList, ok := sourcePerms.([]any)
			if !ok {
				return nil
			}

			// Transform permission codes with prefix
			for _, p := range sourceList {
				if pm, ok := p.(map[string]any); ok {
					if code, ok := pm["code"].(string); ok {
						pm["code"] = prefix + ":" + code
					}
				}
			}

			// Merge into target
			targetPerms, exists := info.Target["x-iam-permissions"]
			if !exists {
				info.Target["x-iam-permissions"] = sourceList
				return nil
			}
			targetList, ok := targetPerms.([]any)
			if !ok {
				info.Target["x-iam-permissions"] = sourceList
				return nil
			}
			info.Target["x-iam-permissions"] = append(targetList, sourceList...)
			return nil
		}
	}

	// PathHandler to transform x-required-permission in paths
	addPathPrefix := func(prefix string) MergePathHandler {
		return func(ctx context.Context, info MergePathInfo) error {
			for _, op := range info.Source.Operations() {
				if op.Extensions != nil {
					if perm, ok := op.Extensions["x-required-permission"].(string); ok {
						op.Extensions["x-required-permission"] = prefix + ":" + perm
					}
				}
			}
			return nil
		}
	}

	spec, err := Merge(context.Background(), []MergeSource{
		{Data: cmsSpec},
		{
			Data:              mediaLibraryGeneric,
			PathHandler:       addPathPrefix("cms"),
			ExtensionsHandler: addPrefix("cms"),
		},
	})
	require.NoError(t, err)

	// Verify paths merged
	require.NotNil(t, spec.Paths.Find("/pages"))
	require.NotNil(t, spec.Paths.Find("/media"))

	// Verify path extensions transformed
	mediaPath := spec.Paths.Find("/media")
	require.Equal(t, "cms:media:read", mediaPath.Get.Extensions["x-required-permission"])
	require.Equal(t, "cms:media:write", mediaPath.Post.Extensions["x-required-permission"])

	// Verify spec extensions merged with prefix
	perms, ok := spec.Extensions["x-iam-permissions"].([]any)
	require.True(t, ok)
	require.Len(t, perms, 3) // 1 from cms + 2 from media_library

	codes := make([]string, 0, len(perms))
	for _, p := range perms {
		if pm, ok := p.(map[string]any); ok {
			if code, ok := pm["code"].(string); ok {
				codes = append(codes, code)
			}
		}
	}
	require.Contains(t, codes, "cms:page:read")
	require.Contains(t, codes, "cms:media:read")
	require.Contains(t, codes, "cms:media:write")
}

// =============================================================================
// Scenario 24: Extensions merger error propagation
// =============================================================================

func TestMergeOpenAPISpecs_ExtensionsHandlerError(t *testing.T) {
	specA := []byte(`
openapi: 3.0.4
info:
  title: Service A
  version: 1.0.0
paths: {}
x-custom: "value"
`)

	specB := []byte(`
openapi: 3.0.4
info:
  title: Service B
  version: 1.0.0
paths: {}
x-custom: "another"
`)

	_, err := Merge(context.Background(), []MergeSource{
		{Data: specA},
		{
			Data: specB,
			ExtensionsHandler: func(ctx context.Context, info MergeExtensionsInfo) error {
				return errors.New("extensions handler failed")
			},
		},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "extensions handler failed")
}

// =============================================================================
// Scenario 25: Extensions handler behavior
// =============================================================================

func TestMergeOpenAPISpecs_ExtensionsHandlerBehavior(t *testing.T) {
	specA := []byte(`
openapi: 3.0.4
info:
  title: Service A
  version: 1.0.0
paths: {}
x-custom-a: "from A"
x-shared: "original"
`)

	specB := []byte(`
openapi: 3.0.4
info:
  title: Service B
  version: 1.0.0
paths: {}
x-custom-b: "from B"
x-shared: "overwritten"
`)

	// Without ExtensionsHandler (nil), no extensions merging is performed
	spec, err := Merge(context.Background(), []MergeSource{
		{Data: specA},
		{Data: specB},
	})
	require.NoError(t, err)

	require.Equal(t, "from A", spec.Extensions["x-custom-a"])
	require.Nil(t, spec.Extensions["x-custom-b"]) // Not merged
	require.Equal(t, "original", spec.Extensions["x-shared"])

	// With ExtensionsHandler that performs override, extensions are merged
	spec, err = Merge(context.Background(), []MergeSource{
		{Data: specA},
		{
			Data: specB,
			ExtensionsHandler: func(ctx context.Context, info MergeExtensionsInfo) error {
				// Handler is fully responsible for merging
				for k, v := range info.Source {
					info.Target[k] = v
				}
				return nil
			},
		},
	})
	require.NoError(t, err)

	require.Equal(t, "from A", spec.Extensions["x-custom-a"])
	require.Equal(t, "from B", spec.Extensions["x-custom-b"])
	require.Equal(t, "overwritten", spec.Extensions["x-shared"])
}

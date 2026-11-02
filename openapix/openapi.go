package openapix

// ComponentType represents the type of OpenAPI component.
type ComponentType string

const (
	ComponentTypeSchema         ComponentType = "SCHEMA"
	ComponentTypeResponse       ComponentType = "RESPONSE"
	ComponentTypeParameter      ComponentType = "PARAMETER"
	ComponentTypeExample        ComponentType = "EXAMPLE"
	ComponentTypeRequestBody    ComponentType = "REQUEST_BODY"
	ComponentTypeHeader         ComponentType = "HEADER"
	ComponentTypeSecurityScheme ComponentType = "SECURITY_SCHEME"
	ComponentTypeLink           ComponentType = "LINK"
	ComponentTypeCallback       ComponentType = "CALLBACK"
)

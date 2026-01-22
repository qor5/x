package openapix

import (
	"context"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// ErrMergeSkip is returned by handlers to indicate the item should be skipped (not merged).
var ErrMergeSkip = errors.New("merge: skip")

// MergePathInfo provides information for path handling decisions during merge.
type MergePathInfo struct {
	Name   string             // The path name (e.g., "/users/{id}")
	Source *openapi3.PathItem // The path item from source spec
	Target *openapi3.PathItem // The existing path item in target spec (nil if not exists)
}

// MergeComponentInfo provides information for component handling decisions during merge.
type MergeComponentInfo struct {
	Name   string        // The component name (e.g., "User", "MediaItem")
	Type   ComponentType // The component type
	Source any           // The component item from source spec
	Target any           // The existing component item in target spec (nil if not exists)
}

// MergeExtensionsInfo provides information for extensions handling decisions during merge.
type MergeExtensionsInfo struct {
	Source map[string]any // Extensions from source spec
	Target map[string]any // Extensions in target spec (never nil, may be empty)
}

// MergePathHandler determines how to handle a path during merge.
// Return nil to merge (overwrite if exists), ErrMergeSkip to skip, or other error to abort.
type MergePathHandler func(ctx context.Context, info *MergePathInfo) error

// MergeComponentHandler determines how to handle a component during merge.
// Return nil to merge (overwrite if exists), ErrMergeSkip to skip, or other error to abort.
type MergeComponentHandler func(ctx context.Context, info *MergeComponentInfo) error

// MergeExtensionsHandler handles extensions merging.
// If nil, no extensions merging is performed.
// If non-nil, the handler is fully responsible for merging info.Source into info.Target.
// The framework does nothing after the handler returns (no default override).
// Return nil on success, or error to abort.
type MergeExtensionsHandler func(ctx context.Context, info *MergeExtensionsInfo) error

// MergeSource represents an OpenAPI spec source to be merged.
type MergeSource struct {
	// Data is the raw YAML/JSON bytes of the OpenAPI spec.
	Data []byte

	// Loader is an optional custom loader for loading the spec.
	// Useful when loading specs from embed.FS or other non-filesystem sources.
	// If nil, a default loader will be used.
	Loader *openapi3.Loader

	// PathHandler handles path merge decisions. If nil, all paths are merged (overwrite on conflict).
	PathHandler MergePathHandler

	// ComponentHandler handles component merge decisions. If nil, all components are merged (overwrite on conflict).
	// This applies to all component types: schemas, responses, parameters, etc.
	ComponentHandler MergeComponentHandler

	// ExtensionsHandler handles spec-level extensions merging (e.g., x-iam-permissions).
	// If nil, no extensions merging is performed.
	// If non-nil, the handler is fully responsible for merging.
	// This is only called for non-first sources.
	ExtensionsHandler MergeExtensionsHandler
}

// Merge merges multiple OpenAPI specs into one.
// The first source is used as the base (info, servers, security, etc. are taken from it).
// Paths and all components from all sources are merged.
//
// Note: OpenAPI has no namespace - names are globally unique.
// Use PathHandler and ComponentHandler to control filtering and conflict behavior.
func Merge(ctx context.Context, sources []*MergeSource) (*openapi3.T, error) {
	if len(sources) == 0 {
		return nil, errors.New("at least one source is required")
	}

	var result *openapi3.T

	for i, source := range sources {
		loader := source.Loader
		if loader == nil {
			loader = openapi3.NewLoader()
		}

		spec, err := loader.LoadFromData(source.Data)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to load source %d", i)
		}

		if result == nil {
			result = spec
			ensureComponents(result)
			if err := handleFirstSource(ctx, result, source.PathHandler, source.ComponentHandler); err != nil {
				return nil, errors.Wrapf(err, "failed to process source %d", i)
			}
		} else {
			if err := mergeIntoSpec(ctx, result, spec, source.PathHandler, source.ComponentHandler, source.ExtensionsHandler); err != nil {
				return nil, errors.Wrapf(err, "failed to merge source %d", i)
			}
		}
	}

	if err := result.Validate(ctx); err != nil {
		return nil, errors.Wrap(err, "merged spec validation failed")
	}

	return result, nil
}

func ensureComponents(spec *openapi3.T) {
	if spec.Components == nil {
		spec.Components = &openapi3.Components{}
	}
	if spec.Components.Schemas == nil {
		spec.Components.Schemas = make(openapi3.Schemas)
	}
	if spec.Components.Responses == nil {
		spec.Components.Responses = make(openapi3.ResponseBodies)
	}
	if spec.Components.Parameters == nil {
		spec.Components.Parameters = make(openapi3.ParametersMap)
	}
	if spec.Components.Examples == nil {
		spec.Components.Examples = make(openapi3.Examples)
	}
	if spec.Components.RequestBodies == nil {
		spec.Components.RequestBodies = make(openapi3.RequestBodies)
	}
	if spec.Components.Headers == nil {
		spec.Components.Headers = make(openapi3.Headers)
	}
	if spec.Components.SecuritySchemes == nil {
		spec.Components.SecuritySchemes = make(openapi3.SecuritySchemes)
	}
	if spec.Components.Links == nil {
		spec.Components.Links = make(openapi3.Links)
	}
	if spec.Components.Callbacks == nil {
		spec.Components.Callbacks = make(openapi3.Callbacks)
	}
}

func handleFirstSource(ctx context.Context, spec *openapi3.T, pathHandler MergePathHandler, componentHandler MergeComponentHandler) error {
	if pathHandler != nil && spec.Paths != nil {
		for name, item := range spec.Paths.Map() {
			err := pathHandler(ctx, &MergePathInfo{Name: name, Source: item, Target: nil})
			if errors.Is(err, ErrMergeSkip) {
				spec.Paths.Delete(name)
			} else if err != nil {
				return errors.Wrapf(err, "path handler error for %s", name)
			}
		}
	}

	if componentHandler != nil && spec.Components != nil {
		if err := processComponents(ctx, spec.Components, nil, componentHandler, true); err != nil {
			return err
		}
	}

	return nil
}

func mergeIntoSpec(ctx context.Context, target, source *openapi3.T, pathHandler MergePathHandler, componentHandler MergeComponentHandler, extensionsHandler MergeExtensionsHandler) error {
	if source.Paths != nil {
		if target.Paths == nil {
			target.Paths = &openapi3.Paths{}
		}
		for name, item := range source.Paths.Map() {
			existingItem := target.Paths.Find(name)
			if pathHandler != nil {
				err := pathHandler(ctx, &MergePathInfo{Name: name, Source: item, Target: existingItem})
				if errors.Is(err, ErrMergeSkip) {
					continue
				} else if err != nil {
					return errors.Wrapf(err, "path handler error for %s", name)
				}
			}
			target.Paths.Set(name, item)
		}
	}

	if source.Components != nil {
		if err := processComponents(ctx, target.Components, source.Components, componentHandler, false); err != nil {
			return err
		}
	}

	if extensionsHandler != nil && len(source.Extensions) > 0 {
		if target.Extensions == nil {
			target.Extensions = make(map[string]any)
		}
		if err := extensionsHandler(ctx, &MergeExtensionsInfo{Source: source.Extensions, Target: target.Extensions}); err != nil {
			return errors.Wrap(err, "extensions handler error")
		}
	}

	return nil
}

// componentEntry represents a single component for processing.
type componentEntry struct {
	name         string
	compType     ComponentType
	item         any
	existingItem any
	merge        func() // Called to perform the merge
	delete       func() // Called to delete from first source
}

// mapGetAny returns the value from map as any, avoiding typed-nil issue.
// When a map returns nil for non-existent key, assigning it to any creates a typed-nil.
// This function returns untyped nil when key doesn't exist.
func mapGetAny[K comparable, V any](m map[K]V, key K) any {
	if v, ok := m[key]; ok {
		return v
	}
	return nil
}

// processComponents handles both first source filtering and merge operations.
// If isFirstSource is true, source should be nil and we filter target in place.
// If isFirstSource is false, we merge source into target.
func processComponents(ctx context.Context, target, source *openapi3.Components, handler MergeComponentHandler, isFirstSource bool) error {
	var entries []componentEntry

	if isFirstSource {
		// First source: filter target in place
		for name, item := range target.Schemas {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeSchema, item: item, existingItem: nil,
				delete: func() { delete(target.Schemas, name) },
			})
		}
		for name, item := range target.Responses {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeResponse, item: item, existingItem: nil,
				delete: func() { delete(target.Responses, name) },
			})
		}
		for name, item := range target.Parameters {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeParameter, item: item, existingItem: nil,
				delete: func() { delete(target.Parameters, name) },
			})
		}
		for name, item := range target.Examples {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeExample, item: item, existingItem: nil,
				delete: func() { delete(target.Examples, name) },
			})
		}
		for name, item := range target.RequestBodies {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeRequestBody, item: item, existingItem: nil,
				delete: func() { delete(target.RequestBodies, name) },
			})
		}
		for name, item := range target.Headers {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeHeader, item: item, existingItem: nil,
				delete: func() { delete(target.Headers, name) },
			})
		}
		for name, item := range target.SecuritySchemes {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeSecurityScheme, item: item, existingItem: nil,
				delete: func() { delete(target.SecuritySchemes, name) },
			})
		}
		for name, item := range target.Links {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeLink, item: item, existingItem: nil,
				delete: func() { delete(target.Links, name) },
			})
		}
		for name, item := range target.Callbacks {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeCallback, item: item, existingItem: nil,
				delete: func() { delete(target.Callbacks, name) },
			})
		}
	} else {
		// Merge: source into target
		for name, item := range source.Schemas {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeSchema, item: item,
				existingItem: mapGetAny(target.Schemas, name),
				merge:        func() { target.Schemas[name] = item },
			})
		}
		for name, item := range source.Responses {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeResponse, item: item,
				existingItem: mapGetAny(target.Responses, name),
				merge:        func() { target.Responses[name] = item },
			})
		}
		for name, item := range source.Parameters {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeParameter, item: item,
				existingItem: mapGetAny(target.Parameters, name),
				merge:        func() { target.Parameters[name] = item },
			})
		}
		for name, item := range source.Examples {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeExample, item: item,
				existingItem: mapGetAny(target.Examples, name),
				merge:        func() { target.Examples[name] = item },
			})
		}
		for name, item := range source.RequestBodies {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeRequestBody, item: item,
				existingItem: mapGetAny(target.RequestBodies, name),
				merge:        func() { target.RequestBodies[name] = item },
			})
		}
		for name, item := range source.Headers {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeHeader, item: item,
				existingItem: mapGetAny(target.Headers, name),
				merge:        func() { target.Headers[name] = item },
			})
		}
		for name, item := range source.SecuritySchemes {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeSecurityScheme, item: item,
				existingItem: mapGetAny(target.SecuritySchemes, name),
				merge:        func() { target.SecuritySchemes[name] = item },
			})
		}
		for name, item := range source.Links {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeLink, item: item,
				existingItem: mapGetAny(target.Links, name),
				merge:        func() { target.Links[name] = item },
			})
		}
		for name, item := range source.Callbacks {
			name, item := name, item
			entries = append(entries, componentEntry{
				name: name, compType: ComponentTypeCallback, item: item,
				existingItem: mapGetAny(target.Callbacks, name),
				merge:        func() { target.Callbacks[name] = item },
			})
		}
	}

	for _, entry := range entries {
		if handler != nil {
			err := handler(ctx, &MergeComponentInfo{
				Name:   entry.name,
				Type:   entry.compType,
				Source: entry.item,
				Target: entry.existingItem,
			})
			if errors.Is(err, ErrMergeSkip) {
				if isFirstSource && entry.delete != nil {
					entry.delete()
				}
				continue
			} else if err != nil {
				return errors.Wrapf(err, "component handler error for %s %s", entry.compType, entry.name)
			}
		}
		if !isFirstSource && entry.merge != nil {
			entry.merge()
		}
	}

	return nil
}

// Marshal writes an OpenAPI spec to YAML bytes.
func Marshal(spec *openapi3.T) ([]byte, error) {
	data, err := yaml.Marshal(spec)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal spec to YAML")
	}
	return data, nil
}

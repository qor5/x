# jsonx

A specialized JSON handling package with enhanced support for Protocol Buffers.

## Features

- **Marshal Method**: When handling `proto.Message` or its pointers (regardless of nesting depth), internally uses `protojson` with `EmitUnpopulated(true)` to ensure consistency across all data consumers.

- **Unmarshal Method**: When handling `proto.Message` or its multi-level pointers, internally uses `protojson` with `DiscardUnknown(true)` to ensure forward compatibility.

- **Patch Method**: An enhanced version of `Unmarshal` that preserves original values.

- **Copy Method**: Implemented based on `Patch`, used for data copying between objects.

## Handling Nil / Null

When marshaling a nil value:

- `protojson.Marshal` returns an empty object (`"{}"`) for nil root Protocol Buffer messages
- This behavior differs from standard `json.Marshal` which would typically return `"null"`
- The jsonx package deliberately follows `json.Marshal`'s behavior for nil Protocol Buffer messages, returning `"null"` instead of an empty object, providing more intuitive and consistent results

When unmarshaling a complete JSON string that is just `"null"`:

- `json.Unmarshal` doesn't produce an error when encountering a JSON string that is just `"null"`; it simply leaves the message unmodified
- `protojson.Unmarshal` returns an error when the entire JSON string is `"null"`, with a message like "unexpected token null"
- The jsonx package follows `json.Unmarshal`'s behavior, ensuring no errors occur when unmarshaling a JSON `"null"` to the destination

## Patch vs. Unmarshal

The key differences between `Patch` and standard unmarshal methods:

1. `json.Unmarshal` can deserialize a patch string to a destination struct while preserving original values for fields not specified in the patch. However, it will lose original values when the target is a nested `map[string]any`.

2. `protojson.Unmarshal` differs from the standard JSON unmarshaler and doesn't preserve original values.

3. `Patch` addresses these limitations by preserving original values even in complex nested structures.

## Deviation from RFC7386

`Patch` doesn't fully comply with the RFC7386 JSON Merge Patch standard for the following reasons:

1. In RFC7386, a `null` value in the patch indicates deletion of the original field in the destination JSON.

2. In Go's typical unmarshal scenarios applied to structs, `null` is usually intended to override the original value.

3. `Patch` works by first using `Marshal(dest)` to create a patch that can restore original values, then combining it with the user-provided patch to create a final patch that meets the required use case.

## Important Note

For array fields, `Patch`, `Copy`, and RFC7386 all use full replacement. If incremental array modifications are needed in your application, you should:

1. Clone the original information before calling `Patch`/`Copy`
2. Manually handle incremental changes after the call

package statusx

import (
	"context"
	"strings"

	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

func BadRequest(inputs ...any) *Status {
	return New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "invalid argument").WithFlattenFieldViolations(inputs...)
}

type ValidatorX interface {
	ValidateX(ctx context.Context) []*errdetails.BadRequest_FieldViolation
}

type ValidateXOption func(*validateXOptions)

type validateXOptions struct {
	fieldPrefix string
}

func WithFieldPrefix(prefix string) ValidateXOption {
	return func(o *validateXOptions) {
		o.fieldPrefix = prefix
	}
}

func ValidateX(ctx context.Context, input any, opts ...ValidateXOption) []*errdetails.BadRequest_FieldViolation {
	options := &validateXOptions{}
	for _, opt := range opts {
		opt(options)
	}
	if val, ok := input.(ValidatorX); ok {
		violations := val.ValidateX(ctx)
		if options.fieldPrefix != "" {
			for _, fv := range violations {
				fv.Field = options.fieldPrefix + fv.Field
			}
		}
		return violations
	}

	// Compatible with proto-gen-validate
	var fvs []*errdetails.BadRequest_FieldViolation
	if val, ok := input.(interface{ ValidateAll() error }); ok {
		err := val.ValidateAll()
		if err == nil {
			return fvs
		}
		inf := err.(interface{ AllErrors() []error })
		for _, vErr := range inf.AllErrors() {
			infPgvErr := vErr.(pgvErr)
			fvs = append(fvs, convertProtoGenErrToFV(infPgvErr, options.fieldPrefix))
		}
	}
	return fvs
}

// proto-gen-validate error
type pgvErr interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
}

const ErrorReasonProtoGenValidate = "PROTO_GEN_VALIDATE"

// convertProtoGenErrToFV converts a protobuf generated validation error into a BadRequest_FieldViolation.
// It processes the field path to ensure proper camelCase formatting while preserving array indices.
//
// Parameters:
//   - pgvErr: The protobuf generated validation error
//   - fieldPrefix: Optional prefix to prepend to the field path (can be empty)
//
// The function:
//  1. Combines the fieldPrefix with the error field if prefix is provided
//  2. Splits the field path on dots and processes each segment:
//     - Preserves array indices (e.g., [0], [1])
//     - Converts path segments to camelCase
//  3. Creates a BadRequest_FieldViolation with:
//     - Formatted field path
//     - Original error reason
//     - Standard proto validation error reason
//
// Returns:
//   - *errdetails.BadRequest_FieldViolation with formatted field information
func convertProtoGenErrToFV(pgvErr pgvErr, fieldPrefix string) *errdetails.BadRequest_FieldViolation {
	field := pgvErr.Field()
	if fieldPrefix != "" {
		field = fieldPrefix + field
	}
	parts := strings.Split(field, ".")
	for i, part := range parts {
		// Find the start position of array index notation (e.g., [0], [1])
		// Returns -1 if no array index is found
		arrStart := strings.LastIndex(part, "[")
		if arrStart != -1 {
			// We cannot use lo.CamelCase(part) directly because it would:
			// 1. Remove the square brackets: "user_addresses[0]" -> "userAddresses0"
			// 2. Make array indices indistinguishable from field names
			// 3. Break client-side error handling that relies on proper array notation
			//
			// Example of incorrect direct usage:
			//   lo.CamelCase("user_addresses[0]") -> "userAddresses0"  // Wrong!
			//
			// Example of current approach:
			//   lo.CamelCase("user_addresses") + "[0]" -> "userAddresses[0]"  // Correct!
			parts[i] = lo.CamelCase(part[:arrStart]) + part[arrStart:]
		} else {
			// For regular fields without array indices, simply convert to camelCase
			parts[i] = lo.CamelCase(part)
		}
	}
	field = strings.Join(parts, ".")
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: pgvErr.Reason(),
		Reason:      ErrorReasonProtoGenValidate, // fmt.Sprintf("%s:%s", pgvErr.ErrorName(), field),
	}
}

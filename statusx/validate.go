package statusx

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
)

func BadRequest(inputs ...any) *Status {
	return New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "invalid argument").WithFlattenFieldViolations(inputs...)
}

type Validator interface {
	Validate() error
}

type ContextValidator interface {
	Validate(ctx context.Context) error
}

func Validate(ctx context.Context, input any) error {
	// Try ContextValidator interface first (has context support)
	if val, ok := input.(ContextValidator); ok {
		return val.Validate(ctx)
	}

	// Try proto-gen-validate interface next (provides better error details than simple Validator)
	var fvs []any
	if val, ok := input.(interface{ ValidateAll() error }); ok {
		err := val.ValidateAll()
		if err == nil {
			return nil
		}
		if inf, ok := err.(interface{ AllErrors() []error }); ok {
			// Multi-error case
			for _, vErr := range inf.AllErrors() {
				if infPgvErr, ok := vErr.(pgvErr); ok {
					fvs = append(fvs, convertProtoGenErrToFV(infPgvErr))
				}
			}
		}
		// Single error case - try to convert directly
		if pgvErr, ok := err.(pgvErr); ok {
			fvs = append(fvs, convertProtoGenErrToFV(pgvErr))
		}

		if len(fvs) > 0 {
			return BadRequest(fvs...).Err()
		}
		// Fallback - return original error if we can't handle it
		return err
	}

	// Fallback to simple Validator interface (for types that only have Validate())
	if val, ok := input.(Validator); ok {
		return val.Validate()
	}

	return errors.New("input does not implement Validator or ContextValidator interface")
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
//
// The function:
//  1. Splits the field path on dots and processes each segment:
//     - Preserves array indices (e.g., [0], [1])
//     - Converts path segments to camelCase
//  2. Creates a BadRequest_FieldViolation with:
//     - Formatted field path
//     - Original error reason
//     - Standard proto validation error reason
//
// Returns:
//   - *errdetails.BadRequest_FieldViolation with formatted field information
func convertProtoGenErrToFV(pgvErr pgvErr) *errdetails.BadRequest_FieldViolation {
	field := pgvErr.Field()
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

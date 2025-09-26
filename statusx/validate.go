package statusx

import (
	"context"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

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
// Returns:
//   - *errdetails.BadRequest_FieldViolation with formatted field information
func convertProtoGenErrToFV(pgvErr pgvErr) *errdetails.BadRequest_FieldViolation {
	field := FormatField(pgvErr.Field(), lo.CamelCase)
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: pgvErr.Reason(),
		Reason:      ErrorReasonProtoGenValidate,
	}
}

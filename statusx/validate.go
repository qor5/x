package statusx

import (
	"context"
	"strings"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"

	"github.com/qor5/x/v3/jsonx"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
)

func BadRequest(fvs ...*errdetails.BadRequest_FieldViolation) *Status {
	return New(codes.InvalidArgument, statusv1.ErrorReason_INVALID_ARGUMENT.String(), "invalid argument").
		WithExtraDetail(&errdetails.BadRequest{FieldViolations: fvs})
}

func PrependFieldPrefix(prefix string, fvs ...*errdetails.BadRequest_FieldViolation) []*errdetails.BadRequest_FieldViolation {
	if prefix == "" {
		return fvs
	}
	for _, fv := range fvs {
		fv.Field = prefix + fv.Field
	}
	return fvs
}

func InvalidArgumentFieldViolation(field string, desc string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: desc,
		Reason:      statusv1.ErrorReason_INVALID_ARGUMENT.String(),
	}
}

func AssertErrorFieldViolation(t *testing.T, err error, fvs ...*errdetails.BadRequest_FieldViolation) {
	st := Convert(err)
	assert.Equal(t, codes.InvalidArgument, st.Code(), "error code mismatch")
	assert.Equal(t, "invalid argument", st.Message(), "error message mismatch")

	pb, ok := lo.Find(st.Details(), func(d any) bool {
		_, ok := d.(*errdetails.BadRequest)
		return ok
	})
	if assert.True(t, ok, "BadRequest not found in error details") {
		bq := pb.(*errdetails.BadRequest)
		for _, v := range bq.GetFieldViolations() {
			v.LocalizedMessage = nil
		}
		for _, fv := range fvs {
			if _, ok := lo.Find(bq.GetFieldViolations(), func(d *errdetails.BadRequest_FieldViolation) bool {
				return proto.Equal(d, fv)
			}); !ok {
				t.Errorf("field violation %s not found in bad request %s", jsonx.MustMarshalX[string](fv), jsonx.MustMarshalX[string](bq))
			}
		}
	}
}

type ValidatorX interface {
	ValidateX(ctx context.Context) []*errdetails.BadRequest_FieldViolation
}

type ValidateXOption func(*validateXOpts)

type validateXOpts struct {
	fieldPrefix string
}

func WithFieldPrefix(prefix string) ValidateXOption {
	return func(o *validateXOpts) {
		o.fieldPrefix = prefix
	}
}

func ValidateX(ctx context.Context, input any, opts ...ValidateXOption) []*errdetails.BadRequest_FieldViolation {
	conf := &validateXOpts{}
	for _, opt := range opts {
		opt(conf)
	}
	if val, ok := input.(ValidatorX); ok {
		return PrependFieldPrefix(conf.fieldPrefix, val.ValidateX(ctx)...)
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
			fvs = append(fvs, convertProtoGenErrToFV(infPgvErr, conf.fieldPrefix))
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

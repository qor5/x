package statusx

import (
	"testing"

	"github.com/qor5/x/v3/jsonx"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

func AssertFieldViolations(t *testing.T, err error, fvs ...*errdetails.BadRequest_FieldViolation) {
	st := Convert(err)
	assert.Equal(t, codes.InvalidArgument, st.Code(), "error code mismatch")
	assert.Equal(t, "invalid argument", st.Message(), "error message mismatch")

	badRequest := ExtractDetail[*errdetails.BadRequest](st.Details())
	if assert.NotNil(t, badRequest, "BadRequest not found in error details") {
		for _, v := range badRequest.GetFieldViolations() {
			v.LocalizedMessage = nil
		}
		for _, fv := range fvs {
			if _, ok := lo.Find(badRequest.GetFieldViolations(), func(d *errdetails.BadRequest_FieldViolation) bool {
				return proto.Equal(d, fv)
			}); !ok {
				t.Errorf("field violation %s not found in bad request %s", jsonx.MustMarshalX[string](fv), jsonx.MustMarshalX[string](badRequest))
			}
		}
	}
}

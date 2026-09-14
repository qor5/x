package httperrors

import (
	"net/http"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func AssertFieldViolations(t *testing.T, err error, fvs ...*FieldViolation) {
	t.Helper()
	st := Convert(err)
	assert.Equal(t, http.StatusBadRequest, st.StatusCode(), "error status code mismatch")
	assert.Equal(t, "invalid argument", st.Message(), "error message mismatch")

	actualFVs := st.FieldViolations()
	for _, expected := range fvs {
		if _, ok := lo.Find(actualFVs, func(actual *FieldViolation) bool {
			return actual.Field() == expected.Field() &&
				actual.Reason() == expected.Reason() &&
				actual.Description() == expected.Description()
		}); !ok {
			t.Errorf("field violation (field=%q, reason=%q, description=%q) not found in actual violations",
				expected.Field(), expected.Reason(), expected.Description())
		}
	}
}

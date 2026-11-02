package httperrors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReasonFromStatus(t *testing.T) {
	tests := []struct {
		httpStatus int
		expected   string
	}{
		{http.StatusOK, ReasonOK},
		{http.StatusBadRequest, ReasonInvalidArgument},
		{http.StatusUnauthorized, ReasonUnauthenticated},
		{http.StatusForbidden, ReasonPermissionDenied},
		{http.StatusNotFound, ReasonNotFound},
		{http.StatusConflict, ReasonAlreadyExists},
		{http.StatusUnprocessableEntity, ReasonInvalidArgument},
		{http.StatusTooManyRequests, ReasonResourceExhausted},
		{http.StatusInternalServerError, ReasonInternal},
		{http.StatusNotImplemented, ReasonUnimplemented},
		{http.StatusBadGateway, ReasonUnavailable},
		{http.StatusServiceUnavailable, ReasonUnavailable},
		{http.StatusGatewayTimeout, ReasonDeadlineExceeded},
		{499, ReasonCanceled},
		{999, ReasonUnknown},
		{http.StatusTeapot, ReasonUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, ReasonFromStatus(tt.httpStatus))
		})
	}
}

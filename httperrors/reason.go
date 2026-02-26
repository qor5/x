package httperrors

import "net/http"

// Common reason constants for HTTP errors.
// These replace the protobuf ErrorReason enum from statusx.
const (
	ReasonOK                = "OK"
	ReasonCanceled          = "CANCELED"
	ReasonUnknown           = "UNKNOWN"
	ReasonBadRequest        = "BAD_REQUEST"
	ReasonInvalidArgument   = "INVALID_ARGUMENT"
	ReasonDeadlineExceeded  = "DEADLINE_EXCEEDED"
	ReasonNotFound          = "NOT_FOUND"
	ReasonAlreadyExists     = "ALREADY_EXISTS"
	ReasonConflict          = "CONFLICT"
	ReasonPermissionDenied  = "PERMISSION_DENIED"
	ReasonResourceExhausted = "RESOURCE_EXHAUSTED"
	ReasonFailedPrecondition = "FAILED_PRECONDITION"
	ReasonAborted           = "ABORTED"
	ReasonOutOfRange        = "OUT_OF_RANGE"
	ReasonUnimplemented     = "UNIMPLEMENTED"
	ReasonInternal          = "INTERNAL"
	ReasonUnavailable       = "UNAVAILABLE"
	ReasonDataLoss          = "DATA_LOSS"
	ReasonUnauthenticated   = "UNAUTHENTICATED"
	ReasonBadGateway        = "BAD_GATEWAY"
)

// ReasonFromStatus returns a default reason string for a given HTTP status code.
func ReasonFromStatus(httpStatus int) string {
	switch httpStatus {
	case http.StatusOK:
		return ReasonOK
	case http.StatusBadRequest:
		return ReasonBadRequest
	case http.StatusUnauthorized:
		return ReasonUnauthenticated
	case http.StatusForbidden:
		return ReasonPermissionDenied
	case http.StatusNotFound:
		return ReasonNotFound
	case http.StatusConflict:
		return ReasonConflict
	case http.StatusUnprocessableEntity:
		return ReasonInvalidArgument
	case http.StatusTooManyRequests:
		return ReasonResourceExhausted
	case http.StatusInternalServerError:
		return ReasonInternal
	case http.StatusNotImplemented:
		return ReasonUnimplemented
	case http.StatusBadGateway:
		return ReasonBadGateway
	case http.StatusServiceUnavailable:
		return ReasonUnavailable
	case http.StatusGatewayTimeout:
		return ReasonDeadlineExceeded
	case 499: // Client Closed Request
		return ReasonCanceled
	default:
		return ReasonUnknown
	}
}

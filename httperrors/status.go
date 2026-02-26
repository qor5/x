package httperrors

import (
	"context"
	"fmt"
	"io"
	"maps"
	"net/http"
	"slices"

	"github.com/pkg/errors"
)

type Status struct {
	httpStatus int
	reason     string
	message    string

	metadata        map[string]string
	localized       *Localized
	fieldViolations []*FieldViolation

	cause error
}

// New creates a Status with the specified HTTP status code, reason, and message.
//
// For non-2xx status codes, it automatically captures a stacktrace at creation time,
// which provides valuable debugging context without manual instrumentation.
//
// Parameters:
//   - httpStatus: The HTTP status code (e.g., 400, 404, 500)
//   - reason: A string identifier used as the error reason and i18n key fallback during translation
//   - message: A human-readable message for debugging purposes
//
// The reason serves as both the error identifier and the i18n key.
// The reason is immediately fixed as the i18n key at creation time.
// Use WithLocalized() to override with a specific i18n key and args if needed.
// Use WithLocalizedArgs() to add template arguments while preserving the current key.
// Returns a Status object that can be further enriched with metadata or localization.
func New(httpStatus int, reason, message string) *Status {
	if reason == "" {
		panic("reason is required")
	}
	s := &Status{
		httpStatus: httpStatus,
		reason:     reason,
		message:    message,
		localized:  &Localized{key: reason},
	}
	if httpStatus < 200 || httpStatus >= 300 {
		s.cause = errors.New(message)
	}
	return s
}

func Newf(httpStatus int, reason, format string, a ...any) *Status {
	return New(httpStatus, reason, fmt.Sprintf(format, a...))
}

// NewStatus creates a Status with automatically derived reason from the HTTP status code.
// This is a convenience function that uses ReasonFromStatus to generate the reason.
func NewStatus(httpStatus int, message string) *Status {
	return New(httpStatus, ReasonFromStatus(httpStatus), message)
}

// NewStatusf creates a Status with automatically derived reason and formatted message.
func NewStatusf(httpStatus int, format string, a ...any) *Status {
	return Newf(httpStatus, ReasonFromStatus(httpStatus), format, a...)
}

func Error(httpStatus int, reason, message string) error {
	return New(httpStatus, reason, message).Err()
}

func Errorf(httpStatus int, reason, format string, a ...any) error {
	return Error(httpStatus, reason, fmt.Sprintf(format, a...))
}

func (s *Status) StatusCode() int {
	if s == nil {
		return http.StatusOK
	}
	if s.httpStatus >= 200 && s.httpStatus < 300 && s.cause != nil {
		return http.StatusInternalServerError
	}
	return s.httpStatus
}

func (s *Status) Reason() string {
	if s == nil {
		return ReasonOK
	}
	if s.httpStatus >= 200 && s.httpStatus < 300 {
		if s.cause != nil {
			return ReasonUnknown
		}
		return ReasonOK
	}
	return s.reason
}

func (s *Status) Message() string {
	if s == nil {
		return ""
	}
	return s.message
}

func (s *Status) Cause() error {
	if s == nil {
		return nil
	}
	return s.cause //nolint:errhandle
}

func (s *Status) Metadata() map[string]string {
	if s == nil {
		return nil
	}
	return maps.Clone(s.metadata)
}

func (s *Status) Localized() *Localized {
	if s == nil || s.localized == nil {
		return nil
	}
	return s.localized.Clone()
}

func (s *Status) FieldViolations() []*FieldViolation {
	if s == nil {
		return nil
	}
	return cloneFieldViolations(s.fieldViolations)
}

// ToFieldViolations converts this Status to field violations for the specified field
func (s *Status) ToFieldViolations(field string) FieldViolations {
	if s == nil {
		return nil
	}
	return ToFieldViolations(s.Err(), field)
}

// Err converts the Status to an error interface.
//
// The returned error type is either:
//   - nil: When StatusCode() is 2xx
//   - *StatusError: An error that implements the error interface
func (s *Status) Err() error {
	if s.StatusCode() >= 200 && s.StatusCode() < 300 {
		return nil
	}
	return &StatusError{s: s} //nolint:errhandle
}

func (s *Status) WithStatusCode(httpStatus int) *Status {
	st := Clone(s)
	st.httpStatus = httpStatus
	return st
}

func (s *Status) WithMessage(message string) *Status {
	st := Clone(s)
	st.message = message
	return st
}

func (s *Status) WithMessagef(format string, a ...any) *Status {
	return s.WithMessage(fmt.Sprintf(format, a...))
}

func (s *Status) WithReason(reason string) *Status {
	st := Clone(s)
	st.reason = reason
	return st
}

func (s *Status) WithCause(cause error) *Status {
	st := Clone(s)
	st.cause = errors.WithStack(cause)
	return st
}

func (s *Status) WithMetadata(md map[string]string) *Status {
	st := Clone(s)
	st.metadata = maps.Clone(md)
	return st
}

func (s *Status) WithLocalized(key string, args ...any) *Status {
	if key == "" {
		panic("key is required")
	}
	st := Clone(s)
	st.localized = &Localized{
		key:  key,
		args: args,
	}
	return st
}

// WithLocalizedArgs sets template arguments for i18n.
// Preserves the existing localized key and adds/replaces the template arguments.
// Since the key is always set (either at creation time or by WithLocalized), no fallback logic is needed.
func (s *Status) WithLocalizedArgs(args ...any) *Status {
	st := Clone(s)
	st.localized = &Localized{
		key:  st.localized.Key(),
		args: args,
	}
	return st
}

// WithFieldViolations adds multiple field-level validation violations in batch.
// Multiple violations for the same field are allowed and will be appended.
func (s *Status) WithFieldViolations(fieldViolations ...*FieldViolation) *Status {
	st := Clone(s)

	for _, fv := range fieldViolations {
		field := fv.Field()
		if field == "" {
			panic(errors.Errorf("field is required"))
		}
		if fv.Reason() == "" {
			panic(errors.Errorf("reason is required"))
		}
		st.fieldViolations = append(st.fieldViolations, fv.Clone())
	}

	return st
}

// WithFlattenFieldViolations accepts various types of field violation inputs and flattens them.
// Supports *FieldViolation, []*FieldViolation, FieldViolations.
// Mixed types are allowed in a single call for maximum flexibility.
//
// Note: For error and *Status inputs, use ToFieldViolations(err, field) or status.ToFieldViolations(field)
// first to specify the field name, then pass the result to this function.
func (s *Status) WithFlattenFieldViolations(inputs ...any) *Status {
	violations, err := FlattenFieldViolations(inputs...)
	if err != nil {
		panic(err)
	}
	return s.WithFieldViolations(violations...)
}

func (s *Status) String() string {
	return fmt.Sprintf("http error: status = %d reason = %s message = %s", s.StatusCode(), s.Reason(), s.message)
}

func Clone(s *Status) *Status {
	if s == nil {
		return nil
	}
	var localized *Localized
	if s.localized != nil {
		localized = s.localized.Clone()
	}
	return &Status{
		httpStatus:      s.httpStatus,
		reason:          s.reason,
		message:         s.message,
		metadata:        maps.Clone(s.metadata),
		localized:       localized,
		fieldViolations: cloneFieldViolations(s.fieldViolations),
		cause:           s.cause,
	}
}

type StatusError struct {
	s *Status
}

func (e *StatusError) Error() string {
	msg := e.s.String()
	if e.s.cause != nil {
		msg += ": " + e.s.cause.Error()
	}
	return msg
}

func (e *StatusError) Cause() error {
	return e.s.Cause()
}

func (e *StatusError) Unwrap() error { return e.s.Cause() }

func (e *StatusError) Format(s fmt.State, verb rune) {
	// nolint:errcheck
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", e.Cause())
			io.WriteString(s, e.s.String())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

// Is compares two StatusErrors by httpStatus + reason (decision: scheme A).
func (e *StatusError) Is(target error) bool {
	tse, ok := target.(*StatusError)
	if !ok {
		return false
	}
	return e.s.StatusCode() == tse.s.StatusCode() && e.s.Reason() == tse.s.Reason()
}

func (e *StatusError) Status() *Status {
	return e.s
}

func FromError(err error) (s *Status, ok bool) {
	if err == nil {
		return New(http.StatusOK, ReasonOK, ""), true
	}
	var se *StatusError
	if errors.As(err, &se) {
		return se.s, true
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return New(http.StatusGatewayTimeout, ReasonDeadlineExceeded, err.Error()), true
	}
	if errors.Is(err, context.Canceled) {
		return New(499, ReasonCanceled, err.Error()), true
	}
	return New(http.StatusInternalServerError, ReasonUnknown, err.Error()), false
}

func StatusCode(err error) int {
	return Convert(err).StatusCode()
}

func Reason(err error) string {
	return Convert(err).Reason()
}

func Convert(err error) *Status {
	s, _ := FromError(err)
	return s
}

func Wrap(err error, httpStatus int, reason, message string) *Status {
	s, ok := FromError(err)
	if ok {
		return s
	}
	s.cause = errors.WithStack(err)
	s.httpStatus = httpStatus
	s.message = message
	s.reason = reason
	s.localized = &Localized{key: s.Reason()}
	return s
}

func Wrapf(err error, httpStatus int, reason, format string, a ...any) *Status {
	return Wrap(err, httpStatus, reason, fmt.Sprintf(format, a...))
}

// WrapStatus wraps an error with automatically derived reason from the HTTP status code.
func WrapStatus(err error, httpStatus int, message string) *Status {
	return Wrap(err, httpStatus, ReasonFromStatus(httpStatus), message)
}

// WrapStatusf wraps an error with automatically derived reason and formatted message.
func WrapStatusf(err error, httpStatus int, format string, a ...any) *Status {
	return Wrapf(err, httpStatus, ReasonFromStatus(httpStatus), format, a...)
}

func cloneFieldViolations(fvs []*FieldViolation) []*FieldViolation {
	if fvs == nil {
		return nil
	}
	return slices.Clone(fvs)
}

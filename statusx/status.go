package statusx

import (
	"context"
	"fmt"
	"io"
	"maps"

	"github.com/pkg/errors"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
)

type Status struct {
	code    codes.Code
	message string

	errorInfo  *errdetails.ErrorInfo
	localized  *statusv1.Localized
	badRequest *statusv1.BadRequest
	details    []proto.Message

	cause error
}

// New creates a Status with the specified code, reason, and message.
//
// For non-OK status codes, it automatically captures a stacktrace at creation time,
// which provides valuable debugging context without manual instrumentation.
//
// Parameters:
//   - c: The status code indicating the type of status
//   - reason: A string identifier used as the error reason and i18n key fallback during translation
//   - message: A human-readable message for debugging purposes
//
// The reason serves as both the error identifier and the i18n key.
// The reason is immediately fixed as the i18n key at creation time.
// Use WithLocalized() to override with a specific i18n key and args if needed.
// Use WithLocalizedArgs() to add template arguments while preserving the current key.
// Returns a Status object that can be further enriched with metadata or localization.
func New(c codes.Code, reason, message string) *Status {
	if reason == "" {
		panic("reason is required")
	}
	s := &Status{
		code:      c,
		message:   message,
		errorInfo: &errdetails.ErrorInfo{Reason: reason},
	}
	if c != codes.OK {
		s.cause = errors.New(message)
	}
	// Immediately fix key to creation-time reason
	s.localized = &statusv1.Localized{Key: s.Reason()}
	return s
}

func Newf(c codes.Code, reason, format string, a ...any) *Status {
	return New(c, reason, fmt.Sprintf(format, a...))
}

func Error(c codes.Code, reason, message string) error {
	return New(c, reason, message).Err()
}

func Errorf(c codes.Code, reason, format string, a ...any) error {
	return Error(c, reason, fmt.Sprintf(format, a...))
}

func (s *Status) GRPCStatus() *status.Status {
	if s == nil {
		return nil
	}

	if s.Code() == codes.OK {
		return status.New(codes.OK, "")
	}

	var details []protoadapt.MessageV1

	errorInfo := proto.Clone(s.errorInfo).(*errdetails.ErrorInfo)
	errorInfo.Reason = s.Reason()
	details = append(details, errorInfo)

	if s.localized != nil {
		details = append(details, proto.Clone(s.localized).(*statusv1.Localized))
	}

	if s.badRequest != nil {
		details = append(details, proto.Clone(s.badRequest).(*statusv1.BadRequest))
	}

	if len(s.details) > 0 {
		for _, d := range s.details {
			details = append(details, protoadapt.MessageV1Of(proto.Clone(d)))
		}
	}

	st, err := status.New(s.Code(), s.message).WithDetails(details...)
	if err != nil {
		panic(err)
	}
	return st
}

func (s *Status) Details() []any {
	return s.GRPCStatus().Details()
}

func (s *Status) Code() codes.Code {
	if s == nil {
		return codes.OK
	}
	if s.code == codes.OK && s.cause != nil {
		return codes.Unknown
	}
	return s.code
}

func (s *Status) Reason() string {
	if s == nil {
		return statusv1.ErrorReason_OK.String()
	}
	if s.code == codes.OK {
		if s.cause != nil {
			// At this point, s.reason should not be trusted either.
			return statusv1.ErrorReason_UNKNOWN.String()
		}
		return statusv1.ErrorReason_OK.String()
	}
	return s.errorInfo.Reason
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
	return maps.Clone(s.errorInfo.Metadata)
}

func (s *Status) Localized() *statusv1.Localized {
	if s == nil || s.localized == nil {
		return nil
	}
	return proto.Clone(s.localized).(*statusv1.Localized)
}

func (s *Status) BadRequest() *statusv1.BadRequest {
	if s == nil || s.badRequest == nil {
		return nil
	}
	return proto.Clone(s.badRequest).(*statusv1.BadRequest)
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
//   - nil: When Status.Code() is codes.OK
//   - *StatusError: An error that implements both error and GRPCStatus interfaces
func (s *Status) Err() error {
	if s.Code() == codes.OK {
		return nil
	}
	return &StatusError{s: s} //nolint:errhandle
}

func (s *Status) WithCode(c codes.Code) *Status {
	st := Clone(s)
	st.code = c
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
	st.errorInfo.Reason = reason
	return st
}

func (s *Status) WithCause(cause error) *Status {
	st := Clone(s)
	st.cause = errors.WithStack(cause)
	return st
}

func (s *Status) WithMetadata(md map[string]string) *Status {
	st := Clone(s)
	st.errorInfo.Metadata = md
	return st
}

func (s *Status) WithLocalized(key string, args ...any) *Status {
	if key == "" {
		panic("key is required")
	}
	st := Clone(s)
	st.localized = &statusv1.Localized{
		Key:  key,
		Args: convertArgsToAny(args),
	}
	return st
}

// WithLocalizedArgs sets template arguments for i18n.
// Preserves the existing localized key and adds/replaces the template arguments.
// Since the key is always set (either at creation time or by WithLocalized), no fallback logic is needed.
func (s *Status) WithLocalizedArgs(args ...any) *Status {
	st := Clone(s)
	st.localized = &statusv1.Localized{
		Key:  st.localized.GetKey(), // Always non-empty due to New() or WithLocalized()
		Args: convertArgsToAny(args),
	}
	return st
}

// WithFieldViolations adds multiple field-level validation violations in batch.
// Multiple violations for the same field are allowed and will be appended.
func (s *Status) WithFieldViolations(fieldViolations ...*FieldViolation) *Status {
	st := Clone(s)

	if st.badRequest == nil {
		st.badRequest = &statusv1.BadRequest{}
	}

	// Simply append all field violations without deduplication
	for _, fv := range fieldViolations {
		field := fv.Field()
		if field == "" {
			panic(errors.Errorf("field is required"))
		}
		if fv.Reason() == "" {
			panic(errors.Errorf("reason is required"))
		}
		st.badRequest.FieldViolations = append(st.badRequest.FieldViolations, fv.Proto())
	}

	return st
}

// WithFlattenFieldViolations accepts various types of field violation inputs and flattens them.
// Supports *FieldViolation, []*FieldViolation, FieldViolations, and their protobuf equivalents.
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

func (s *Status) WithDetails(details ...proto.Message) *Status {
	st := Clone(s)

	// Validate that details don't contain reserved types
	for _, detail := range details {
		switch detail.(type) {
		case *errdetails.ErrorInfo:
			panic("cannot add *errdetails.ErrorInfo via WithDetails, it's managed internally")
		case *statusv1.Localized:
			panic("cannot add *statusv1.Localized via WithDetails, use WithLocalized instead")
		case *statusv1.BadRequest:
			panic("cannot add *statusv1.BadRequest via WithDetails, use WithFieldViolations instead")
		}
	}

	st.details = append(st.details, details...)
	return st
}

func (s *Status) String() string {
	return fmt.Sprintf("rpc error: code = %s reason = %s message = %s", s.Code().String(), s.Reason(), s.message)
}

func Clone(s *Status) *Status {
	if s == nil {
		return nil
	}
	var errorInfo *errdetails.ErrorInfo
	if s.errorInfo != nil {
		errorInfo = proto.Clone(s.errorInfo).(*errdetails.ErrorInfo)
	}
	var localized *statusv1.Localized
	if s.localized != nil {
		localized = proto.Clone(s.localized).(*statusv1.Localized)
	}
	var badRequest *statusv1.BadRequest
	if s.badRequest != nil {
		badRequest = proto.Clone(s.badRequest).(*statusv1.BadRequest)
	}
	var details []proto.Message
	if s.details != nil {
		details = lo.Map(s.details, func(d proto.Message, _ int) proto.Message { return proto.Clone(d) })
	}
	return &Status{
		code:       s.code,
		message:    s.message,
		errorInfo:  errorInfo,
		localized:  localized,
		badRequest: badRequest,
		details:    details,
		cause:      s.cause,
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

func (e *StatusError) Is(target error) bool {
	tse, ok := target.(*StatusError)
	if !ok {
		return false
	}
	return proto.Equal(e.GRPCStatus().Proto(), tse.GRPCStatus().Proto())
}

func (e *StatusError) Status() *Status {
	return e.s
}

func (e *StatusError) GRPCStatus() *status.Status {
	return e.s.GRPCStatus()
}

func FromError(err error) (s *Status, ok bool) {
	if err == nil {
		// This is intentionally different from status.FromError(nil) behavior for convenience.
		return New(codes.OK, statusv1.ErrorReason_OK.String(), ""), true
	}
	var se *StatusError
	if errors.As(err, &se) {
		// if err is already a *StatusError, we don't want to lose the original status.
		return se.s, true
	}
	ss, ok := status.FromError(err)
	if !ok {
		if errors.Is(err, context.DeadlineExceeded) {
			return New(codes.DeadlineExceeded, statusv1.ErrorReason_DEADLINE_EXCEEDED.String(), err.Error()), true
		}
		if errors.Is(err, context.Canceled) {
			return New(codes.Canceled, statusv1.ErrorReason_CANCELED.String(), err.Error()), true
		}
		return New(codes.Unknown, statusv1.ErrorReason_UNKNOWN.String(), err.Error()), false
	}
	s = New(ss.Code(), statusv1.ErrorReason_UNKNOWN.String(), ss.Message())
	for _, detail := range ss.Details() {
		switch d := detail.(type) {
		case *errdetails.ErrorInfo:
			s.errorInfo = proto.Clone(d).(*errdetails.ErrorInfo)
		case *statusv1.Localized:
			s.localized = proto.Clone(d).(*statusv1.Localized)
		case *statusv1.BadRequest:
			s.badRequest = proto.Clone(d).(*statusv1.BadRequest)
		default:
			s.details = append(s.details, proto.Clone(d.(proto.Message)))
		}
	}
	return s, true
}

func Code(err error) codes.Code {
	return Convert(err).Code()
}

func Reason(err error) string {
	return Convert(err).Reason()
}

func Convert(err error) *Status {
	s, _ := FromError(err)
	return s
}

func Wrap(err error, c codes.Code, reason, message string) *Status {
	s, ok := FromError(err)
	if ok {
		return s
	}
	s.cause = errors.WithStack(err)
	s.code = c
	s.message = message
	s.errorInfo.Reason = reason
	// Immediately fix key to creation-time reason
	s.localized = &statusv1.Localized{Key: s.Reason()}
	return s
}

func Wrapf(err error, c codes.Code, reason, format string, a ...any) *Status {
	return Wrap(err, c, reason, fmt.Sprintf(format, a...))
}

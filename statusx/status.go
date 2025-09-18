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

	proto   *statusv1.Status
	details []proto.Message

	cause error
}

// New creates a Status with the specified code, reason, and message.
//
// For non-OK status codes, it automatically captures a stacktrace at creation time,
// which provides valuable debugging context without manual instrumentation.
//
// Parameters:
//   - c: The status code indicating the type of status
//   - reason: A string identifier describing the reason for this status
//   - message: A human-readable message describing the status
//
// Returns a Status object that can be further enriched with metadata or localization.
func New(c codes.Code, reason, message string) *Status {
	s := &Status{
		code:    c,
		message: message,
		proto: &statusv1.Status{
			ErrorInfo: &errdetails.ErrorInfo{
				Reason: reason,
			},
		},
	}
	if c != codes.OK {
		s.cause = errors.New(message)
	}
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

	proto := proto.Clone(s.proto).(*statusv1.Status)
	proto.ErrorInfo.Reason = s.Reason()
	details = append(details, proto.ErrorInfo)

	// Add other custom details
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
	if s.proto != nil && s.proto.ErrorInfo != nil {
		return s.proto.ErrorInfo.Reason
	}
	return ""
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
	if s == nil || s.proto == nil || s.proto.ErrorInfo == nil {
		return nil
	}
	return maps.Clone(s.proto.ErrorInfo.Metadata)
}

func (s *Status) Localized() *statusv1.Localized {
	if s == nil || s.proto == nil {
		return nil
	}
	return proto.Clone(s.proto.Localized).(*statusv1.Localized)
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

func (s *Status) WithReason(reason string) *Status {
	st := Clone(s)
	if st.proto.ErrorInfo == nil {
		st.proto.ErrorInfo = &errdetails.ErrorInfo{}
	}
	st.proto.ErrorInfo.Reason = reason
	return st
}

func (s *Status) WithCause(cause error) *Status {
	st := Clone(s)
	st.cause = errors.WithStack(cause)
	return st
}

func (s *Status) WithMetadata(md map[string]string) *Status {
	st := Clone(s)
	if st.proto.ErrorInfo == nil {
		st.proto.ErrorInfo = &errdetails.ErrorInfo{}
	}
	st.proto.ErrorInfo.Metadata = md
	return st
}

// WithTopLevelLocalized sets the top-level localization for the entire error
func (s *Status) WithTopLevelLocalized(key string, args ...any) *Status {
	st := Clone(s)
	if st.proto == nil {
		st.proto = &statusv1.Status{}
	}
	st.proto.Localized = &statusv1.Localized{
		Key:  key,
		Args: convertArgsToAny(args),
	}
	return st
}

// WithFieldViolation adds a single field-level validation violation
func (s *Status) WithFieldViolation(field, reason, description string, localizedKey string, args ...any) *Status {
	return s.WithFieldViolations(&FieldViolation{
		Field:       field,
		Description: description,
		Reason:      reason,
		Key:         localizedKey,
		Args:        args,
	})
}

// WithFieldViolations adds multiple field-level validation violations in batch
func (s *Status) WithFieldViolations(violations ...*FieldViolation) *Status {
	st := Clone(s)
	if st.proto == nil {
		st.proto = &statusv1.Status{}
	}

	// Convert FieldViolation structs to proto FieldViolation
	for _, violation := range violations {
		protoViolation := &statusv1.FieldViolation{
			Field:       violation.Field,
			Description: violation.Description,
			Reason:      violation.Reason,
		}

		// Set localization if provided
		if violation.Key != "" {
			protoViolation.Localized = &statusv1.Localized{
				Key:  violation.Key,
				Args: convertArgsToAny(violation.Args),
			}
		}

		st.proto.FieldViolations = append(st.proto.FieldViolations, protoViolation)
	}

	return st
}

func (s *Status) WithDetails(details ...proto.Message) *Status {
	st := Clone(s)

	// Validate that details don't contain reserved types
	for _, detail := range details {
		switch detail.(type) {
		case *statusv1.Localized:
			panic("cannot add *statusv1.Localized via WithDetails, use WithLocalized instead")
		case *errdetails.ErrorInfo:
			panic("cannot add *errdetails.ErrorInfo via WithDetails, it's managed internally")
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

	var clonedProto *statusv1.Status
	if s.proto != nil {
		clonedProto = proto.Clone(s.proto).(*statusv1.Status)
	}

	return &Status{
		code:    s.code,
		message: s.message,
		proto:   clonedProto,
		details: lo.Map(s.details, func(d proto.Message, _ int) proto.Message { return proto.Clone(d) }),
		cause:   s.cause,
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
	if se := new(StatusError); errors.As(err, &se) {
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
		case *statusv1.Status:
			s.proto = proto.Clone(d).(*statusv1.Status)
		case *errdetails.ErrorInfo:
			// Handle legacy ErrorInfo by merging into our proto structure
			if s.proto == nil {
				s.proto = &statusv1.Status{}
			}
			s.proto.ErrorInfo = proto.Clone(d).(*errdetails.ErrorInfo)
		case *statusv1.Localized:
			// Handle legacy Localized by merging into our proto structure
			if s.proto == nil {
				s.proto = &statusv1.Status{}
			}
			s.proto.Localized = proto.Clone(d).(*statusv1.Localized)
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
	s = New(c, reason, message)
	s.cause = errors.WithStack(err)
	return s
}

func Wrapf(err error, c codes.Code, reason, format string, a ...any) *Status {
	return Wrap(err, c, reason, fmt.Sprintf(format, a...))
}

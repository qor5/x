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

	"github.com/qor5/x/v3/i18nx"
)

type Status struct {
	code    codes.Code
	message string

	errorInfo    *errdetails.ErrorInfo
	localized    *statusv1.Localized
	extraDetails []proto.Message

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
		errorInfo: &errdetails.ErrorInfo{
			Reason: reason,
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
	errorInfo := proto.Clone(s.errorInfo).(*errdetails.ErrorInfo)
	errorInfo.Reason = s.Reason()
	details = append(details, errorInfo)
	if s.localized != nil {
		details = append(details, proto.Clone(s.localized).(*statusv1.Localized))
	}
	if len(s.extraDetails) > 0 {
		for _, d := range s.extraDetails {
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
	if s == nil {
		return nil
	}
	return proto.Clone(s.localized).(*statusv1.Localized)
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
	err := Clone(s)
	err.code = c
	return err
}

func (s *Status) WithCause(cause error) *Status {
	err := Clone(s)
	err.cause = errors.WithStack(cause)
	return err
}

func (s *Status) WithMetadata(md map[string]string) *Status {
	err := Clone(s)
	err.errorInfo.Metadata = md
	return err
}

func (s *Status) WithLocalized(key i18nx.Key, args ...string) *Status {
	err := Clone(s)
	err.localized = &statusv1.Localized{
		Key:  string(key),
		Args: args,
	}
	return err
}

func (s *Status) WithExtraDetail(message ...proto.Message) *Status {
	err := Clone(s)
	err.extraDetails = append(err.extraDetails, message...)
	return err
}

func (s *Status) String() string {
	return fmt.Sprintf("rpc error: code = %s reason = %s message = %s", s.Code().String(), s.Reason(), s.message)
}

func Clone(s *Status) *Status {
	if s == nil {
		return nil
	}
	return &Status{
		code:         s.code,
		message:      s.message,
		errorInfo:    proto.Clone(s.errorInfo).(*errdetails.ErrorInfo),
		localized:    proto.Clone(s.localized).(*statusv1.Localized),
		extraDetails: lo.Map(s.extraDetails, func(d proto.Message, _ int) proto.Message { return proto.Clone(d) }),
		cause:        s.cause,
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
		case *errdetails.ErrorInfo:
			s.errorInfo = proto.Clone(d).(*errdetails.ErrorInfo)
		case *statusv1.Localized:
			s.localized = proto.Clone(d).(*statusv1.Localized)
		default:
			s.extraDetails = append(s.extraDetails, proto.Clone(d.(proto.Message)))
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
	return s
}

func Wrapf(err error, c codes.Code, reason, format string, a ...any) *Status {
	return Wrap(err, c, reason, fmt.Sprintf(format, a...))
}

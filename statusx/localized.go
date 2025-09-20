package statusx

import (
	"log/slog"
	"time"

	"github.com/pkg/errors"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/qor5/x/v3/i18nx"
)

// NewLocalized creates a new Status with the given code, key, and args
func NewLocalized(code codes.Code, key string, args ...any) *Status {
	return New(code, ReasonFromCode(code).String(), key).WithLocalized(key, args...)
}

// WrapLocalized wraps an error with the given code, key, and args
func WrapLocalized(err error, code codes.Code, key string, args ...any) *Status {
	return Wrap(err, code, ReasonFromCode(code).String(), key).WithLocalized(key, args...)
}

// toProtoMessage converts Go values to proto messages for use with Any.
func toProtoMessage(v any) (proto.Message, error) {
	switch val := v.(type) {
	case proto.Message:
		return val, nil
	case string:
		return wrapperspb.String(val), nil
	case int:
		return wrapperspb.Int64(int64(val)), nil
	case int32:
		return wrapperspb.Int32(val), nil
	case int64:
		return wrapperspb.Int64(val), nil
	case uint:
		return wrapperspb.UInt64(uint64(val)), nil
	case uint32:
		return wrapperspb.UInt32(val), nil
	case uint64:
		return wrapperspb.UInt64(val), nil
	case float32:
		return wrapperspb.Float(val), nil
	case float64:
		return wrapperspb.Double(val), nil
	case bool:
		return wrapperspb.Bool(val), nil
	case []byte:
		return wrapperspb.Bytes(val), nil
	case time.Time:
		return timestamppb.New(val), nil
	case time.Duration:
		return durationpb.New(val), nil
	case map[string]any:
		s, err := structpb.NewStruct(val)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert map to struct")
		}
		return s, nil
	case nil:
		return &emptypb.Empty{}, nil
	default:
		return nil, errors.Errorf("unsupported type for protobuf Any conversion: %T", val)
	}
}

// extractValueFromAny converts protobuf Any back to Go values.
func extractValueFromAny(anyVal *anypb.Any) (any, error) {
	if anyVal == nil {
		return nil, nil
	}

	msg, err := anypb.UnmarshalNew(anyVal, proto.UnmarshalOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal Any")
	}

	switch v := msg.(type) {
	case *wrapperspb.StringValue:
		return v.GetValue(), nil
	case *wrapperspb.Int64Value:
		return v.GetValue(), nil
	case *wrapperspb.Int32Value:
		return v.GetValue(), nil
	case *wrapperspb.UInt64Value:
		return v.GetValue(), nil
	case *wrapperspb.UInt32Value:
		return v.GetValue(), nil
	case *wrapperspb.FloatValue:
		return v.GetValue(), nil
	case *wrapperspb.DoubleValue:
		return v.GetValue(), nil
	case *wrapperspb.BoolValue:
		return v.GetValue(), nil
	case *wrapperspb.BytesValue:
		return v.GetValue(), nil
	case *timestamppb.Timestamp:
		return v.AsTime(), nil
	case *durationpb.Duration:
		return v.AsDuration(), nil
	case *structpb.Struct:
		return v.AsMap(), nil
	case *emptypb.Empty:
		return nil, nil
	default:
		return nil, errors.Errorf("unsupported well-known type in Any: %T", v)
	}
}

// convertArgsToAny converts Go values to protobuf Any types
func convertArgsToAny(args []any) []*anypb.Any {
	anyArgs := make([]*anypb.Any, len(args))
	for i, arg := range args {
		protoMsg, err := toProtoMessage(arg)
		if err != nil {
			panic(errors.Wrap(err, "failed to convert arg to proto message"))
		}
		anyArg, marshalErr := anypb.New(protoMsg)
		if marshalErr != nil {
			panic(errors.Wrap(marshalErr, "failed to marshal proto message to Any"))
		}
		anyArgs[i] = anyArg
	}
	return anyArgs
}

// Localized represents a localized message
type Localized struct {
	Key  string
	Args []any
}

func (l *Localized) Proto() *statusv1.Localized {
	if l == nil {
		return nil
	}
	return &statusv1.Localized{
		Key:  l.Key,
		Args: convertArgsToAny(l.Args),
	}
}

func LocalizedFromProto(pb *statusv1.Localized) *Localized {
	if pb == nil {
		return nil
	}
	args := pb.GetArgs()
	goArgs := make([]any, len(args))
	for i, anyArg := range args {
		val, err := extractValueFromAny(anyArg)
		if err != nil {
			// Use nil instead of placeholder strings to avoid breaking template rendering
			val = nil
			slog.Warn("failed to extract value from Any", "error", err)
		}
		goArgs[i] = val
	}
	return &Localized{
		Key:  pb.GetKey(),
		Args: goArgs,
	}
}

// FieldViolation represents a field-level validation violation with localization capability
//
// Priority order for localized messages:
//  1. LocalizedMessage (highest priority - pre-translated, ready to use)
//  2. Localized (lower priority - template that needs translation via interceptor)
type FieldViolation struct {
	Field            string
	Reason           string
	Description      string
	Localized        *Localized                   // Localization template (requires translation via interceptor)
	LocalizedMessage *errdetails.LocalizedMessage // Pre-translated message (ready to use)
}

// Proto converts FieldViolation to protobuf message
func (f *FieldViolation) Proto() *statusv1.BadRequest_FieldViolation {
	if f == nil {
		return nil
	}

	return &statusv1.BadRequest_FieldViolation{
		Field:            f.Field,
		Reason:           f.Reason,
		Description:      f.Description,
		Localized:        f.Localized.Proto(),
		LocalizedMessage: f.LocalizedMessage,
	}
}

// ToFieldViolations converts any error to field violations for the specified field.
// Returns a slice where the first element is the main field error, followed by
// nested field violations with the field path properly prefixed.
func ToFieldViolations(err error, field string) []*FieldViolation {
	if err == nil {
		return nil
	}

	s, _ := FromError(err)

	// Check for existing translated content in details
	var localizedMessage *errdetails.LocalizedMessage
	var badRequest *errdetails.BadRequest
	for _, d := range s.details {
		switch v := d.(type) {
		case *errdetails.LocalizedMessage:
			localizedMessage = v
		case *errdetails.BadRequest:
			badRequest = v
		}
	}

	// Main field error (always first)
	var mainLocalized *Localized
	if localizedMessage == nil {
		// Only use s.localized if no existing translated message
		mainLocalized = LocalizedFromProto(s.localized)
	}

	mainViolation := &FieldViolation{
		Field:            field,
		Reason:           s.Reason(),
		Description:      s.Message(),
		Localized:        mainLocalized,
		LocalizedMessage: localizedMessage,
	}

	result := []*FieldViolation{mainViolation}

	// Process field violations - use existingBadRequest if present, otherwise s.badRequest
	if badRequest != nil {
		// Use Google standard BadRequest from details
		for _, fv := range badRequest.FieldViolations {
			result = append(result, &FieldViolation{
				Field:            field + "." + fv.Field,
				Reason:           fv.Reason,
				Description:      fv.Description,
				Localized:        nil,
				LocalizedMessage: fv.LocalizedMessage,
			})
		}
	} else if s.badRequest != nil {
		// Use our custom badRequest
		for _, fv := range s.badRequest.FieldViolations {
			result = append(result, &FieldViolation{
				Field:            field + "." + fv.Field,
				Reason:           fv.Reason,
				Description:      fv.Description,
				Localized:        LocalizedFromProto(fv.Localized),
				LocalizedMessage: nil,
			})
		}
	}

	return result
}

// FlattenFieldViolations flattens various field violation types into a unified []*FieldViolation slice.
// Supports *FieldViolation, []*FieldViolation, and their protobuf equivalents.
// Mixed types are allowed in a single call.
//
// Note: For error and *Status inputs, use ToFieldViolations(err, field) or status.ToFieldViolations(field)
// first to specify the field name, then pass the result to this function.
func FlattenFieldViolations(inputs ...any) []*FieldViolation {
	var result []*FieldViolation

	for _, input := range inputs {
		switch v := input.(type) {
		case *FieldViolation:
			if v != nil {
				result = append(result, v)
			}
		case []*FieldViolation:
			result = append(result, v...)
		case *errdetails.BadRequest_FieldViolation:
			if v != nil {
				result = append(result, &FieldViolation{
					Field:            v.Field,
					Reason:           v.Reason,
					Description:      v.Description,
					LocalizedMessage: v.LocalizedMessage,
				})
			}
		case []*errdetails.BadRequest_FieldViolation:
			for _, pbFv := range v {
				if pbFv != nil {
					result = append(result, &FieldViolation{
						Field:            pbFv.Field,
						Reason:           pbFv.Reason,
						Description:      pbFv.Description,
						LocalizedMessage: pbFv.LocalizedMessage,
					})
				}
			}
		}
	}

	return result
}

// TranslateError translates error messages and field violations using the provided i18n instance and language.
// Returns the original error if translation is not possible or if localized details already exist.
//
// Translation priority:
//  1. If LocalizedMessage already exists -> skip translation (highest priority)
//  2. If Localized template exists -> translate template (medium priority)
//  3. Use error reason for translation -> fallback (lowest priority)
func TranslateError(ib *i18nx.I18N, lang language.Tag, err error) error {
	if err == nil {
		return nil
	}

	s, ok := FromError(err)
	if !ok {
		s.message = "unknown error"
	}

	var localizedMessage *errdetails.LocalizedMessage
	var badRequest *errdetails.BadRequest
	for _, d := range s.details {
		switch v := d.(type) {
		case *errdetails.LocalizedMessage:
			localizedMessage = v
		case *errdetails.BadRequest:
			badRequest = v
		}
	}

	if localizedMessage == nil {
		var text string

		if s.localized != nil && s.localized.GetKey() != "" {
			localized := LocalizedFromProto(s.localized)
			text = ib.Sprintf(lang, localized.Key, localized.Args...)
		}

		reason := s.Reason()
		if text == "" && reason != "" {
			text = ib.Sprintf(lang, reason)
		}

		if text != "" {
			s.details = append(s.details, &errdetails.LocalizedMessage{
				Locale:  lang.String(),
				Message: text,
			})
		}

		// Clear original localized information after translation to avoid duplication
		s.localized = nil
	}

	if badRequest == nil && s.badRequest != nil {
		br := &errdetails.BadRequest{}
		for _, fieldViolation := range s.badRequest.FieldViolations {
			fv := &errdetails.BadRequest_FieldViolation{
				Field:       fieldViolation.Field,
				Description: fieldViolation.Description,
				Reason:      fieldViolation.Reason,
			}

			// Check if LocalizedMessage already exists (highest priority)
			if fieldViolation.LocalizedMessage != nil {
				// Use existing pre-translated message directly
				fv.LocalizedMessage = proto.Clone(fieldViolation.LocalizedMessage).(*errdetails.LocalizedMessage)
			} else {
				// Need to translate from template or reason
				var text string

				if fieldViolation.Localized != nil && fieldViolation.Localized.GetKey() != "" {
					localized := LocalizedFromProto(fieldViolation.Localized)
					text = ib.Sprintf(lang, localized.Key, localized.Args...)
				}

				reason := fieldViolation.GetReason()
				if text == "" && reason != "" {
					text = ib.Sprintf(lang, reason)
				}

				if text != "" {
					fv.LocalizedMessage = &errdetails.LocalizedMessage{
						Locale:  lang.String(),
						Message: text,
					}
				}
			}

			br.FieldViolations = append(br.FieldViolations, fv)
		}
		s.details = append(s.details, br)

		// Clear original localized information after translation to avoid duplication
		s.badRequest = nil
	}

	return s.Err()
}

// TranslateStatusErrorOnly translates only StatusError types, returning the error and a boolean indicating success
func TranslateStatusErrorOnly(ib *i18nx.I18N, lang language.Tag, err error) (error, bool) {
	if err != nil {
		if se := new(StatusError); !errors.As(err, &se) {
			return err, false //nolint:errhandle
		}
	}
	return TranslateError(ib, lang, err), true
}

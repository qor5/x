package statusx

import (
	"cmp"
	"log/slog"
	"slices"
	"time"

	"github.com/pkg/errors"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/qor5/x/v3/i18nx"
)

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
	key  string
	args []any
}

func (l *Localized) Key() string {
	if l == nil {
		return ""
	}
	return l.key
}

func (l *Localized) Args() []any {
	if l == nil {
		return nil
	}
	return slices.Clone(l.args)
}

func (l *Localized) Proto() *statusv1.Localized {
	if l == nil {
		return nil
	}
	return &statusv1.Localized{
		Key:  l.key,
		Args: convertArgsToAny(l.args),
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
			slog.Warn("failed to extract value from Any", "error", err, "index", i, "typeURL", anyArg.GetTypeUrl())
		}
		goArgs[i] = val
	}
	return &Localized{
		key:  pb.GetKey(),
		args: goArgs,
	}
}

// TranslateError translates error messages and field violations using the provided i18n instance and language.
// Returns the original error if translation is not possible or if localized details already exist.
//
// Translation priority:
//  1. If LocalizedMessage already exists -> skip translation (highest priority)
//  2. If Localized template exists -> translate template (medium priority)
//  3. Use error reason for translation -> fallback (lowest priority)
func TranslateError(err error, ib *i18nx.I18N, lang language.Tag) error {
	if err == nil {
		return nil
	}

	s, ok := FromError(err)
	if !ok {
		s.message = "unknown error"
	}

	// Use the new Status method for translation
	translatedStatus := s.Translated(ib, lang)
	return translatedStatus.Err()
}

// TranslateStatusErrorOnly translates only StatusError types, returning the error and a boolean indicating success
func TranslateStatusErrorOnly(err error, ib *i18nx.I18N, lang language.Tag) (error, bool) {
	if err != nil {
		var se *StatusError
		if !errors.As(err, &se) {
			return err, false //nolint:errhandle
		}
	}
	return TranslateError(err, ib, lang), true
}

// Translated returns a new Status with translated messages and field violations.
//
// Translation priority:
//  1. If LocalizedMessage already exists -> skip translation (highest priority)
//  2. If Localized template exists -> translate template (medium priority)
//  3. Use error reason for translation -> fallback (lowest priority)
func (s *Status) Translated(ib *i18nx.I18N, lang language.Tag) *Status {
	if s == nil {
		return nil
	}

	st := Clone(s)
	st.translateMainMessage(ib, lang)
	st.translateFieldViolations(ib, lang)
	return st
}

// translateMainMessage handles the translation of the main status message
func (s *Status) translateMainMessage(ib *i18nx.I18N, lang language.Tag) {
	// Check if LocalizedMessage already exists
	for _, d := range s.details {
		if _, ok := d.(*errdetails.LocalizedMessage); ok {
			return // Already translated
		}
	}

	defer func() {
		s.localized = nil // Clear to avoid duplication
	}()

	var text string

	localized := cmp.Or(LocalizedFromProto(s.localized), &Localized{})
	if localized.key == "" {
		localized.key = s.Reason()
	}
	if localized.key != "" {
		text = ib.Sprintf(lang, localized.key, localized.args...)
	}

	if text != "" {
		s.details = append(s.details, &errdetails.LocalizedMessage{
			Locale:  lang.String(),
			Message: text,
		})
	}
}

// translateFieldViolations handles the translation of field violations
func (s *Status) translateFieldViolations(ib *i18nx.I18N, lang language.Tag) {
	if s.badRequest == nil || len(s.badRequest.FieldViolations) == 0 {
		return
	}

	// Check if BadRequest already exists in details
	for _, d := range s.details {
		if _, ok := d.(*errdetails.BadRequest); ok {
			return // Already translated
		}
	}

	defer func() {
		s.badRequest = nil // Clear to avoid duplication
	}()

	br := &errdetails.BadRequest{}

	for _, fieldViolation := range s.badRequest.FieldViolations {
		fv := &errdetails.BadRequest_FieldViolation{
			Field:       fieldViolation.Field,
			Description: fieldViolation.Description,
			Reason:      fieldViolation.Reason,
		}

		// Check if LocalizedMessage already exists (highest priority)
		if fieldViolation.LocalizedMessage != nil {
			fv.LocalizedMessage = proto.Clone(fieldViolation.LocalizedMessage).(*errdetails.LocalizedMessage)
		} else {
			var text string

			localized := cmp.Or(LocalizedFromProto(fieldViolation.Localized), &Localized{})
			if localized.key == "" {
				localized.key = fieldViolation.GetReason()
			}
			if localized.key != "" {
				text = ib.Sprintf(lang, localized.key, localized.args...)
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
}

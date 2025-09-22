package statusx

import (
	"cmp"
	"fmt"
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

// FieldViolation represents a field-level validation violation with localization capability
//
// Priority order for localized messages:
//  1. LocalizedMessage (highest priority - pre-translated, ready to use)
//  2. Localized (lower priority - template that needs translation via interceptor)
type FieldViolation struct {
	field            string
	reason           string
	description      string
	localized        *Localized                   // Localization template (requires translation via interceptor)
	localizedMessage *errdetails.LocalizedMessage // Pre-translated message (ready to use)
}

type FieldViolations []*FieldViolation

func (fvs FieldViolations) PrependField(field string) FieldViolations {
	return PrependField(field, fvs...)
}

// Field returns the field name that caused the violation.
func (f *FieldViolation) Field() string {
	if f == nil {
		return ""
	}
	return f.field
}

// Reason returns the error reason code.
func (f *FieldViolation) Reason() string {
	if f == nil {
		return ""
	}
	return f.reason
}

// Description returns the human-readable description of the violation.
func (f *FieldViolation) Description() string {
	if f == nil {
		return ""
	}
	return f.description
}

// Localized returns the localization template if set.
// Returns nil if no localization template is available.
func (f *FieldViolation) Localized() *Localized {
	if f.localized == nil {
		return nil
	}
	// Return a copy to prevent external modification
	return &Localized{
		key:  f.localized.Key(),
		args: f.localized.Args(),
	}
}

// LocalizedMessage returns the pre-translated message if available.
// Returns nil if no pre-translated message is set.
func (f *FieldViolation) LocalizedMessage() *errdetails.LocalizedMessage {
	if f.localizedMessage == nil {
		return nil
	}
	return proto.Clone(f.localizedMessage).(*errdetails.LocalizedMessage)
}

// NewFieldViolation creates a new field validation violation.
// The reason serves as the error identifier and will be used as the i18n key fallback during translation.
func NewFieldViolation(field, reason, description string) *FieldViolation {
	if field == "" {
		panic("field is required")
	}
	if reason == "" {
		panic("reason is required")
	}
	return &FieldViolation{
		field:       field,
		reason:      reason,
		description: description,
		localized:   &Localized{key: reason},
	}
}

// NewFieldViolationf creates a new field validation violation with a formatted description.
func NewFieldViolationf(field, reason, format string, args ...any) *FieldViolation {
	return NewFieldViolation(field, reason, fmt.Sprintf(format, args...))
}

// WithLocalized sets a custom i18n key and template arguments.
// This sets a specific i18n key instead of relying on the reason as fallback during translation.
func (f *FieldViolation) WithLocalized(key string, args ...any) *FieldViolation {
	if key == "" {
		panic("key is required")
	}
	return &FieldViolation{
		field:            f.Field(),
		reason:           f.Reason(),
		description:      f.Description(),
		localized:        &Localized{key: key, args: args},
		localizedMessage: f.LocalizedMessage(),
	}
}

// WithLocalizedArgs sets template arguments for i18n.
// Preserves the existing localized key if present, or leaves it empty for the translator to use reason as fallback.
// This is useful when you want to add template arguments without setting a specific i18n key.
func (f *FieldViolation) WithLocalizedArgs(args ...any) *FieldViolation {
	return f.WithLocalized(f.localized.Key(), args...)
}

// Proto converts FieldViolation to protobuf message
func (f *FieldViolation) Proto() *statusv1.BadRequest_FieldViolation {
	if f == nil {
		return nil
	}

	return &statusv1.BadRequest_FieldViolation{
		Field:            f.Field(),
		Reason:           f.Reason(),
		Description:      f.Description(),
		Localized:        f.Localized().Proto(),
		LocalizedMessage: f.LocalizedMessage(),
	}
}

func PrependField(field string, fvs ...*FieldViolation) FieldViolations {
	for _, fv := range fvs {
		fv.field = field + "." + fv.field
	}
	return fvs
}

// ToFieldViolations converts any error to field violations for the specified field.
// Simple behavior:
//   - If field is empty: returns only nested field violations without prefix
//   - If field is non-empty: returns only nested field violations with the specified field prefix
//
// This design extracts meaningful field-level violations from container errors.
func ToFieldViolations(err error, field string) FieldViolations {
	if err == nil {
		return nil
	}

	s, ok := FromError(err)
	if !ok {
		s.message = "unknown error"
	}

	// Check for existing translated content in details
	var localizedMessage *errdetails.LocalizedMessage
	var badRequest *errdetails.BadRequest
	for _, d := range s.details {
		switch v := d.(type) {
		case *errdetails.LocalizedMessage:
			localizedMessage = proto.Clone(v).(*errdetails.LocalizedMessage)
		case *errdetails.BadRequest:
			badRequest = proto.Clone(v).(*errdetails.BadRequest)
		}
	}

	// Main field error (skip if field is empty)
	var result []*FieldViolation
	if field != "" {
		// Include main error as first field violation
		var mainLocalized *Localized
		if localizedMessage == nil {
			// Only use s.localized if no existing translated message
			mainLocalized = LocalizedFromProto(s.localized)
		}

		result = []*FieldViolation{{
			field:            field,
			reason:           s.Reason(),
			description:      s.Message(),
			localized:        mainLocalized,
			localizedMessage: localizedMessage,
		}}
	}

	// Process field violations - use existingBadRequest if present, otherwise s.badRequest
	var fieldPrefix string
	if field != "" {
		fieldPrefix = field + "."
	}
	if badRequest != nil {
		// Use Google standard BadRequest from details
		for _, fv := range badRequest.FieldViolations {
			result = append(result, &FieldViolation{
				field:            fieldPrefix + fv.Field,
				reason:           fv.Reason,
				description:      fv.Description,
				localized:        nil,
				localizedMessage: fv.LocalizedMessage,
			})
		}
	} else if s.badRequest != nil {
		// Use our custom badRequest
		for _, fv := range s.badRequest.FieldViolations {
			result = append(result, &FieldViolation{
				field:            fieldPrefix + fv.Field,
				reason:           fv.Reason,
				description:      fv.Description,
				localized:        LocalizedFromProto(fv.Localized),
				localizedMessage: nil,
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
func FlattenFieldViolations(inputs ...any) ([]*FieldViolation, error) {
	var result []*FieldViolation

	for _, input := range inputs {
		if input == nil {
			continue // Skip nil inputs
		}
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
					field:            v.Field,
					reason:           v.Reason,
					description:      v.Description,
					localizedMessage: v.LocalizedMessage,
				})
			}
		case []*errdetails.BadRequest_FieldViolation:
			for _, pbFv := range v {
				if pbFv != nil {
					result = append(result, &FieldViolation{
						field:            pbFv.Field,
						reason:           pbFv.Reason,
						description:      pbFv.Description,
						localizedMessage: pbFv.LocalizedMessage,
					})
				}
			}
		default:
			return nil, errors.Errorf("unsupported type for flatten field violations: %T", v)
		}
	}

	return result, nil
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
		if se := new(StatusError); !errors.As(err, &se) {
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

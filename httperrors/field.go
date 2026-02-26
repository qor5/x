package httperrors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// ValidationError creates a new Status with http.StatusUnprocessableEntity (422) and a flattened list of field violations.
func ValidationError(inputs ...any) *Status {
	violations, err := FlattenFieldViolations(inputs...)
	if err != nil {
		panic(err)
	}
	if len(violations) == 0 {
		return New(http.StatusOK, ReasonOK, "ok")
	}
	return New(http.StatusUnprocessableEntity, ReasonInvalidArgument, "invalid argument").WithFieldViolations(violations...)
}

// FormatField formats a dotted field path by applying a formatting function to each segment
// while preserving array index notations (e.g., [0], [1]).
//
// Parameters:
//   - field: The original field path (e.g., "user_info.addresses[0].street_name")
//   - formatFunc: Function to apply to each field segment (e.g., lo.CamelCase)
//
// Returns:
//   - The formatted field path with proper array index preservation
//
// Example:
//
//	FormatField("user_info.addresses[0].street_name", lo.CamelCase)
//	Returns: "userInfo.addresses[0].streetName"
func FormatField(field string, formatFunc func(string) string) string {
	parts := strings.Split(field, ".")
	for i, part := range parts {
		// Find the start position of array index notation (e.g., [0], [1])
		// Use Index (not LastIndex) to handle multiple array indices correctly
		// For "matrix_data[1][2]", we want to split at the first "[", not the last one
		arrStart := strings.Index(part, "[")
		if arrStart != -1 {
			// We cannot apply formatFunc(part) directly because it would:
			// 1. Remove the square brackets: "user_addresses[0]" -> formatted result without brackets
			// 2. Make array indices indistinguishable from field names
			// 3. Break client-side error handling that relies on proper array notation
			//
			// Example of incorrect direct usage:
			//   formatFunc("user_addresses[0]") -> "userAddresses0"  // Wrong!
			//
			// Example of current approach:
			//   formatFunc("user_addresses") + "[0]" -> "userAddresses[0]"  // Correct!
			parts[i] = formatFunc(part[:arrStart]) + part[arrStart:]
		} else {
			// For regular fields without array indices, simply apply the formatting function
			parts[i] = formatFunc(part)
		}
	}
	return strings.Join(parts, ".")
}

// PrependField prepends a field name to the field name of each field violation.
func PrependField(field string, fvs ...*FieldViolation) FieldViolations {
	for _, fv := range fvs {
		fv.field = field + "." + fv.field
	}
	return fvs
}

// FieldViolation represents a field-level validation violation with localization capability
//
// Priority order for localized messages:
//  1. LocalizedMessage (highest priority - pre-translated, ready to use)
//  2. Localized (lower priority - template that needs translation via middleware)
type FieldViolation struct {
	field            string
	reason           string
	description      string
	localized        *Localized        // Localization template (requires translation via middleware)
	localizedMessage *LocalizedMessage // Pre-translated message (ready to use)
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

// GetLocalized returns the localization template if set.
// Returns nil if no localization template is available.
func (f *FieldViolation) GetLocalized() *Localized {
	if f == nil || f.localized == nil {
		return nil
	}
	return f.localized.Clone()
}

// GetLocalizedMessage returns the pre-translated message if available.
// Returns nil if no pre-translated message is set.
func (f *FieldViolation) GetLocalizedMessage() *LocalizedMessage {
	if f == nil || f.localizedMessage == nil {
		return nil
	}
	return f.localizedMessage.Clone()
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
		localizedMessage: f.GetLocalizedMessage(),
	}
}

// WithLocalizedArgs sets template arguments for i18n.
// Preserves the existing localized key if present, or leaves it empty for the translator to use reason as fallback.
// This is useful when you want to add template arguments without setting a specific i18n key.
func (f *FieldViolation) WithLocalizedArgs(args ...any) *FieldViolation {
	return f.WithLocalized(f.localized.Key(), args...)
}

// Clone creates a deep copy of this FieldViolation.
func (f *FieldViolation) Clone() *FieldViolation {
	if f == nil {
		return nil
	}
	return &FieldViolation{
		field:            f.field,
		reason:           f.reason,
		description:      f.description,
		localized:        f.localized.Clone(),
		localizedMessage: f.localizedMessage.Clone(),
	}
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

	// Main field error (skip if field is empty)
	var result FieldViolations
	if field != "" {
		result = append(result, &FieldViolation{
			field:       field,
			reason:      s.Reason(),
			description: s.Message(),
			localized:   s.Localized(),
		})
	}

	// Process field violations
	var fieldPrefix string
	if field != "" {
		fieldPrefix = field + "."
	}
	for _, fv := range s.fieldViolations {
		result = append(result, &FieldViolation{
			field:            fieldPrefix + fv.field,
			reason:           fv.reason,
			description:      fv.description,
			localized:        fv.localized.Clone(),
			localizedMessage: fv.localizedMessage.Clone(),
		})
	}

	return result
}

// FlattenFieldViolations flattens various field violation types into a unified FieldViolations slice.
// Supports *FieldViolation, []*FieldViolation, FieldViolations.
// Mixed types are allowed in a single call.
//
// Note: For error and *Status inputs, use ToFieldViolations(err, field) or status.ToFieldViolations(field)
// first to specify the field name, then pass the result to this function.
func FlattenFieldViolations(inputs ...any) (FieldViolations, error) {
	var result FieldViolations

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
		case FieldViolations:
			result = append(result, v...)
		default:
			return nil, errors.Errorf("unsupported type for flatten field violations: %T", v)
		}
	}

	return result, nil
}

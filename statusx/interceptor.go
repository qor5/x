package statusx

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	"github.com/qor5/x/v3/i18nx"
	statusv1 "github.com/qor5/x/v3/statusx/gen/status/v1"
)

func TranslateStatusErrorOnly(ctx context.Context, ib *i18nx.I18N, err error) (error, bool) {
	if err != nil {
		if se := new(StatusError); !errors.As(err, &se) {
			return err, false //nolint:errhandle
		}
	}
	return TranslateError(ctx, ib, err), true
}

// TranslateError translates error details according to the locale in the context.
//
// Given an error, it translates the error message and field violations into the
// locale specified in the context. If the error doesn't contain any localized
// message or field violations, it falls back to the error reason. If the error
// doesn't contain any localized message but has a localized key, it uses the key
// to translate the message. If the error doesn't contain any localized message or
// key, it leaves the error as is.
//
// The function returns the translated error. If the error is nil, it returns nil.
//
// For example, if a request is invalid, the error will contain a
// *errdetails.BadRequest error detail. The field violations in the bad request
// will be translated according to the locale specified in the context. The
// translated field violations will be returned as a new *errdetails.LocalizedMessage
// error detail.
func TranslateError(ctx context.Context, ib *i18nx.I18N, err error) error {
	if err == nil {
		return nil
	}

	s, ok := FromError(err)
	if !ok {
		s.message = "unknown error"
	}

	tag := ib.LanguageFromContext(ctx)
	locale := tag.String()

	var badRequest *errdetails.BadRequest
	var localizedMessage *errdetails.LocalizedMessage
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

		// Handle top-level localization from our proto
		if s.proto != nil && s.proto.Localized != nil {
			localized := s.proto.Localized
			if localized.GetKey() != "" {
				args := localized.GetArgs()
				// Convert protobuf Any args back to Go values for i18n formatting
				goArgs := make([]any, len(args))
				for i, anyArg := range args {
					val, err := extractValueFromAny(anyArg)
					if err != nil {
						val = fmt.Sprintf("<%s>", anyArg.GetTypeUrl())
					}
					goArgs[i] = val
				}
				text = ib.Sprintf(tag, localized.GetKey(), goArgs...)
			}
		}

		if text == "" {
			text = ib.Sprintf(tag, s.Reason())
		}

		if text != "" {
			s.details = append(s.details, &errdetails.LocalizedMessage{
				Locale:  locale,
				Message: text,
			})
		}
	}

	// Only construct BadRequest if none exists AND we have field violations
	if badRequest == nil && s.proto != nil && len(s.proto.GetFieldViolations()) > 0 {
		// Create new BadRequest from field violations
		fieldViolations := make([]*errdetails.BadRequest_FieldViolation, 0)

		for _, fieldLoc := range s.proto.GetFieldViolations() {
			fieldViolation := &errdetails.BadRequest_FieldViolation{
				Field:       fieldLoc.GetField(),
				Description: fieldLoc.GetDescription(),
				Reason:      fieldLoc.GetReason(),
			}

			// Handle field-level localization if present
			if fieldLoc.GetLocalized() != nil {
				localized := fieldLoc.GetLocalized()
				args := localized.GetArgs()

				// Convert protobuf Any args back to Go values for template rendering
				goArgs := make([]any, len(args))
				for i, anyArg := range args {
					val, err := extractValueFromAny(anyArg)
					if err != nil {
						val = fmt.Sprintf("<%s>", anyArg.GetTypeUrl())
					}
					goArgs[i] = val
				}

				// Render localized message
				localizedText := ib.Sprintf(tag, localized.GetKey(), goArgs...)
				if localizedText != "" {
					fieldViolation.LocalizedMessage = &errdetails.LocalizedMessage{
						Locale:  locale,
						Message: localizedText,
					}
				}
			}

			fieldViolations = append(fieldViolations, fieldViolation)
		}

		if len(fieldViolations) > 0 {
			s.details = append(s.details, &errdetails.BadRequest{
				FieldViolations: fieldViolations,
			})
		}

		// Clear field violations as they're now converted to BadRequest
		if s.proto != nil {
			s.proto.FieldViolations = nil
		}

		return s.Err()
	}

	// If badRequest already exists, only update existing FieldViolations that lack LocalizedMessage
	if badRequest != nil {
		// Create field violation map for quick lookup
		fieldViolationMap := make(map[string]*statusv1.FieldViolation)
		if s.proto != nil {
			for _, fieldLoc := range s.proto.GetFieldViolations() {
				fieldViolationMap[fieldLoc.GetField()] = fieldLoc
			}
		}

		for _, fv := range badRequest.GetFieldViolations() {
			if fv.GetLocalizedMessage() != nil {
				continue
			}

			// Check if we have field-specific localization
			if fieldLoc, exists := fieldViolationMap[fv.GetField()]; exists {
				if fieldLoc.GetLocalized() != nil {
					localized := fieldLoc.GetLocalized()
					args := localized.GetArgs()

					// Convert protobuf Any args back to Go values for template rendering
					goArgs := make([]any, len(args))
					for i, anyArg := range args {
						val, err := extractValueFromAny(anyArg)
						if err != nil {
							val = fmt.Sprintf("<%s>", anyArg.GetTypeUrl())
						}
						goArgs[i] = val
					}

					// Use i18nx template rendering with field-specific key and args
					text := ib.Sprintf(tag, localized.GetKey(), goArgs...)
					if text != "" {
						fv.LocalizedMessage = &errdetails.LocalizedMessage{
							Locale:  locale,
							Message: text,
						}
						continue
					}
				}
			}

			// Fallback: use reason-based localization (existing behavior)
			text := ib.Sprintf(tag, fv.GetReason())
			if text != "" {
				fv.LocalizedMessage = &errdetails.LocalizedMessage{
					Locale:  locale,
					Message: text,
				}
			}
		}

		// Clear field violations as they're now processed
		if s.proto != nil && len(s.proto.GetFieldViolations()) > 0 {
			s.proto.FieldViolations = nil
		}
	}

	return s.Err()
}

func ConvertToFieldViolation(st *status.Status, field string) *errdetails.BadRequest_FieldViolation {
	details := st.Details()
	errorInfo := ExtractDetail[*errdetails.ErrorInfo](details)
	localizedMessage := ExtractDetail[*errdetails.LocalizedMessage](details)
	return &errdetails.BadRequest_FieldViolation{
		Field:            field,
		Description:      st.Message(),
		Reason:           errorInfo.GetReason(),
		LocalizedMessage: localizedMessage,
	}
}

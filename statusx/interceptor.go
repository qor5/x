package statusx

import (
	"github.com/pkg/errors"
	"golang.org/x/text/language"
	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"github.com/qor5/x/v3/i18nx"
)

func TranslateStatusErrorOnly(ib *i18nx.I18N, lang language.Tag, err error) (error, bool) {
	if err != nil {
		if se := new(StatusError); !errors.As(err, &se) {
			return err, false //nolint:errhandle
		}
	}
	return TranslateError(ib, lang, err), true
}

// TranslateError translates error messages and field violations using the provided i18n instance and language.
// Returns the original error if translation is not possible or if localized details already exist.
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

		if s.localized.GetKey() != "" {
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
	}

	if badRequest == nil && s.badRequest != nil {
		br := &errdetails.BadRequest{}
		for _, fieldViolation := range s.badRequest.FieldViolations {
			fv := &errdetails.BadRequest_FieldViolation{
				Field:       fieldViolation.Field,
				Description: fieldViolation.Description,
				Reason:      fieldViolation.Reason,
			}

			var text string

			if fieldViolation.Localized.GetKey() != "" {
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

			br.FieldViolations = append(br.FieldViolations, fv)
		}
		s.details = append(s.details, br)
	}

	return s.Err()
}

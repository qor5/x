package statusx

import (
	"context"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"github.com/qor5/x/v3/i18nx"
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
	for _, d := range s.extraDetails {
		switch v := d.(type) {
		case *errdetails.LocalizedMessage:
			localizedMessage = v
		case *errdetails.BadRequest:
			badRequest = v
		}
	}

	if localizedMessage == nil {
		var text string

		localized := s.localized
		if localized.GetKey() != "" {
			args := localized.GetArgs()
			text = ib.Sprintf(tag, localized.GetKey(), lo.Map(args, func(v string, _ int) any {
				return v
			})...)
		}
		s.localized = nil

		if text == "" {
			text = ib.Sprintf(tag, s.Reason())
		}

		if text != "" {
			s.extraDetails = append(s.extraDetails, &errdetails.LocalizedMessage{
				Locale:  locale,
				Message: text,
			})
		}
	}

	if badRequest == nil {
		return s.Err()
	}

	for _, fv := range badRequest.GetFieldViolations() {
		if fv.GetLocalizedMessage() != nil {
			continue
		}
		text := ib.Sprintf(tag, fv.GetReason())
		if text != "" {
			fv.LocalizedMessage = &errdetails.LocalizedMessage{
				Locale:  locale,
				Message: text,
			}
		}
	}

	return s.Err()
}

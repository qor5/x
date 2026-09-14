package httperrors

import (
	"cmp"
	"slices"

	"github.com/pkg/errors"
	"golang.org/x/text/language"

	"github.com/qor5/x/v3/i18nx"
)

// Localized represents a localization template with a key and optional arguments.
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

func (l *Localized) Clone() *Localized {
	if l == nil {
		return nil
	}
	return &Localized{
		key:  l.key,
		args: slices.Clone(l.args),
	}
}

// LocalizedMessage represents a pre-translated message with its locale.
type LocalizedMessage struct {
	Locale  string
	Message string
}

func (lm *LocalizedMessage) Clone() *LocalizedMessage {
	if lm == nil {
		return nil
	}
	return &LocalizedMessage{
		Locale:  lm.Locale,
		Message: lm.Message,
	}
}

// TranslateError translates error messages and field violations using the provided i18n instance and language.
// It always returns an error derived from a status representation of the input error (typically a *StatusError);
//
// Translation behavior:
//   - The original message is preserved in the `message` field.
//   - The translated text is stored in the `localizedMessage` field.
//   - If the Localized template is set, its key and args are used for translation.
//   - Otherwise, the error reason is used as the i18n key fallback.
//   - If already translated (localized is nil), translation is skipped.
func TranslateError(err error, ib *i18nx.I18N, lang language.Tag) error {
	if err == nil {
		return nil
	}

	s, ok := FromError(err)
	if !ok {
		s.message = "unknown error"
	}

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

// Translated returns a new Status with translated messages stored in localizedMessage fields.
// The original message is preserved unchanged.
func (s *Status) Translated(ib *i18nx.I18N, lang language.Tag) *Status {
	if s == nil {
		return nil
	}

	st := Clone(s)
	st.translateMainMessage(ib, lang)
	st.translateFieldViolations(ib, lang)
	return st
}

// translateMainMessage handles the translation of the main status message.
// The original message is preserved; the translated text is stored in localizedMessage.
func (s *Status) translateMainMessage(ib *i18nx.I18N, lang language.Tag) {
	if s.localizedMessage != nil {
		return // Already translated
	}
	if s.localized == nil {
		return // No template
	}

	localized := cmp.Or(s.localized, &Localized{})
	if localized.key == "" {
		localized.key = s.Reason()
	}

	var text string
	if localized.key != "" {
		text = ib.Sprintf(lang, localized.key, localized.args...)
	}

	if text != "" {
		s.localizedMessage = &LocalizedMessage{
			Locale:  lang.String(),
			Message: text,
		}
	}

	s.localized = nil // Clear to indicate translation is done
}

// translateFieldViolations handles the translation of field violations.
// The original description is preserved; the translated text is stored in localizedMessage.
func (s *Status) translateFieldViolations(ib *i18nx.I18N, lang language.Tag) {
	if len(s.fieldViolations) == 0 {
		return
	}

	for _, fv := range s.fieldViolations {
		if fv.localizedMessage != nil {
			continue // Already translated
		}
		if fv.localized == nil {
			continue // No template
		}

		localized := cmp.Or(fv.localized, &Localized{})
		if localized.key == "" {
			localized.key = fv.Reason()
		}

		var text string
		if localized.key != "" {
			text = ib.Sprintf(lang, localized.key, localized.args...)
		}

		if text != "" {
			fv.localizedMessage = &LocalizedMessage{
				Locale:  lang.String(),
				Message: text,
			}
		}

		fv.localized = nil // Clear to indicate translation is done
	}
}

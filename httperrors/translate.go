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

// LocalizedMessage represents a pre-translated message.
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
// Returns the original error if translation is not possible or if localized details already exist.
//
// Translation priority:
//  1. If localizedMessage already exists -> skip translation (highest priority)
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
//  1. If localizedMessage already exists -> skip translation (highest priority)
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
	// Check if already translated (localizedMessage already set via translation)
	// For Status, we track this by checking if localized has been consumed (nil).
	// After translation, localized is cleared to indicate it's been processed.
	if s.localized == nil {
		return // Already translated or no template
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
		s.message = text
	}

	s.localized = nil // Clear to indicate translation is done
}

// translateFieldViolations handles the translation of field violations
func (s *Status) translateFieldViolations(ib *i18nx.I18N, lang language.Tag) {
	if len(s.fieldViolations) == 0 {
		return
	}

	for _, fv := range s.fieldViolations {
		// Check if LocalizedMessage already exists (highest priority)
		if fv.localizedMessage != nil {
			continue
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

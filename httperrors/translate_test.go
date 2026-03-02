package httperrors

import (
	"net/http"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/qor5/x/v3/i18nx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

func newTestI18N(t *testing.T) *i18nx.I18N {
	t.Helper()
	csv := `key,en,zh
NOT_FOUND,Not Found,未找到
INTERNAL,Internal Error,内部错误
REQUIRED,Required,必填
TOO_SHORT,Too Short,太短
INVALID_FORMAT,Invalid Format,格式无效
custom.greeting,Hello %s,你好 %s
`
	ib, err := i18nx.New(strings.NewReader(csv))
	require.NoError(t, err)
	return ib
}

func TestLocalized(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		l := &Localized{key: "test_key", args: []any{"arg1", 42}}
		assert.Equal(t, "test_key", l.Key())
		assert.Equal(t, []any{"arg1", 42}, l.Args())
	})

	t.Run("nil safety", func(t *testing.T) {
		var l *Localized
		assert.Equal(t, "", l.Key())
		assert.Nil(t, l.Args())
		assert.Nil(t, l.Clone())
	})

	t.Run("clone independence", func(t *testing.T) {
		l := &Localized{key: "key", args: []any{"a", "b"}}
		cloned := l.Clone()
		assert.Equal(t, l.Key(), cloned.Key())
		assert.Equal(t, l.Args(), cloned.Args())

		// Modify clone args should not affect original
		cloned.args[0] = "modified"
		assert.Equal(t, "a", l.args[0])
	})
}

func TestLocalizedMessage(t *testing.T) {
	t.Run("clone", func(t *testing.T) {
		lm := &LocalizedMessage{Locale: "zh", Message: "你好"}
		cloned := lm.Clone()
		assert.Equal(t, lm.Locale, cloned.Locale)
		assert.Equal(t, lm.Message, cloned.Message)
	})

	t.Run("nil clone", func(t *testing.T) {
		var lm *LocalizedMessage
		assert.Nil(t, lm.Clone())
	})
}

func TestTranslateError(t *testing.T) {
	ib := newTestI18N(t)

	t.Run("nil error", func(t *testing.T) {
		result := TranslateError(nil, ib, language.English)
		assert.Nil(t, result)
	})

	t.Run("translate main message by reason", func(t *testing.T) {
		err := Error(http.StatusNotFound, "NOT_FOUND", "user not found")
		result := TranslateError(err, ib, language.English)
		require.NotNil(t, result)

		st := Convert(result)
		assert.Equal(t, "user not found", st.Message()) // original preserved
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "Not Found", st.GetLocalizedMessage().Message)
	})

	t.Run("translate to Chinese", func(t *testing.T) {
		err := Error(http.StatusNotFound, "NOT_FOUND", "user not found")
		result := TranslateError(err, ib, language.Chinese)
		require.NotNil(t, result)

		st := Convert(result)
		assert.Equal(t, "user not found", st.Message()) // original preserved
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "未找到", st.GetLocalizedMessage().Message)
		assert.Equal(t, "zh", st.GetLocalizedMessage().Locale)
	})

	t.Run("translate with custom localized key", func(t *testing.T) {
		err := New(http.StatusBadRequest, "CUSTOM", "bad request").
			WithLocalized("NOT_FOUND").Err()
		result := TranslateError(err, ib, language.English)

		st := Convert(result)
		assert.Equal(t, "bad request", st.Message()) // original preserved
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "Not Found", st.GetLocalizedMessage().Message)
	})

	t.Run("translate field violations", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "email is required")
		err := New(http.StatusUnprocessableEntity, ReasonInvalidArgument, "invalid").
			WithFieldViolations(fv).Err()

		result := TranslateError(err, ib, language.Chinese)
		require.NotNil(t, result)

		st := Convert(result)
		fvs := st.FieldViolations()
		require.Len(t, fvs, 1)
		assert.Equal(t, "email is required", fvs[0].Description()) // original preserved
		require.NotNil(t, fvs[0].GetLocalizedMessage())
		assert.Equal(t, "必填", fvs[0].GetLocalizedMessage().Message)
		assert.Equal(t, "zh", fvs[0].GetLocalizedMessage().Locale)
	})

	t.Run("field violation with custom localized key", func(t *testing.T) {
		fv := NewFieldViolation("name", "TOO_SHORT", "name is too short").
			WithLocalized("REQUIRED")
		err := New(http.StatusUnprocessableEntity, ReasonInvalidArgument, "invalid").
			WithFieldViolations(fv).Err()

		result := TranslateError(err, ib, language.English)
		st := Convert(result)
		fvs := st.FieldViolations()
		require.Len(t, fvs, 1)
		assert.Equal(t, "name is too short", fvs[0].Description()) // original preserved
		require.NotNil(t, fvs[0].GetLocalizedMessage())
		assert.Equal(t, "Required", fvs[0].GetLocalizedMessage().Message)
	})

	t.Run("skip already translated field violation", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "email is required")
		fv.localizedMessage = &LocalizedMessage{Locale: "fr", Message: "Obligatoire"}

		err := New(http.StatusUnprocessableEntity, ReasonInvalidArgument, "invalid").
			WithFieldViolations(fv).Err()

		result := TranslateError(err, ib, language.English)
		st := Convert(result)
		fvs := st.FieldViolations()
		require.Len(t, fvs, 1)
		// Should preserve the existing translation
		assert.Equal(t, "Obligatoire", fvs[0].GetLocalizedMessage().Message)
		assert.Equal(t, "fr", fvs[0].GetLocalizedMessage().Locale)
	})

	t.Run("idempotent field violation translation", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "email is required")
		err := New(http.StatusUnprocessableEntity, ReasonInvalidArgument, "invalid").
			WithFieldViolations(fv).Err()

		// First translation
		result := TranslateError(err, ib, language.English)
		// Second translation should not change anything
		result2 := TranslateError(result, ib, language.English)
		st := Convert(result2)
		fvs := st.FieldViolations()
		require.Len(t, fvs, 1)
		require.NotNil(t, fvs[0].GetLocalizedMessage())
		assert.Equal(t, "Required", fvs[0].GetLocalizedMessage().Message)
	})

	t.Run("unknown error", func(t *testing.T) {
		err := errors.New("some random error")
		result := TranslateError(err, ib, language.English)
		require.NotNil(t, result)

		st := Convert(result)
		assert.Equal(t, http.StatusInternalServerError, st.StatusCode())
	})
}

func TestTranslateStatusErrorOnly(t *testing.T) {
	ib := newTestI18N(t)

	t.Run("translates StatusError", func(t *testing.T) {
		err := Error(http.StatusNotFound, "NOT_FOUND", "not found")
		result, ok := TranslateStatusErrorOnly(err, ib, language.English)
		assert.True(t, ok)
		require.NotNil(t, result)

		st := Convert(result)
		assert.Equal(t, "not found", st.Message()) // original preserved
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "Not Found", st.GetLocalizedMessage().Message)
	})

	t.Run("does not translate non-StatusError", func(t *testing.T) {
		err := errors.New("plain error")
		result, ok := TranslateStatusErrorOnly(err, ib, language.English)
		assert.False(t, ok)
		assert.Equal(t, err, result)
	})

	t.Run("nil error", func(t *testing.T) {
		result, ok := TranslateStatusErrorOnly(nil, ib, language.English)
		assert.True(t, ok)
		assert.Nil(t, result)
	})
}

func TestTranslated(t *testing.T) {
	ib := newTestI18N(t)

	t.Run("nil status", func(t *testing.T) {
		var s *Status
		assert.Nil(t, s.Translated(ib, language.English))
	})

	t.Run("translates main message and field violations", func(t *testing.T) {
		fv := NewFieldViolation("email", "REQUIRED", "required")
		s := New(http.StatusUnprocessableEntity, "INVALID_FORMAT", "bad format").
			WithFieldViolations(fv)

		translated := s.Translated(ib, language.Chinese)

		assert.Equal(t, "bad format", translated.Message()) // original preserved
		require.NotNil(t, translated.GetLocalizedMessage())
		assert.Equal(t, "格式无效", translated.GetLocalizedMessage().Message)
		fvs := translated.FieldViolations()
		require.Len(t, fvs, 1)
		assert.Equal(t, "required", fvs[0].Description()) // original preserved
		require.NotNil(t, fvs[0].GetLocalizedMessage())
		assert.Equal(t, "必填", fvs[0].GetLocalizedMessage().Message)
	})

	t.Run("idempotent translation", func(t *testing.T) {
		s := New(http.StatusNotFound, "NOT_FOUND", "not found")
		translated1 := s.Translated(ib, language.English)
		translated2 := translated1.Translated(ib, language.English)

		// Second translation should not change anything (localizedMessage already set)
		assert.Equal(t, translated1.Message(), translated2.Message())
		assert.Equal(t, translated1.GetLocalizedMessage().Message, translated2.GetLocalizedMessage().Message)
	})
}

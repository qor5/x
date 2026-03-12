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
TOO_SHORT_MIN,Must be at least %d characters,至少需要%d个字符
OUT_OF_RANGE,Must be between %d and %d,必须在%d到%d之间
custom.welcome,Welcome {{.Name}} to {{.App}},欢迎 {{.Name}} 使用 {{.App}}
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
		err := New(http.StatusBadRequest, ReasonInvalidArgument, "invalid").
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
		err := New(http.StatusBadRequest, ReasonInvalidArgument, "invalid").
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

		err := New(http.StatusBadRequest, ReasonInvalidArgument, "invalid").
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
		err := New(http.StatusBadRequest, ReasonInvalidArgument, "invalid").
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

func TestTranslateError_LocalizedWithI18N(t *testing.T) {
	ib := newTestI18N(t)

	t.Run("status reason as i18n key - English", func(t *testing.T) {
		// reason "NOT_FOUND" is automatically used as i18n key
		err := Error(http.StatusNotFound, "NOT_FOUND", "user xyz not found")
		result := TranslateError(err, ib, language.English)

		st := Convert(result)
		assert.Equal(t, http.StatusNotFound, st.StatusCode())
		assert.Equal(t, "NOT_FOUND", st.Reason())
		assert.Equal(t, "user xyz not found", st.Message()) // original preserved
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "Not Found", st.GetLocalizedMessage().Message)
		assert.Equal(t, "en", st.GetLocalizedMessage().Locale)
	})

	t.Run("status reason as i18n key - Chinese", func(t *testing.T) {
		err := Error(http.StatusInternalServerError, "INTERNAL", "database connection failed")
		result := TranslateError(err, ib, language.Chinese)

		st := Convert(result)
		assert.Equal(t, "database connection failed", st.Message()) // original preserved
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "内部错误", st.GetLocalizedMessage().Message)
		assert.Equal(t, "zh", st.GetLocalizedMessage().Locale)
	})

	t.Run("field violation reason as i18n key", func(t *testing.T) {
		fvs := []*FieldViolation{
			NewFieldViolation("email", "REQUIRED", "email is required"),
			NewFieldViolation("password", "TOO_SHORT", "password is too short"),
		}
		err := BadRequest(fvs).Err()
		result := TranslateError(err, ib, language.Chinese)

		st := Convert(result)
		violations := st.FieldViolations()
		require.Len(t, violations, 2)

		assert.Equal(t, "email is required", violations[0].Description())
		require.NotNil(t, violations[0].GetLocalizedMessage())
		assert.Equal(t, "必填", violations[0].GetLocalizedMessage().Message)

		assert.Equal(t, "password is too short", violations[1].Description())
		require.NotNil(t, violations[1].GetLocalizedMessage())
		assert.Equal(t, "太短", violations[1].GetLocalizedMessage().Message)
	})

	t.Run("custom localized key overrides reason", func(t *testing.T) {
		// reason is "BAD_REQUEST" but i18n key is overridden to "NOT_FOUND"
		err := New(http.StatusBadRequest, "BAD_REQUEST", "something went wrong").
			WithLocalized("NOT_FOUND").Err()
		result := TranslateError(err, ib, language.Chinese)

		st := Convert(result)
		assert.Equal(t, "BAD_REQUEST", st.Reason())              // reason unchanged
		assert.Equal(t, "something went wrong", st.Message())    // original preserved
		assert.Equal(t, "未找到", st.GetLocalizedMessage().Message) // translated by custom key
	})
}

func TestTranslateError_TemplateVariables(t *testing.T) {
	ib := newTestI18N(t)

	t.Run("positional args - single", func(t *testing.T) {
		// WithLocalizedArgs sets positional %d args
		fv := NewFieldViolation("password", "TOO_SHORT_MIN", "must be at least 8 characters").
			WithLocalizedArgs(8)
		err := BadRequest(fv).Err()
		result := TranslateError(err, ib, language.Chinese)

		st := Convert(result)
		violations := st.FieldViolations()
		require.Len(t, violations, 1)
		assert.Equal(t, "must be at least 8 characters", violations[0].Description())
		require.NotNil(t, violations[0].GetLocalizedMessage())
		assert.Equal(t, "至少需要8个字符", violations[0].GetLocalizedMessage().Message)
	})

	t.Run("positional args - multiple", func(t *testing.T) {
		fv := NewFieldViolation("age", "OUT_OF_RANGE", "must be between 18 and 120").
			WithLocalizedArgs(18, 120)
		err := BadRequest(fv).Err()
		result := TranslateError(err, ib, language.Chinese)

		st := Convert(result)
		violations := st.FieldViolations()
		require.Len(t, violations, 1)
		assert.Equal(t, "must be between 18 and 120", violations[0].Description())
		require.NotNil(t, violations[0].GetLocalizedMessage())
		assert.Equal(t, "必须在18到120之间", violations[0].GetLocalizedMessage().Message)
	})

	t.Run("positional args - English", func(t *testing.T) {
		fv := NewFieldViolation("password", "TOO_SHORT_MIN", "must be at least 8 characters").
			WithLocalizedArgs(8)
		err := BadRequest(fv).Err()
		result := TranslateError(err, ib, language.English)

		st := Convert(result)
		violations := st.FieldViolations()
		require.Len(t, violations, 1)
		require.NotNil(t, violations[0].GetLocalizedMessage())
		assert.Equal(t, "Must be at least 8 characters", violations[0].GetLocalizedMessage().Message)
	})

	t.Run("named args with map", func(t *testing.T) {
		// WithLocalized with a map triggers Go template rendering {{.Name}}
		err := New(http.StatusBadRequest, "BAD_REQUEST", "invalid request").
			WithLocalized("custom.welcome", map[string]any{
				"Name": "Alice",
				"App":  "PIM",
			}).Err()
		result := TranslateError(err, ib, language.Chinese)

		st := Convert(result)
		assert.Equal(t, "invalid request", st.Message())
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "欢迎 Alice 使用 PIM", st.GetLocalizedMessage().Message)
	})

	t.Run("named args with map - English", func(t *testing.T) {
		err := New(http.StatusBadRequest, "BAD_REQUEST", "invalid request").
			WithLocalized("custom.welcome", map[string]any{
				"Name": "Bob",
				"App":  "Console",
			}).Err()
		result := TranslateError(err, ib, language.English)

		st := Convert(result)
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "Welcome Bob to Console", st.GetLocalizedMessage().Message)
	})

	t.Run("positional args on status level", func(t *testing.T) {
		// Status-level translation with positional args
		err := New(http.StatusBadRequest, "BAD_REQUEST", "greeting failed").
			WithLocalized("custom.greeting", "World").Err()
		result := TranslateError(err, ib, language.Chinese)

		st := Convert(result)
		assert.Equal(t, "greeting failed", st.Message())
		require.NotNil(t, st.GetLocalizedMessage())
		assert.Equal(t, "你好 World", st.GetLocalizedMessage().Message)
	})

	t.Run("mixed field violations with and without args", func(t *testing.T) {
		fvs := []*FieldViolation{
			NewFieldViolation("email", "REQUIRED", "email is required"),
			NewFieldViolation("password", "TOO_SHORT_MIN", "must be at least 8 characters").
				WithLocalizedArgs(8),
			NewFieldViolation("age", "OUT_OF_RANGE", "must be between 18 and 120").
				WithLocalizedArgs(18, 120),
		}
		err := BadRequest(fvs).Err()
		result := TranslateError(err, ib, language.Chinese)

		st := Convert(result)
		violations := st.FieldViolations()
		require.Len(t, violations, 3)

		// No args - simple key lookup
		require.NotNil(t, violations[0].GetLocalizedMessage())
		assert.Equal(t, "必填", violations[0].GetLocalizedMessage().Message)

		// Single positional arg
		require.NotNil(t, violations[1].GetLocalizedMessage())
		assert.Equal(t, "至少需要8个字符", violations[1].GetLocalizedMessage().Message)

		// Multiple positional args
		require.NotNil(t, violations[2].GetLocalizedMessage())
		assert.Equal(t, "必须在18到120之间", violations[2].GetLocalizedMessage().Message)
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
		s := New(http.StatusBadRequest, "INVALID_FORMAT", "bad format").
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

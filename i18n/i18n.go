package i18n

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/text/language"
)

type ModuleKey string

type Builder struct {
	supportLanguages                   []language.Tag
	getSupportLanguagesFromRequestFunc func(R *http.Request) []language.Tag
	moduleMessages                     map[language.Tag]context.Context
	matcher                            language.Matcher
	cookieName                         string
	queryName                          string
}

type Messages interface{}

func New() *Builder {
	b := &Builder{
		supportLanguages: []language.Tag{
			language.English,
		},
		moduleMessages: map[language.Tag]context.Context{language.English: context.TODO()},
		cookieName:     "lang",
		queryName:      "lang",
	}
	b.matcher = language.NewMatcher(b.supportLanguages)
	return b
}

func (b *Builder) defaultLanguage() language.Tag {
	return b.supportLanguages[0]
}

func (b *Builder) GetCookieName() string {
	return b.cookieName
}

func (b *Builder) GetQueryName() string {
	return b.queryName
}

func (b *Builder) SupportLanguages(vs ...language.Tag) (r *Builder) {
	if len(vs) == 0 {
		panic("have to support at least one language")
	}
	b.supportLanguages = vs
	for _, l := range b.supportLanguages {
		if b.moduleMessages[l] == nil {
			b.moduleMessages[l] = context.TODO()
		}
	}
	b.matcher = language.NewMatcher(b.supportLanguages)
	return b
}

func (b *Builder) GetSupportLanguages() []language.Tag {
	return b.supportLanguages
}

func (b *Builder) GetSupportLanguagesFromRequest(R *http.Request) []language.Tag {
	if b.getSupportLanguagesFromRequestFunc != nil {
		return b.getSupportLanguagesFromRequestFunc(R)
	}
	return b.GetSupportLanguages()
}

func (b *Builder) GetSupportLanguagesFromRequestFunc(v func(R *http.Request) []language.Tag) (r *Builder) {
	b.getSupportLanguagesFromRequestFunc = v
	return b
}

func (b *Builder) RegisterForModule(lang language.Tag, module ModuleKey, msg Messages) (r *Builder) {
	c := b.moduleMessages[lang]
	if c == nil {
		c = context.TODO()
	}

	c = context.WithValue(c, module, msg)
	b.moduleMessages[lang] = c
	return b
}

func MustGetModuleMessages(r *http.Request, module ModuleKey, defaultMessages Messages) Messages {
	v := r.Context().Value(moduleMessagesKey)
	if v == nil {
		return defaultMessages
	}

	msg := v.(context.Context).Value(module)
	if msg == nil {
		msg = defaultMessages
	}
	return msg
}

type i18nContextKey int

const (
	moduleMessagesKey i18nContextKey = iota
	dynaBuilderKey
	languageTagKey
)

func LanguageTagFromContext(ctx context.Context, fallback language.Tag) language.Tag {
	v, ok := ctx.Value(languageTagKey).(language.Tag)
	if !ok {
		return fallback
	}
	return v
}

func (b *Builder) EnsureLanguage(in http.Handler) (out http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := ""
		lang = r.FormValue(b.queryName)
		if len(lang) > 0 {
			maxAge := 365 * 24 * 60 * 60
			http.SetCookie(w, &http.Cookie{
				Name:    b.cookieName,
				Value:   lang,
				Path:    "/",
				MaxAge:  maxAge,
				Expires: time.Now().Add(time.Duration(maxAge) * time.Second),
			})
		} else {
			lang = b.GetCurrentLangFromCookie(r)
		}

		accept := r.Header.Get("Accept-Language")

		var availableLanguages []language.Tag
		var matcher language.Matcher
		if len(lang) > 0 {
			availableLanguages = b.GetSupportLanguages()
			matcher = b.matcher
		} else {
			availableLanguages = b.GetSupportLanguagesFromRequest(r)
			matcher = language.NewMatcher(availableLanguages)
		}
		_, i := language.MatchStrings(matcher, lang, accept)
		tag := availableLanguages[i]

		moduleMsgs := b.moduleMessages[tag]
		if moduleMsgs == nil {
			moduleMsgs = b.moduleMessages[b.defaultLanguage()]
		}
		if moduleMsgs == nil {
			panic(fmt.Sprintf("language %s not supported", tag.String()))
		}
		dyna := DynaNew().Language(tag.String())
		ctx := context.WithValue(r.Context(), moduleMessagesKey, moduleMsgs)
		ctx = context.WithValue(ctx, dynaBuilderKey, dyna)
		ctx = context.WithValue(ctx, languageTagKey, tag)
		in.ServeHTTP(w, r.WithContext(ctx))
		if dyna.HaveMissingKeys() {
			log.Println(dyna.PrettyMissingKeys())
		}
	})
}

func (b *Builder) GetCurrentLangFromCookie(r *http.Request) (lang string) {
	langCookie, _ := r.Cookie(b.cookieName)
	if langCookie != nil {
		lang = langCookie.Value
	}
	return
}

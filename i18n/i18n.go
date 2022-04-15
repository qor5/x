package i18n

import (
	"context"
	"fmt"
	"log"
	"net/http"

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

type Messages interface {
}

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
)

func (b *Builder) EnsureLanguage(in http.Handler) (out http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var lang = ""
		lang = r.FormValue(b.queryName)
		if len(lang) > 0 {
			http.SetCookie(w, &http.Cookie{Name: b.cookieName, Value: lang})
		} else {
			langCookie, _ := r.Cookie(b.cookieName)
			if langCookie != nil {
				lang = langCookie.Value
			}
		}

		var tag language.Tag
		if len(lang) > 0 {
			accept := r.Header.Get("Accept-Language")
			tag, _ = language.MatchStrings(b.matcher, lang, accept)
		} else {
			var matcher = language.NewMatcher(b.GetSupportLanguagesFromRequest(r))
			tag, _ = language.MatchStrings(matcher, lang)
		}

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
		in.ServeHTTP(w, r.WithContext(ctx))
		if dyna.HaveMissingKeys() {
			log.Println(dyna.PrettyMissingKeys())
		}
	})
}

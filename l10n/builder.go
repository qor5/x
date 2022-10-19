package l10n

import (
	"context"
	"net/http"
	"time"

	"github.com/biter777/countries"
)

type Builder struct {
	supportLocales                   []countries.CountryCode
	getSupportLocalesFromRequestFunc func(R *http.Request) []countries.CountryCode
	cookieName                       string
	queryName                        string
}

func New() *Builder {
	b := &Builder{
		supportLocales: []countries.CountryCode{},
		cookieName:     "locale",
		queryName:      "locale",
	}
	return b
}

func (b *Builder) IsTurnedOn() bool {
	return len(b.GetSupportLocales()) > 0
}

func (b *Builder) GetCookieName() string {
	return b.cookieName
}

func (b *Builder) GetQueryName() string {
	return b.queryName
}

func (b *Builder) SupportLocales(vs ...countries.CountryCode) (r *Builder) {
	if len(vs) == 0 {
		panic("have to support at least one language")
	}
	b.supportLocales = vs
	return b
}

func (b *Builder) GetSupportLocales() []countries.CountryCode {
	return b.supportLocales
}

func (b *Builder) GetSupportLocalesFromRequest(R *http.Request) []countries.CountryCode {
	if b.getSupportLocalesFromRequestFunc != nil {
		return b.getSupportLocalesFromRequestFunc(R)
	}
	return b.GetSupportLocales()
}

func (b *Builder) GetSupportLocalesFromRequestFunc(v func(R *http.Request) []countries.CountryCode) (r *Builder) {
	b.getSupportLocalesFromRequestFunc = v
	return b
}

func (b *Builder) GetCurrentLocaleFromCookie(r *http.Request) (locale string) {
	localeCookie, _ := r.Cookie(b.cookieName)
	if localeCookie != nil {
		locale = localeCookie.Value
	}
	return
}

func (b *Builder) GetCorrectLocale(r *http.Request) string {
	locale := r.FormValue(b.queryName)
	if locale == "" {
		locale = b.GetCurrentLocaleFromCookie(r)
	}

	supportLocales := b.GetSupportLocalesFromRequest(r)
	for _, v := range supportLocales {
		if locale == v.String() {
			return locale
		}
	}

	return supportLocales[0].String()
}

type l10nContextKey int

const (
	HasLocaleKey l10nContextKey = iota
	LocaleCode
)

func (b *Builder) EnsureLocale(in http.Handler) (out http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(b.GetSupportLocalesFromRequest(r)) == 0 {
			in.ServeHTTP(w, r)
			return
		}

		var locale = b.GetCorrectLocale(r)
		if len(locale) > 0 {
			maxAge := 365 * 24 * 60 * 60
			http.SetCookie(w, &http.Cookie{
				Name:    b.cookieName,
				Value:   locale,
				Path:    "/",
				MaxAge:  maxAge,
				Expires: time.Now().Add(time.Duration(maxAge) * time.Second),
			})
		}

		ctx := context.WithValue(r.Context(), LocaleCode, locale)

		in.ServeHTTP(w, r.WithContext(ctx))
	})
}

package page

import (
	"net/http"
	"net/url"
	"strings"
)

const PathBeforeStripPrefixKey = "X-Path-Before-StripPrefix"

func StripPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pathBeforeStripPrefix := r.URL.Path
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.Header.Set(PathBeforeStripPrefixKey, pathBeforeStripPrefix)
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}

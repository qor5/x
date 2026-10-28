package login

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewarePublicPaths(t *testing.T) {
	b := New().Secret("test-secret").PublicPathPrefixes("/static/", "/favicon.ico")
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	h := b.Middleware()(next)

	for _, c := range []struct {
		name   string
		path   string
		expect int
	}{
		{name: "static extension outside public paths", path: "/orders/1.js", expect: http.StatusFound},
		{name: "uppercase static extension outside public paths", path: "/orders/1.CSS", expect: http.StatusFound},
		{name: "no extension", path: "/orders/1", expect: http.StatusFound},
		{name: "login assets", path: "/auth/assets/style.css", expect: http.StatusOK},
		{name: "declared prefix", path: "/static/logo.png", expect: http.StatusOK},
		{name: "declared exact path", path: "/favicon.ico", expect: http.StatusOK},
		{name: "declared prefix but not static", path: "/static/orders", expect: http.StatusFound},
		{name: "unclean path", path: "/static/../orders/1.js", expect: http.StatusFound},
		{name: "prefix without trailing slash boundary", path: "/staticx/a.js", expect: http.StatusFound},
	} {
		t.Run(c.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, c.path, nil)
			w := httptest.NewRecorder()
			h.ServeHTTP(w, r)
			if w.Code != c.expect {
				t.Errorf("GET %s: expect %d, got %d", c.path, c.expect, w.Code)
			}
		})
	}
}

package login

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSanitizeContinueURL(t *testing.T) {
	const host = "admin.example.com"
	for _, c := range []struct {
		name      string
		raw       string
		allowHost string
		expect    string
	}{
		{name: "relative path with query", raw: "/admin/x?a=1", expect: "/admin/x?a=1"},
		{name: "root", raw: "/", expect: "/"},
		{name: "relative with fragment", raw: "/admin/x?a=1#top", expect: "/admin/x?a=1#top"},
		{name: "escaped path kept", raw: "/admin/a%20b?q=%2F", expect: "/admin/a%20b?q=%2F"},
		{name: "empty", raw: "", expect: ""},
		{name: "protocol relative", raw: "//evil.com", expect: ""},
		{name: "protocol relative with path", raw: "//evil.com/a", expect: ""},
		{name: "backslash after slash", raw: `/\evil.com`, expect: ""},
		{name: "backslash later in path", raw: `/a\..\\evil.com`, expect: ""},
		{name: "absolute https", raw: "https://evil.com/a", expect: ""},
		{name: "absolute http other host even if allowHost set", raw: "http://evil.com/a", allowHost: host, expect: ""},
		{name: "absolute same host without allowHost", raw: "http://" + host + "/a", expect: ""},
		{name: "referer same host", raw: "http://" + host + "/a", allowHost: host, expect: "/a"},
		{name: "referer same host case insensitive with query", raw: "https://ADMIN.example.com/a/b?x=1#f", allowHost: host, expect: "/a/b?x=1#f"},
		{name: "referer same host empty path", raw: "https://" + host, allowHost: host, expect: "/"},
		{name: "referer same host with userinfo", raw: "https://u:p@" + host + "/a", allowHost: host, expect: ""},
		{name: "referer same host double slash path", raw: "https://" + host + "//evil.com", allowHost: host, expect: ""},
		{name: "referer other host", raw: "https://evil.com/admin", allowHost: host, expect: ""},
		{name: "referer host with port mismatch", raw: "https://" + host + ":8443/a", allowHost: host, expect: ""},
		{name: "javascript scheme", raw: "javascript:alert(1)", expect: ""},
		{name: "javascript scheme with allowHost", raw: "javascript:alert(1)", allowHost: host, expect: ""},
		{name: "data scheme", raw: "data:text/html,hi", allowHost: host, expect: ""},
		{name: "relative without leading slash", raw: "admin/x", expect: ""},
		{name: "tab inside", raw: "/\t/evil.com", expect: ""},
		{name: "newline inside", raw: "/a\r\nSet-Cookie: x=1", expect: ""},
		{name: "DEL char", raw: "/a\x7f", expect: ""},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := sanitizeContinueURL(c.raw, c.allowHost)
			if got != c.expect {
				t.Errorf("sanitizeContinueURL(%q, %q) = %q, want %q", c.raw, c.allowHost, got, c.expect)
			}
		})
	}
}

func continueCookie(w *httptest.ResponseRecorder, name string) *http.Cookie {
	for _, c := range w.Result().Cookies() {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func TestSetContinueURL(t *testing.T) {
	b := New().Secret("test-secret")
	for _, c := range []struct {
		name    string
		target  string
		referer string
		expect  string // "" means cookie must not be set
	}{
		{name: "plain request uri", target: "/admin/orders?page=2", expect: "/admin/orders?page=2"},
		{name: "execute event uses same-host referer", target: "/admin/orders?__execute_event__=x", referer: "http://example.com/admin/orders?page=3", expect: "/admin/orders?page=3"},
		{name: "execute event with foreign referer", target: "/admin/orders?__execute_event__=x", referer: "https://evil.com/phish", expect: ""},
		{name: "execute event with protocol relative referer", target: "/admin/orders?__execute_event__=x", referer: "//evil.com/phish", expect: ""},
		{name: "login page ignored", target: b.loginPageURL, expect: ""},
	} {
		t.Run(c.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, c.target, nil) // Host: example.com
			if c.referer != "" {
				r.Header.Set("Referer", c.referer)
			}
			w := httptest.NewRecorder()
			b.setContinueURL(w, r)
			ck := continueCookie(w, b.continueUrlCookieName)
			if c.expect == "" {
				if ck != nil {
					t.Fatalf("expected no continue cookie, got %q", ck.Value)
				}
				return
			}
			if ck == nil || ck.Value != c.expect {
				t.Fatalf("expected continue cookie %q, got %+v", c.expect, ck)
			}
		})
	}
}

func TestGetContinueURL(t *testing.T) {
	b := New().Secret("test-secret")
	for _, c := range []struct {
		name   string
		cookie string
		expect string
	}{
		{name: "safe value", cookie: "/admin/x?a=1", expect: "/admin/x?a=1"},
		{name: "legacy same-host absolute value", cookie: "http://example.com/admin/x", expect: "/admin/x"},
		{name: "foreign absolute", cookie: "https://evil.com/a", expect: ""},
		{name: "protocol relative", cookie: "//evil.com", expect: ""},
	} {
		t.Run(c.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/auth/login", nil)
			r.AddCookie(&http.Cookie{Name: b.continueUrlCookieName, Value: c.cookie})
			w := httptest.NewRecorder()
			got := b.getContinueURL(w, r)
			if got != c.expect {
				t.Fatalf("getContinueURL = %q, want %q", got, c.expect)
			}
			ck := continueCookie(w, b.continueUrlCookieName)
			if ck == nil || ck.MaxAge >= 0 || ck.Value != "" {
				t.Fatalf("expected continue cookie to be cleared, got %+v", ck)
			}
		})
	}
}

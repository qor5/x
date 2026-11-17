package login

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSanitizeContinueURL(t *testing.T) {
	const host = "admin.example.com"
	for _, c := range []struct {
		name          string
		raw           string
		allowAbsolute bool
		expect        string
	}{
		{name: "relative path with query", raw: "/admin/x?a=1", expect: "/admin/x?a=1"},
		{name: "root", raw: "/", expect: "/"},
		{name: "relative with fragment", raw: "/admin/x?a=1#top", expect: "/admin/x?a=1#top"},
		{name: "escaped path kept", raw: "/admin/a%20b?q=%2F", expect: "/admin/a%20b?q=%2F"},
		{name: "empty", raw: "", expect: ""},
		{name: "protocol relative", raw: "//evil.com", expect: ""},
		{name: "protocol relative with path", raw: "//evil.com/a", expect: ""},
		{name: "protocol relative even if absolute allowed", raw: "//evil.com/a", allowAbsolute: true, expect: ""},
		{name: "backslash after slash", raw: `/\evil.com`, expect: ""},
		{name: "backslash later in path", raw: `/a\..\\evil.com`, expect: ""},
		{name: "absolute https rejected for relative-only source", raw: "https://evil.com/a", expect: ""},
		{name: "absolute same host rejected for relative-only source", raw: "http://" + host + "/a", expect: ""},
		{name: "referer reduced to path", raw: "http://" + host + "/a", allowAbsolute: true, expect: "/a"},
		{name: "referer other host reduced to same-origin path", raw: "https://evil.com/admin?x=1", allowAbsolute: true, expect: "/admin?x=1"},
		{name: "referer with query and fragment", raw: "https://ADMIN.example.com/a/b?x=1#f", allowAbsolute: true, expect: "/a/b?x=1#f"},
		{name: "referer empty path", raw: "https://" + host, allowAbsolute: true, expect: "/"},
		{name: "referer with port", raw: "https://" + host + ":8443/a", allowAbsolute: true, expect: "/a"},
		{name: "referer with userinfo", raw: "https://u:p@" + host + "/a", allowAbsolute: true, expect: ""},
		{name: "referer double slash path", raw: "https://" + host + "//evil.com", allowAbsolute: true, expect: ""},
		{name: "referer backslash path", raw: "https://" + host + `/\evil.com`, allowAbsolute: true, expect: ""},
		{name: "scheme without host", raw: "http:/a", allowAbsolute: true, expect: ""},
		{name: "ftp scheme", raw: "ftp://" + host + "/a", allowAbsolute: true, expect: ""},
		{name: "javascript scheme", raw: "javascript:alert(1)", expect: ""},
		{name: "javascript scheme with absolute allowed", raw: "javascript:alert(1)", allowAbsolute: true, expect: ""},
		{name: "data scheme", raw: "data:text/html,hi", allowAbsolute: true, expect: ""},
		{name: "relative without leading slash", raw: "admin/x", expect: ""},
		{name: "tab inside", raw: "/\t/evil.com", expect: ""},
		{name: "newline inside", raw: "/a\r\nSet-Cookie: x=1", expect: ""},
		{name: "DEL char", raw: "/a\x7f", expect: ""},
	} {
		t.Run(c.name, func(t *testing.T) {
			got := sanitizeContinueURL(c.raw, c.allowAbsolute)
			if got != c.expect {
				t.Errorf("sanitizeContinueURL(%q, %v) = %q, want %q", c.raw, c.allowAbsolute, got, c.expect)
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
		{name: "execute event with referer from another host is reduced to a path", target: "/admin/orders?__execute_event__=x", referer: "https://proxy.internal/admin/orders?page=4", expect: "/admin/orders?page=4"},
		{name: "execute event with javascript referer", target: "/admin/orders?__execute_event__=x", referer: "javascript:alert(1)", expect: ""},
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
		{name: "legacy foreign absolute is reduced to a path", cookie: "https://evil.com/a", expect: "/a"},
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

package login

import (
	"net/url"
	"strings"
)

// sanitizeContinueURL normalizes raw into a safe, same-site relative URL
// (path + query + fragment) suitable for redirecting after login.
// It returns "" if raw cannot be made safe.
//
// A relative reference must start with a single "/". Protocol-relative
// references ("//host"), backslashes in the path (which browsers treat as
// "/"), control characters, userinfo and non-http(s) schemes are rejected.
//
// If allowAbsolute is true, an absolute http(s) URL (any host, e.g. a Referer)
// is accepted and reduced to its path + query + fragment. The host is
// deliberately not compared with the request host: the result is always a
// relative path on the current origin, and comparing would break behind
// proxies that rewrite Host. Absolute URLs are rejected otherwise.
func sanitizeContinueURL(raw string, allowAbsolute bool) string {
	if raw == "" {
		return ""
	}
	for i := 0; i < len(raw); i++ {
		if c := raw[i]; c < 0x20 || c == 0x7f {
			return ""
		}
	}
	pathEnd := strings.IndexAny(raw, "?#")
	if pathEnd < 0 {
		pathEnd = len(raw)
	}
	if strings.Contains(raw[:pathEnd], `\`) {
		return ""
	}

	u, err := url.Parse(raw)
	if err != nil || u.Opaque != "" || u.User != nil {
		return ""
	}
	if u.Scheme != "" || u.Host != "" {
		if !allowAbsolute || u.Host == "" {
			return ""
		}
		if s := strings.ToLower(u.Scheme); s != "http" && s != "https" {
			return ""
		}
	} else if !strings.HasPrefix(raw, "/") {
		return ""
	}

	p := u.EscapedPath()
	if p == "" {
		p = "/"
	}
	if !strings.HasPrefix(p, "/") || strings.HasPrefix(p, "//") || strings.HasPrefix(u.Path, "//") {
		return ""
	}
	out := p
	if u.RawQuery != "" {
		out += "?" + u.RawQuery
	}
	if u.Fragment != "" {
		out += "#" + u.EscapedFragment()
	}
	return out
}

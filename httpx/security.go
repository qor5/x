package httpx

import (
	"fmt"
	"mime"
	"net/http"
	"slices"
	"strings"

	connectcors "connectrpc.com/cors"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/samber/lo"
)

var (
	HeaderContentType = http.CanonicalHeaderKey("Content-Type")
	HeaderRequestedBy = http.CanonicalHeaderKey("X-Requested-By")
)

// SimpleRequestContentTypes are the only Content-Type values that qualify as CORS simple requests.
// Per the Fetch Standard (https://fetch.spec.whatwg.org/#cors-safelisted-request-header),
// only these three Content-Types are allowed in simple requests without triggering a preflight.
var SimpleRequestContentTypes = []string{
	"application/x-www-form-urlencoded",
	"multipart/form-data",
	"text/plain",
}

// DenySimpleRequests is the default middleware instance that denies CORS simple requests
// without the X-Requested-By header. Use DenySimpleRequestsFactory for custom skip logic.
var DenySimpleRequests = DenySimpleRequestsFactory(nil)

// DenySimpleRequestsFactory creates a configurable middleware that prevents CORS simple requests.
// This provides CSRF protection by requiring either:
//   - The X-Requested-By header to be present, OR
//   - A Content-Type that is NOT a simple request type (which triggers CORS preflight)
//
// Per the Fetch Standard, CORS simple requests are limited to:
//   - Methods: GET, HEAD, POST
//   - Content-Types: application/x-www-form-urlencoded, multipart/form-data, text/plain
//
// Requests with other Content-Types (e.g., application/json) automatically trigger a preflight
// OPTIONS request, which can be validated by CORS policies, so they are allowed through.
//
// The skipCheck function allows selective exemption of certain requests from this check.
// When skipCheck returns true, the request bypasses all validation.
// If skipCheck is nil, all requests will be checked.
var DenySimpleRequestsFactory = func(skipCheck func(r *http.Request) bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if skipCheck != nil && skipCheck(r) {
				next.ServeHTTP(w, r)
				return
			}

			if r.Header.Get(HeaderRequestedBy) != "" {
				next.ServeHTTP(w, r)
				return
			}

			// Check if this could be a CORS simple request (GET/HEAD/POST with simple Content-Type)
			switch r.Method {
			case http.MethodGet, http.MethodHead, http.MethodPost:
				mediaType, _, _ := ParseContentType(r)
				if mediaType == "" || slices.Contains(SimpleRequestContentTypes, mediaType) {
					// Simple request without X-Requested-By header - deny it
					http.Error(w, fmt.Sprintf("%s header is required", HeaderRequestedBy), http.StatusBadRequest)
					return
				}
			}

			// Non-simple request (other methods or non-simple Content-Type) - allow through
			// These will trigger CORS preflight which provides the necessary protection
			next.ServeHTTP(w, r)
		})
	}
}

var Security = func(conf SecurityConfig) func(next http.Handler) http.Handler {
	corsOpts := cors.Options{
		AllowedOrigins:   conf.CORS.AllowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   lo.Uniq(slices.Concat([]string{http.MethodPost}, conf.CORS.AllowedMethods /*, connectcors.AllowedMethods()*/)),
		AllowedHeaders:   lo.Uniq(slices.Concat([]string{HeaderContentType, HeaderRequestedBy}, conf.CORS.AllowedHeaders, connectcors.AllowedHeaders())),
		ExposedHeaders:   lo.Uniq(slices.Concat(conf.CORS.ExposedHeaders, connectcors.ExposedHeaders())),
		MaxAge:           int(conf.CORS.MaxAge.Seconds()),
		Debug:            conf.CORS.Debug,
	}
	if len(corsOpts.AllowedOrigins) == 0 {
		corsOpts.AllowOriginFunc = func(_ string) bool {
			return false // default deny all cross origins
		}
	}
	c := cors.New(corsOpts)
	frameAncestors := buildFrameAncestors(conf.CORS.AllowedOrigins)
	denySimpleRequests := DenySimpleRequestsFactory(conf.CORS.SkipDenySimpleRequests)
	return func(next http.Handler) http.Handler {
		var handler http.Handler
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if conf.DenyMIMETypeSniffing {
				w.Header().Set("X-Content-Type-Options", "nosniff") // prevent MIME type sniffing
			}

			if conf.DenyClickjacking {
				w.Header().Set("Content-Security-Policy", frameAncestors) // prevent clickjacking
				w.Header().Set("X-Frame-Options", "SAMEORIGIN")           // prevent clickjacking
			}

			if conf.EnableHSTS {
				w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload") // HSTS
			}

			next.ServeHTTP(w, r)
		})
		if conf.CORS.DenySimpleRequests {
			handler = denySimpleRequests(handler)
		}
		return c.Handler(handler)
	}
}

func buildFrameAncestors(origins []string) string {
	if len(origins) == 0 {
		return "frame-ancestors 'self'"
	}

	ancestors := []string{"'self'"}
	ancestors = append(ancestors, origins...)
	return "frame-ancestors " + strings.Join(ancestors, " ") + ";"
}

func ParseContentType(r *http.Request) (mediaType string, params map[string]string, err error) {
	contentTypeVals := r.Header.Values(HeaderContentType)
	if len(contentTypeVals) == 0 {
		return "", nil, errors.Errorf("%s header not found", HeaderContentType)
	}
	if len(contentTypeVals) > 1 {
		return "", nil, errors.Errorf("multiple %s headers found", HeaderContentType)
	}
	mediaType, params, err = mime.ParseMediaType(contentTypeVals[0])
	if err != nil {
		return "", nil, errors.Wrapf(err, "failed to parse %s %s", HeaderContentType, contentTypeVals[0])
	}
	return mediaType, params, nil
}

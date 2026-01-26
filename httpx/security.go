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

// DenySimpleRequests is the default middleware instance that denies all simple requests
// without the X-Requested-By header. Use DenySimpleRequestsFactory for custom skip logic.
var DenySimpleRequests = DenySimpleRequestsFactory(nil)

// DenySimpleRequestsFactory creates a configurable middleware that prevents CORS simple requests.
// The skipCheck function allows selective exemption of certain requests from the header requirement.
// When skipCheck returns true, the request is allowed through without checking the X-Requested-By header.
// If skipCheck is nil, all requests will be checked.
// This provides CSRF protection when used with CORS, as browsers will not automatically include
// custom headers in simple requests, forcing a preflight OPTIONS request.
var DenySimpleRequestsFactory = func(skipCheck func(r *http.Request) bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if skipCheck != nil && skipCheck(r) {
				next.ServeHTTP(w, r)
				return
			}
			if r.Header.Get(HeaderRequestedBy) == "" {
				http.Error(w, fmt.Sprintf("%s header is required", HeaderRequestedBy), http.StatusBadRequest)
				return
			}
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

package httpx

import (
	"mime"
	"net/http"
	"slices"
	"strings"

	connectcors "connectrpc.com/cors"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"github.com/samber/lo"
)

var HeaderContentType = http.CanonicalHeaderKey("Content-Type")

var DenySimpleRequests = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Deny CORS simple Request for GET/HEAD/POST
		switch r.Method {
		case http.MethodGet, http.MethodHead, http.MethodPost:
			// Only allow application/json or application/proto content type for simple requests
			// In this way, all simple requests will be pre-checked by cors before they can be executed.
			mediaType, _, err := ParseContentType(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			switch mediaType {
			case "application/json", "application/proto":
			default:
				http.Error(w, "Content-Type must be application/json or application/proto", http.StatusUnsupportedMediaType)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

var Security = func(conf SecurityConfig) func(next http.Handler) http.Handler {
	corsOpts := cors.Options{
		AllowedOrigins:   conf.CORS.AllowedOrigins,
		AllowCredentials: true,
		AllowedMethods:   lo.Uniq(slices.Concat([]string{http.MethodPost}, conf.CORS.AllowedMethods /*, connectcors.AllowedMethods()*/)),
		AllowedHeaders:   lo.Uniq(slices.Concat([]string{HeaderContentType}, conf.CORS.AllowedHeaders, connectcors.AllowedHeaders())),
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
			handler = DenySimpleRequests(handler)
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

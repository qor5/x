package httpx

import "net/http"

// NoStore is a middleware that sets HTTP headers to prevent caching of responses.
// It applies the appropriate cache control headers based on the HTTP protocol version:
// - For HTTP/1.1+: Sets Cache-Control: no-store
// - For HTTP/1.0 and below: Also sets Pragma: no-cache and Expires: 0
var NoStore = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set Cache-Control for HTTP/1.1+
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")

		// Handle HTTP/1.0 and below
		if !r.ProtoAtLeast(1, 1) {
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}

		next.ServeHTTP(w, r)
	})
}

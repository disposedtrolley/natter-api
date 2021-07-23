package middleware

import "net/http"

func SetJSONResponseHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func SetSecurityResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Content-Type-Options", "nosniff")
		w.Header().Add("X-Frame-Options", "DENY")
		w.Header().Add("X-XSS-Protection", "0")
		w.Header().Add("Cache-Control", "no-store")
		w.Header().Add("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; sandbox")

		next.ServeHTTP(w, r)
	})
}

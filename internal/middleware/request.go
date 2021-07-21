package middleware

import "net/http"

// NewEnsureHTTPMethod returns a middleware func configured to reject
// requests not matching the specified method, returning a 405 status
// code and stopping the chain.
func NewEnsureHTTPMethod(method string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				http.Error(w, "", http.StatusMethodNotAllowed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

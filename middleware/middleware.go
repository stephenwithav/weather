package middleware

import "net/http"

// ProtectMiddleware
func ProtectMiddleware(next http.Handler) http.HandlerFunc {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}

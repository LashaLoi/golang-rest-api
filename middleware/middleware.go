package middleware

import "net/http"

// SetResponseHeader ...
func SetResponseHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next(w, r)
	}
}

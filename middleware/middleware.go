package middleware

import (
	"log"
	"net/http"
)

// SetResponseHeader ...
func SetResponseHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next(w, r)
	}
}

// Logger ...
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Method: %v\n", r.Method)
		log.Printf("URL: %v\n", r.URL)

		next(w, r)
	}
}

// ComposeMiddleware ...
func ComposeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return Logger(SetResponseHeader(next))
}

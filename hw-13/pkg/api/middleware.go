package api

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// loggMiddleware log request to API
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RemoteAddr, r.RequestURI, r.Context().Value("requestID"))

		next.ServeHTTP(w, r)
	})
}

// requestIDMiddleware assigns unique ID for request
func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "requestID", uuid.New())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// headersMiddleware makes a response via JSON
func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		next.ServeHTTP(w, r)
	})
}

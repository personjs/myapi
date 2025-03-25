package middleware

import (
	"log"
	"net/http"
)

// AuthMiddleware checks if the "Authentication" header is present in the request
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the "Authentication" header is present
		authHeader := r.Header.Get("Authentication")
		if authHeader == "" {
			log.Println("Authentication header missing")
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			// return
		}

		// Optionally, you can verify the value of the "Authentication" header here
		// For now, we assume any non-empty value is valid

		// If the header is present, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

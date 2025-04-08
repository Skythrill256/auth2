package utils

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	UserEmailContextKey contextKey = "user_email"
)

func AuthMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			// Check if the header has the Bearer prefix
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			// Extract and validate the JWT token
			token := parts[1]
			email, err := ParseJWT(token, jwtSecret)
			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			// Set the user email in the request context
			ctx := context.WithValue(r.Context(), UserEmailContextKey, email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserEmailFromContext retrieves the user email from the request context
func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(UserEmailContextKey).(string)
	return email, ok
}

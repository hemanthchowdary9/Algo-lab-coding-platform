package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var secretKey = []byte("ABCDEFG")

// JWTAuthMiddleware checks for valid JWT token
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Extract user info from claims if token is valid
			username := claims["sub"].(string)

			// Optionally, you could store the username in the request context
			ctx := context.WithValue(r.Context(), "username", username)
			r = r.WithContext(ctx)

			// Continue to the next handler
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
		}
	})
}

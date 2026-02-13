package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type contextKey string

const UserIdContextKey contextKey = "userId"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.SplitAfter(r.Header.Get("Authorization"), "Bearer ")[1]

		var claims jwt.RegisteredClaims
		_, _, err := jwt.NewParser().ParseUnverified(token, &claims)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if claims.Subject == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIdContextKey, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package middleware

import (
	"context"
	"net/http"
	"ozonTech/internal/pkg/auth/usecase"
)

type contextKey string

const (
	authorizationHeader            = "Authorization"
	userContextKey      contextKey = "userID"
)

func AuthMiddleware(authUsecase *usecase.AuthUsecase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				tokenStr := r.Header.Get(authorizationHeader)
				if tokenStr != "" {
					claims, err := authUsecase.VerifyJWT(tokenStr)
					if err != nil {
						http.Error(w, "invalid or expired token", http.StatusUnauthorized)
						return
					}

					ctx := context.WithValue(r.Context(), userContextKey, claims.UserID)
					r = r.WithContext(ctx)
				} else {
					http.Error(w, "missing auth token", http.StatusUnauthorized)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

// GetUserIDFromContext извлекает UserID из контекста
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(userContextKey).(int)
	return userID, ok
}

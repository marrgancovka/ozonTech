package middleware

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"ozonTech/internal/utils"
)

type contextKey string

const (
	authorizationHeader            = "Authorization"
	userContextKey      contextKey = "userID"
)

func AuthMiddleware(next http.Handler, log *logrus.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Info("Method not allowed for authentication")
			return
		}
		token := r.Header.Get("Authorization")
		claims, err := utils.ParseToken(token)
		if err != nil {
			log.Error("Error parsing token: ", err.Error())
			http.Error(w, "missing auth token", http.StatusUnauthorized)
			return
		}
		id, err := utils.ParseClaims(claims)
		if err != nil {
			log.Error("Error parsing claims: ", err.Error())
			http.Error(w, "missing auth token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userContextKey, id)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// GetUserIDFromContext извлекает UserID из контекста
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(userContextKey).(int)
	return userID, ok
}

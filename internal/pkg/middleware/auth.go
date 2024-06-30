package middleware

import (
	"context"
	"encoding/json"
	"github.com/99designs/gqlgen/graphql"
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
		if r.Method == http.MethodPost {
			token := r.Header.Get(authorizationHeader)
			claims, err := utils.ParseToken(token)
			if err == nil {
				id, err := utils.ParseClaims(claims)
				if err == nil {
					ctx := context.WithValue(r.Context(), userContextKey, id)
					r = r.WithContext(ctx)
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

// GetUserIDFromContext извлекает UserID из контекста
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(userContextKey).(int)
	return userID, ok
}

func AuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	httpReq := graphql.GetRequestContext(ctx)
	token := httpReq.Headers.Get(authorizationHeader)
	claims, err := utils.ParseToken(token)
	if err != nil {
		return nil, err
	}
	id, err := utils.ParseClaims(claims)
	if err != nil {
		return nil, err
	}

	// Обновление контекста с идентификатором пользователя
	ctx = context.WithValue(ctx, userContextKey, id)

	// Вызов следующего резолвера с обновленным контекстом
	return next(ctx)
}

type errorResponse struct {
	Message string `json:"message"`
}

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse{Message: message})
}

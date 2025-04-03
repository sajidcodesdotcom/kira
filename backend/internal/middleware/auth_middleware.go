package middleware

import (
	"context"
	"net/http"

	"github.com/sajidcodesdotcom/kira/internal/auth"
	"github.com/sajidcodesdotcom/kira/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := auth.ExtractTokenFromRequest(r)
		if err != nil {
			utils.RespondWithError(w, "Unauthorized, failed to get auth token from Bearer or http cookie"+err.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			utils.RespondWithError(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.ID)
		ctx = context.WithValue(r.Context(), "username", claims.Username)
		ctx = context.WithValue(r.Context(), "role", claims.Role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

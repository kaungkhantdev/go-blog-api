package middleware

import (
	"net/http"
	"strings"

	"go-blog-api/pkg/jwt"
	"go-blog-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			utils.ErrorResponse(ctx, "Authorization header required", http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.ErrorResponse(ctx, "Bearer token required", http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		claims, err := jwt.VerifyJWT(tokenString, jwt.GetJWTSecret())
		if err != nil {
			utils.ErrorResponse(ctx, "Invalid token", http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		// Optionally, you can set claims (like the username) in the context for downstream handlers
		ctx.Set("userId", claims.UserId)

		// Call the next handler in the chain
		ctx.Next()
	}
}

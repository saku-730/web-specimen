// internal/middleware/auth_middleware.go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/saku-730/web-specimen/backend/internal/handler"
)

// JWT token certification
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "認証ヘッダーが必要"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "トークンの形式が 'Bearer <token>' ではない"})
			return
		}
		tokenString := parts[1]

		claims := &handler.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return handler.jwtSecret, nil
		})

		if err != nil || !token.Valid {
			//401 code
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid or expired"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userName", claims.UserName)
		c.Next()
	}
}

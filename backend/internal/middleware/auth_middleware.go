// internal/middleware/auth_middleware.go
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/saku-730/web-specimen/backend/internal/model" // modelを参照する
)

// AuthMiddleware のインターフェースを定義するのだ
type AuthMiddleware interface {
	Auth() gin.HandlerFunc
}

// authMiddleware 構造体が秘密鍵を保管する場所になるのだ
type authMiddleware struct {
	jwtSecret []byte
}

// NewAuthMiddleware は秘密鍵を受け取って、ミドルウェアのインスタンスを生成するのだ
func NewAuthMiddleware(secret string) AuthMiddleware {
	return &authMiddleware{jwtSecret: []byte(secret)}
}

// Auth メソッドが、実際のミドルウェア処理 (gin.HandlerFunc) を返すのだ
func (m *authMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "認証ヘッダーが必要なのだ"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "トークンの形式が 'Bearer <token>' ではないのだ"})
			return
		}
		tokenString := parts[1]

		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// ★★★ ここが修正点！ ★★★
			// 構造体が持っている秘密鍵 m.jwtSecret を使うのだ！
			return m.jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "トークンが無効または期限切れなのだ"})
			return
		}

		c.Set("userID", int(claims.UserID))
		c.Set("userName", claims.UserName)
		c.Next()
	}
}

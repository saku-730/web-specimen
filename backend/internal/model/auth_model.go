// internal/model/auth_model.go
package model

import "github.com/golang-jwt/jwt/v5"

// LoginRequest はログイン時にクライアントから受け取るJSONの形なのだ
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse はログイン成功時にクライアントに返すJSON
type LoginResponse struct {
	Token string `json:"token"`
}

// Claims はJWTに埋め込む情報の構造体
type Claims struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

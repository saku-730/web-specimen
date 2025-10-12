// internal/service/auth_service.go
package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/saku-730/web-specimen/backend/config"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"github.com/saku-730/web-specimen/backend/internal/repository"
	"github.com/saku-730/web-specimen/backend/internal/util"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService interface {
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret []byte
}

func NewAuthService(userRepo repository.UserRepository, cfg *configs.Config) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: []byte(cfg.JWTSecret), // get secret key from cfg
	}
}

func (s *authService) Login(email, password string) (string, error) {
	// 1. Repositoryを使ってEntityを取得
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// 2. utilを使ってパスワードを検証
	if !util.CheckPasswordHash(password, *user.Password) {
		// パスワードが一致しない場合
		return "", ErrInvalidCredentials
	}

	// 3. 認証成功！JWTのClaimsを作成
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &model.Claims{
		UserID:   user.UserID,
		UserName: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// 4. トークンを生成して返す
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

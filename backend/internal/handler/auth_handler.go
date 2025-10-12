// internal/handler/auth_handler.go
package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saku-730/web-specimen/backend/internal/model"
	"github.com/saku-730/web-specimen/backend/internal/service"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

func (h *authHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "main address or password id different"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error occured"})
		return
	}

	// 成功したら、model.LoginResponseの形でトークンを返す
	c.JSON(http.StatusOK, model.LoginResponse{Token: token})
}

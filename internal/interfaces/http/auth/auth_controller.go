package auth

import (
	"errors"
	authUsecase "go-gin/internal/application/auth/usecase"
	dtoAuth "go-gin/internal/application/dto/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	LoginUsecase *authUsecase.LoginUsecase
}

func NewAuthController(loginUsecase *authUsecase.LoginUsecase) *AuthController {
	return &AuthController{LoginUsecase: loginUsecase}
}

// RegisterRoutes 註冊所有的路由
func (a *AuthController) RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", a.Login) // POST /auth/login
	}
}

// Login 處理用戶登錄請求
func (a *AuthController) Login(c *gin.Context) {
	var req dtoAuth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	res, err := a.LoginUsecase.Execute(&req)
	if err != nil {
		if errors.Is(err, authUsecase.ErrInvalidCredentials) || errors.Is(err, authUsecase.ErrUserNotFound) {
			log.Printf("Login attempt failed for %s: invalid credentials or user not found", req.EmailOrUsername)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			log.Printf("Internal server error during login for %s: %v", req.EmailOrUsername, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed due to an internal error"})
		}
		return
	}

	c.JSON(200, res) // 返回成功結果
}

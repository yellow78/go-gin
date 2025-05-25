package auth

import (
	"errors"
	authUsecase "go-gin/internal/application/auth/usecase" // Corrected import path
	dtoAuth "go-gin/internal/application/dto/auth"
	"log" // For server-side logging
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	loginUsecase *authUsecase.LoginUsecase
}

func NewAuthController(loginUsecase *authUsecase.LoginUsecase) *AuthController {
	return &AuthController{loginUsecase: loginUsecase}
}

func (ac *AuthController) Login(c *gin.Context) {
	var req dtoAuth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	response, err := ac.loginUsecase.Execute(&req)
	if err != nil {
		if errors.Is(err, authUsecase.ErrInvalidCredentials) || errors.Is(err, authUsecase.ErrUserNotFound) {
			// Log the attempt with a less specific message for security
			log.Printf("Login attempt failed for %s: invalid credentials or user not found", req.EmailOrUsername)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"}) // Generic message for client
		} else {
			// Log detailed error server-side
			log.Printf("Internal server error during login for %s: %v", req.EmailOrUsername, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed due to an internal error"})
		}
		return
	}
	c.JSON(http.StatusOK, response)
}

func (ac *AuthController) RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", ac.Login)
	}
}

package account

import (
	"go-gin/internal/application/account/usecase"
	"go-gin/internal/application/dto/account" // DTO for request binding

	"github.com/gin-gonic/gin" // Gin framework for HTTP handling
)

// UserController handles HTTP requests related to user accounts.
// It relies on UserUsecase to perform business logic.
type UserController struct {
	UserUsecase *usecase.UserUsecase // Injected dependency
}

// NewAccountController creates a new instance of UserController.
// It requires a UserUsecase to be injected.
func NewAccountController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

// RegisterRoutes sets up the routes for user account operations.
func (u *UserController) RegisterRoutes(router *gin.Engine) {
	users := router.Group("/account")
	{
		users.POST("", u.CreateUser) // Route for creating a new user
	}
}

// CreateUser handles the HTTP request for creating a new user.
// It binds the request body to CreateUserRequest DTO and calls the usecase.
func (u *UserController) CreateUser(c *gin.Context) {
	var req account.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload: " + err.Error()})
		return
	}

	// The UserUsecase is already initialized and available via u.UserUsecase
	// No need to initialize DB or repository here.
	res, err := u.UserUsecase.CreateUser(&req)
	if err != nil {
		// Consider more specific error handling based on error types from usecase
		c.JSON(500, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(200, res) // Return the successful result
}

package account

import (
	"go-gin/internal/application/account/usecase"
	"go-gin/internal/application/dto/account"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase *usecase.UserUsecase
}

func NewAccountController(userUsecase *usecase.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

// RegisterRoutes 註冊所有的路由
func (u *UserController) RegisterRoutes(router *gin.Engine) {
	users := router.Group("/account")
	{
		users.POST("", u.CreateUser) // POST /users
	}
}

// CreateUser 處理註冊用戶的請求
func (u *UserController) CreateUser(c *gin.Context) {
	var req account.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := u.UserUsecase.CreateUser(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(200, res) // 返回成功結果
}

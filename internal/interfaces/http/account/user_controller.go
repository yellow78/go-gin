package account

import (
	"go-gin/internal/application/account/usecase"
	"go-gin/internal/application/dto/account"
	"log"
	"time"

	pkgsql "go-gin/pkg/db"

	"go-gin/internal/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	mysqldriver "github.com/go-sql-driver/mysql"
)

type UserController struct {
	UserUsecase *usecase.UserUsecase
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

	mgr := pkgsql.NewGormManager()

	sqlconfig := mysqldriver.Config{
		User:                 "digimon",
		Passwd:               "digimon123",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "digimon_game",
		ParseTime:            true,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}

	err := mgr.InitDBWithConfig(&sqlconfig, "default")
	if err != nil {
		log.Fatalf("DB init failed: %v", err)
	}

	db, _ := mgr.GetDB("default")

	userRepo := persistence.NewUserRepository(db)
	u.UserUsecase = usecase.NewUserUsecase(userRepo)

	res, err := u.UserUsecase.CreateUser(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res) // 返回成功結果
}

package http

import (
	applicationService "go-gin/internal/application/usecase"
	domainService "go-gin/internal/domain/service"
	"go-gin/internal/infrastructure/persistence"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	userRepo := persistence.NewUserRepository()
	userService := domainService.NewUserService(userRepo)
	userAppService := applicationService.NewUserApplicationService(userService)
	r.GET("/users", func(c *gin.Context) {
		users := userAppService.GetUsers()
		c.JSON(http.StatusOK, users)
	})
	return r
}

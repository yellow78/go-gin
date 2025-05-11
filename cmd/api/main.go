package main

import (
	account "go-gin/internal/interfaces/http/account"

	"github.com/gin-gonic/gin"
)

func main() {
	// // 設置 Gin 框架為 debug 模式
	// gin.SetMode(gin.DebugMode)

	// router := http.NewRouter()
	// router.Run() // 預設在 0.0.0.0:8080 啟動服務

	router := gin.Default()
	account := &account.UserController{}
	account.RegisterRoutes(router)
	if err := router.Run(":8080"); err != nil {
		panic(err) // 啟動服務失敗時，拋出錯誤
	}
}

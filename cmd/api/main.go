package main

import (
	"go-gin/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 設置 Gin 框架為 debug 模式
	gin.SetMode(gin.DebugMode)

	router := http.NewRouter()
	router.Run() // 預設在 0.0.0.0:8080 啟動服務
}

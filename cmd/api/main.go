package main

import (
	"fmt"
	"log"
	"time"

	"go-gin/pkg/config"
	pkgsql "go-gin/pkg/db"

	mysqldriver "github.com/go-sql-driver/mysql"

	accountUsecase "go-gin/internal/application/account/usecase" // Usecase for account
	authUsecase "go-gin/internal/application/auth/usecase"       // Usecase for auth
	"go-gin/internal/infrastructure/persistence"
	accountHttp "go-gin/internal/interfaces/http/account" // HTTP handlers for account
	authHttp "go-gin/internal/interfaces/http/auth"       // HTTP handlers for auth
	"go-gin/internal/interfaces/http/middleware"          // HTTP middleware
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load application configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Auth Middleware (needs cfg for JWT secret)
	authMw := middleware.AuthMiddleware(cfg)

	// Initialize database connection
	// Note: Using mysqldriver.Config for db configuration
	dbConfig := mysqldriver.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Net:                  "tcp",
		Addr:                 cfg.DBHost + ":" + cfg.DBPort,
		DBName:               cfg.DBName,
		ParseTime:            true,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}

	gormManager := pkgsql.NewGormManager()
	// Initialize DB with the specific configuration
	if err := gormManager.InitDBWithConfig(&dbConfig, cfg.DBName); err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}

	// Get the default DB instance (or specific one if needed)
	defaultDB, ok := gormManager.GetDB(cfg.DBName) // Using DBName from config to get the specific DB
	if !ok {
		log.Fatalf("Failed to get database instance for %s", cfg.DBName)
	}

	// Initialize repositories
	// For the 'account' domain, we use UserRepository
	accountUserRepo := persistence.NewUserRepository(defaultDB)

	// Initialize use cases
	// The UserUsecase requires a UserRepository
	userUsecase := accountUsecase.NewUserUsecase(accountUserRepo)
	// The LoginUsecase requires UserRepository and AppConfig (for JWT secret)
	loginUsecase := authUsecase.NewLoginUsecase(accountUserRepo, cfg)

	// Initialize Gin router
	router := gin.Default()

	// Initialize controllers and register routes
	// Account controller
	accountController := accountHttp.NewAccountController(userUsecase)
	accountController.RegisterRoutes(router)

	// Auth controller
	authController := authHttp.NewAuthController(loginUsecase)
	authController.RegisterRoutes(router)

	// Public health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Example protected route group
	apiV1 := router.Group("/api/v1")
	apiV1.Use(authMw) // Apply middleware to all routes in this group
	{
		apiV1.GET("/me", func(c *gin.Context) {
			currentUser, exists := c.Get(middleware.ContextKeyUser)
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Current user not found in context"})
				return
			}

			userData, ok := currentUser.(middleware.UserContextData)
			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data type in context"})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message":  "Welcome to the protected area!",
				"user_id":  userData.ID,
				"username": userData.Username,
			})
		})
	}

	// Start the server
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Starting server on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err) // Use log.Fatalf for consistency
	}
}

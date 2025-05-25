package main

import (
	"fmt"
	"log"
	"net/http" // For http.StatusOK and other constants

	"go-gin/pkg/config"
	"go-gin/pkg/database/pkgsql"
	"go-gin/pkg/driver/mysqldriver"

	accountUsecase "go-gin/internal/application/account/usecase"
	authUsecase "go-gin/internal/application/auth/usecase"
	digimonUsecase "go-gin/internal/application/game/usecase" // New import for Digimon usecases

	"go-gin/internal/infrastructure/persistence"
	// digimonPersistence is not explicitly aliased as persistence is already used for account.
	// We will use persistence.NewMySQLDigimonRepository directly if it's clear.
	// If there were a naming conflict or for extreme clarity, an alias like:
	// digimonInfra "go-gin/internal/infrastructure/persistence"
	// could be used, but the current structure seems to allow direct use.

	accountHttp "go-gin/internal/interfaces/http/account"
	authHttp "go-gin/internal/interfaces/http/auth"
	digimonHttp "go-gin/internal/interfaces/http/game"       // New import for Digimon controller
	"go-gin/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load Config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// 2. Initialize DB
	dbConfig := mysqldriver.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		Username: cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	}
	gormManager := pkgsql.NewGormManager()
	if err := gormManager.InitDBWithConfig(&dbConfig); err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	defaultDB, err := gormManager.GetDB(cfg.DBName)
	if err != nil {
		log.Fatalf("Failed to get database instance for %s: %v", cfg.DBName, err)
	}

	// 3. Setup Account components
	accountUserRepo := persistence.NewUserRepository(defaultDB)
	userUsecase := accountUsecase.NewUserUsecase(accountUserRepo)
	accountController := accountHttp.NewAccountController(userUsecase)

	// 4. Setup Auth components
	loginUsecase := authUsecase.NewLoginUsecase(accountUserRepo, cfg)
	authController := authHttp.NewAuthController(loginUsecase)

	// 5. Setup Gin Router
	router := gin.Default()

	// 6. Setup Public Routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// 7. Register Account and Auth routes (public or semi-public)
	accountController.RegisterRoutes(router) // Assuming /account for user creation is public
	authController.RegisterRoutes(router)    // Assuming /auth for login is public

	// 8. Setup Auth Middleware
	authMw := middleware.AuthMiddleware(cfg)

	// 9. Setup Protected Route Group (API v1)
	apiV1 := router.Group("/api/v1")
	apiV1.Use(authMw)
	{
		// Register protected test route (/api/v1/me)
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

		// === NEWLY ADDED FOR DIGIMON ===
		// 10. Instantiate Digimon Repositories
		// Using the existing 'persistence' import for NewMySQLDigimonSpeciesRepository and NewMySQLDigimonRepository
		digimonSpeciesRepo := persistence.NewMySQLDigimonSpeciesRepository(defaultDB)
		digimonRepo := persistence.NewMySQLDigimonRepository(defaultDB)

		// 11. Instantiate Digimon Usecases
		acquireDigimonUC := digimonUsecase.NewAcquireDigimonUsecase(digimonRepo, digimonSpeciesRepo)
		listPlayerDigimonUC := digimonUsecase.NewListPlayerDigimonUsecase(digimonRepo, digimonSpeciesRepo)
		getDigimonDetailsUC := digimonUsecase.NewGetDigimonDetailsUsecase(digimonRepo, digimonSpeciesRepo)
		// Constructor for UpdateDigimonNicknameUsecase already includes digimonSpeciesRepo from previous step
		updateDigimonNicknameUC := digimonUsecase.NewUpdateDigimonNicknameUsecase(digimonRepo, digimonSpeciesRepo)

		// 12. Instantiate DigimonController
		digimonAPIController := digimonHttp.NewDigimonController(
			acquireDigimonUC,
			listPlayerDigimonUC,
			getDigimonDetailsUC,
			updateDigimonNicknameUC,
		)

		// 13. Register DigimonController Routes under the protected group
		digimonAPIController.RegisterRoutes(apiV1) // Pass the protected group apiV1
		// === END OF NEWLY ADDED FOR DIGIMON ===
	}

	// 14. Start Server
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Starting server on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

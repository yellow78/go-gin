package game // Package name for the controller

import (
	"errors" // Added for errors.Is
	"net/http"
	"strconv" // For parsing limit/offset from query params

	dto "go-gin/internal/application/dto/game"
	usecase "go-gin/internal/application/game/usecase"
	"go-gin/internal/interfaces/http/middleware" // For UserContextData and ContextKeyUser

	"github.com/gin-gonic/gin"
	// "log" // For server-side logging if needed
)

// DigimonController handles HTTP requests related to Digimon management.
type DigimonController struct {
	acquireDigimonUsecase      *usecase.AcquireDigimonUsecase
	listPlayerDigimonUsecase   *usecase.ListPlayerDigimonUsecase
	getDigimonDetailsUsecase   *usecase.GetDigimonDetailsUsecase
	updateDigimonNicknameUsecase *usecase.UpdateDigimonNicknameUsecase
}

// NewDigimonController creates a new DigimonController.
func NewDigimonController(
	acquireUC *usecase.AcquireDigimonUsecase,
	listUC *usecase.ListPlayerDigimonUsecase,
	getDetailsUC *usecase.GetDigimonDetailsUsecase,
	updateNicknameUC *usecase.UpdateDigimonNicknameUsecase,
) *DigimonController {
	return &DigimonController{
		acquireDigimonUsecase:      acquireUC,
		listPlayerDigimonUsecase:   listUC,
		getDigimonDetailsUsecase:   getDetailsUC,
		updateDigimonNicknameUsecase: updateNicknameUC,
	}
}

// RegisterRoutes registers all Digimon management routes.
// All routes in this group should be protected by the AuthMiddleware.
func (dc *DigimonController) RegisterRoutes(routerGroup *gin.RouterGroup) {
	// Assuming routerGroup is already protected by AuthMiddleware (e.g., /api/v1)
	digimonRoutes := routerGroup.Group("/my/digimon") // Player-specific Digimon actions
	{
		digimonRoutes.POST("", dc.AcquireDigimon)
		digimonRoutes.GET("", dc.ListPlayerDigimon)
		digimonRoutes.GET("/:digimonID", dc.GetMyDigimonDetails) // Specific to player's collection
		digimonRoutes.PUT("/:digimonID/nickname", dc.UpdateDigimonNickname)
	}
	// Potentially a more public route for species if needed, e.g., routerGroup.GET("/digimon-species/:speciesID", dc.GetSpeciesDetails)
}

// Helper to get authenticated player ID from context
func getPlayerIDFromContext(c *gin.Context) (string, bool) {
	userCtxData, exists := c.Get(middleware.ContextKeyUser)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated (no context)"})
		return "", false
	}
	userData, ok := userCtxData.(middleware.UserContextData)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data in context"})
		return "", false
	}
	if userData.ID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return "", false
	}
	return userData.ID, true
}

// POST /api/v1/my/digimon
func (dc *DigimonController) AcquireDigimon(c *gin.Context) {
	playerID, ok := getPlayerIDFromContext(c)
	if !ok { return }

	var req dto.AcquireDigimonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	digimonResponse, err := dc.acquireDigimonUsecase.Execute(playerID, &req)
	if err != nil {
		// Handle specific errors from usecase, e.g., ErrSpeciesNotFound
		if errors.Is(err, usecase.ErrSpeciesNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if errors.Is(err, usecase.ErrMaxDigimonSlotsReached) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			// log.Printf("Error acquiring digimon: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to acquire digimon"})
		}
		return
	}
	c.JSON(http.StatusCreated, digimonResponse)
}

// GET /api/v1/my/digimon
func (dc *DigimonController) ListPlayerDigimon(c *gin.Context) {
	playerID, ok := getPlayerIDFromContext(c)
	if !ok { return }

	offsetQuery := c.DefaultQuery("offset", "0")
	limitQuery := c.DefaultQuery("limit", "10")
	offset, _ := strconv.Atoi(offsetQuery)
	limit, _ := strconv.Atoi(limitQuery)
	if limit > 100 { limit = 100 } // Max limit

	digimonList, _ /*totalCount*/, err := dc.listPlayerDigimonUsecase.Execute(playerID, offset, limit)
	if err != nil {
		// log.Printf("Error listing player digimon: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list player digimon"})
		return
	}
	// TODO: Consider how to include totalCount in response for pagination if available
	c.JSON(http.StatusOK, gin.H{"digimon": digimonList /*, "total": totalCount */})
}

// GET /api/v1/my/digimon/:digimonID
func (dc *DigimonController) GetMyDigimonDetails(c *gin.Context) {
	playerID, ok := getPlayerIDFromContext(c)
	if !ok { return }
	
	digimonID := c.Param("digimonID")
	if digimonID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Digimon ID is required"})
		return
	}

	digimonResponse, err := dc.getDigimonDetailsUsecase.Execute(digimonID, playerID)
	if err != nil {
		if errors.Is(err, usecase.ErrDigimonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if errors.Is(err, usecase.ErrDigimonNotOwned) {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()}) // Or StatusNotFound to not reveal existence
		} else if errors.Is(err, usecase.ErrSpeciesNotFound) { // If species not found for an existing digimon
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving digimon details (species missing)"})
		} else {
			// log.Printf("Error getting digimon details: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get digimon details"})
		}
		return
	}
	c.JSON(http.StatusOK, digimonResponse)
}

// PUT /api/v1/my/digimon/:digimonID/nickname
func (dc *DigimonController) UpdateDigimonNickname(c *gin.Context) {
	playerID, ok := getPlayerIDFromContext(c)
	if !ok { return }

	digimonID := c.Param("digimonID")
	if digimonID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Digimon ID is required"})
		return
	}

	var req dto.UpdateDigimonNicknameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	digimonResponse, err := dc.updateDigimonNicknameUsecase.Execute(playerID, digimonID, req.Nickname)
	if err != nil {
	   if errors.Is(err, usecase.ErrDigimonNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	   } else if errors.Is(err, usecase.ErrDigimonNotOwned) {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()}) // Or StatusNotFound
	   } else {
			// log.Printf("Error updating digimon nickname: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update nickname"})
	   }
	   return
	}
	c.JSON(http.StatusOK, digimonResponse)
}

package middleware

import (
	"errors"
	"go-gin/pkg/config" // For AppConfig to get JWT secret
	"log"               // For server-side logging of errors
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthorizationHeaderKey  = "Authorization"
	AuthorizationTypeBearer = "bearer"
	ContextKeyUser          = "currentUser"
)

type UserContextData struct {
	ID       string
	Username string
	// Role string // Example for future use
}

func AuthMiddleware(appConfig *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeaderKey)
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 || !strings.EqualFold(fields[0], AuthorizationTypeBearer) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format. Expected 'Bearer <token>'."})
			return
		}

		accessToken := fields[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Printf("Unexpected signing method: %v", token.Header["alg"])
				return nil, errors.New("unexpected signing method")
			}
			return []byte(appConfig.JWTSecretKey), nil
		})

		if err != nil {
			log.Printf("Token parsing error: %v. Token: %s", err, accessToken)
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			return
		}

		userID, okUserID := claims["sub"].(string)
		username, okUsername := claims["username"].(string)
		// role, _ := claims["role"].(string) // Example for future

		if !okUserID || !okUsername {
			log.Printf("Invalid token claims: UserID present: %t, Username present: %t. Claims: %+v", okUserID, okUsername, claims)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		userCtxData := UserContextData{
			ID:       userID,
			Username: username,
			// Role: role,
		}
		c.Set(ContextKeyUser, userCtxData)

		c.Next()
	}
}

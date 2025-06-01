package middleware

import (
	"errors"
	"go-gin/pkg/config"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthorizationHeaderKey  = "Authorization"
	AuthorizationTypeBearer = "Bearer"
	ContextKeyUser          = "currentUser"
)

type UserContextData struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

func AuthMiddleware(appConfig *config.AppConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		field := strings.Fields(authHeader)
		if len(field) < 2 || !strings.EqualFold(field[0], AuthorizationTypeBearer) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format. Expected 'Bearer <token>'."})
			return
		}

		accessToken := field[1]
		// jwt 自訂儲存任意 key/value
		claims := jwt.MapClaims{}

		// 解析並驗證 JWT token
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
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
			} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is invalid"})
			return
		}

		userID, okUserID := claims["sub"].(string)
		username, okUsername := claims["username"].(string)

		if !okUserID || !okUsername {
			log.Printf("Invalid token claims: UserID present: %t, Username present: %t. Claims: %+v", okUserID, okUsername, claims)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		userCtxData := UserContextData{
			UserID:   userID,
			Username: username,
			// Role: role,
		}
		ctx.Set(ContextKeyUser, userCtxData)

		ctx.Next()
	}
}

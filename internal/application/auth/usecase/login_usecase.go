package usecase

import (
	"errors"
	"go-gin/internal/application/dto/auth"
	accountModel "go-gin/internal/domain/account/model/entity"
	accountRepo "go-gin/internal/domain/account/repository"
	"go-gin/pkg/config" // For AppConfig to get JWT secret
	"log"               // Added for logging token signing errors
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email/username or password")
	ErrUserNotFound       = errors.New("user not found")
	// jwtIssuer (can be a constant or from config)
	jwtIssuer = "go-gin-digimon-api"
)

type LoginUsecase struct {
	userRepo  accountRepo.UserRepository
	appConfig *config.AppConfig // To access JWTSecretKey
}

func NewLoginUsecase(userRepo accountRepo.UserRepository, appConfig *config.AppConfig) *LoginUsecase {
	return &LoginUsecase{userRepo: userRepo, appConfig: appConfig}
}

func (uc *LoginUsecase) Execute(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	var user *accountModel.User
	var err error

	// Attempt to fetch user by email first
	user, err = uc.userRepo.GetByEmail(req.EmailOrUsername)
	if err != nil {
		// If not found by email, try by username
		// This also handles cases where GetByEmail might return other errors,
		// though GORM's ErrRecordNotFound is typical for "not found".
		user, err = uc.userRepo.GetByUserName(req.EmailOrUsername)
		if err != nil {
			// If still not found or another error occurs, return ErrUserNotFound.
			// This helps prevent user enumeration by giving a generic "not found"
			// regardless of whether the lookup was by email or username, or if an error occurred.
			return nil, ErrUserNotFound
		}
	}
	
	// Additional safeguard: if user is still nil after the lookups (e.g., if a repository
	// returns nil, nil for "not found" instead of an error), explicitly return ErrUserNotFound.
	if user == nil {
	    return nil, ErrUserNotFound
	}


	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		// If passwords don't match, or if there's an error during comparison (e.g., hash format issue)
		return nil, ErrInvalidCredentials
	}

	// Create JWT claims
	claims := jwt.MapClaims{
		"sub":      user.ID,                                 // Subject (user ID)
		"iss":      jwtIssuer,                               // Issuer
		"iat":      time.Now().Unix(),                       // Issued at
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Expiration time (7 days)
		"username": user.Username,                           // Custom claim for username
	}

	// Create token with claims and sign it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(uc.appConfig.JWTSecretKey))
	if err != nil {
		// Log the error for server-side diagnosis
		log.Printf("Error signing token for user %s: %v", user.ID, err)
		return nil, errors.New("failed to generate token") // Generic error for client
	}

	return &auth.LoginResponse{
		Token:    tokenString,
		UserID:   user.ID,
		Username: user.Username,
	}, nil
}

package usecase

import (
	"errors"
	"log"
	"time"

	"go-gin/internal/application/dto/auth"
	accountRepo "go-gin/internal/domain/account/repository"
	"go-gin/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email/username or password")
	ErrUserNotFound       = errors.New("user not found")
	jwtIssuer             = "go-gin-digimon-api"
)

type LoginUsecase struct {
	userRepo  accountRepo.UserRepository
	appConfig *config.AppConfig
}

func NewLoginUsecase(userRepo accountRepo.UserRepository, appConfig *config.AppConfig) *LoginUsecase {
	return &LoginUsecase{
		userRepo:  userRepo,
		appConfig: appConfig,
	}
}

func (lu *LoginUsecase) Execute(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	// 1. 檢查使用者Email是否存在
	user, err := lu.userRepo.GetByEmail(req.EmailOrUsername)
	if err != nil {
		// 如果找不到使用者，則嘗試使用使用者名稱
		user, err = lu.userRepo.GetByUserName(req.EmailOrUsername)
		if err != nil {
			return nil, ErrUserNotFound
		}
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	// 2. 驗證密碼
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// 3. 生成 JWT claims
	claims := jwt.MapClaims{
		"sub":      user.ID,                                   // Subject (user ID)
		"iss":      jwtIssuer,                                 // Issuer
		"iat":      time.Now().Unix(),                         // Issued at
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Expiration time (7 days)
		"username": user.Username,                             // Custom claim for username
	}

	// 4. 創建 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(lu.appConfig.JWTSecretKey))
	if err != nil {
		log.Printf("Error signing token for user %s: %v", user.ID, err)
		return nil, errors.New("failed to generate token")
	}

	return &auth.LoginResponse{
		Token:    tokenString,
		UserID:   user.ID,
		Username: user.Username,
	}, nil
}

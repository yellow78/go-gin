package usecase

import (
	"go-gin/internal/application/dto/account"
	"go-gin/internal/domain/account/model/entity"
	"go-gin/internal/domain/account/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (u *UserUsecase) CreateUser(req *account.CreateUserRequest) (*account.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:               uuid.NewString(),
		Username:         req.Username,
		Email:            req.Email,
		PasswordHash:     string(hashedPassword),
		CreatedAt:        time.Now(),
		InGameName:       req.Username, // Default InGameName to Username as per instruction
		Level:            1,            // Default Level
		ExperiencePoints: 0,            // Default ExperiencePoints
	}

	if err := u.UserRepo.Create(user); err != nil {
		return nil, err
	}

	return &account.UserResponse{
		ID:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		CreatedAt:        user.CreatedAt.Format(time.RFC3339),
		InGameName:       user.InGameName,
		Level:            user.Level,
		ExperiencePoints: user.ExperiencePoints,
		LastLoginAt:      formatLastLoginAt(user.LastLoginAt),
	}, nil
}

// formatLastLoginAt formats the LastLoginAt time.
// If the time is zero (meaning it hasn't been set), it returns an empty string.
// This works well with `omitempty` in the DTO.
func formatLastLoginAt(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

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
		ID:           uuid.NewString(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
	}

	if err := u.UserRepo.Create(user); err != nil {
		return nil, err
	}

	return &account.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
	}, nil
}

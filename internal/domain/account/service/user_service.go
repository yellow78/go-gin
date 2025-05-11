package service

import (
	"go-gin/internal/domain/account/repository"
)

type UserDomainService struct {
	UserRepo repository.UserRepository
}

func NewUserDomainService(userRepo repository.UserRepository) *UserDomainService {
	return &UserDomainService{UserRepo: userRepo}
}

func (u *UserDomainService) existUsernameOrEmail(username, email string) (bool, error) {
	if user, _ := u.UserRepo.GetByUserName(username); user != nil {
		return true, nil
	}

	if user, _ := u.UserRepo.GetByEmail(email); user != nil {
		return true, nil
	}

	return false, nil
}

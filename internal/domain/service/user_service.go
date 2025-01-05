package service

import (
	"go-gin/internal/domain/model/entity"
	"go-gin/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() []entity.User {
	return s.repo.GetAll()
}

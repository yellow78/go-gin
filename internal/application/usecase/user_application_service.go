package service

import (
	"go-gin/internal/application/dto"
	"go-gin/internal/domain/service"
)

type UserApplicationService struct {
	userService *service.UserService
}

func NewUserApplicationService(userService *service.UserService) *UserApplicationService {
	return &UserApplicationService{userService: userService}
}

func (u *UserApplicationService) GetUsers() []dto.UserDTO {
	users := u.userService.GetAllUsers()
	var userDTOs []dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, dto.UserDTO{ID: user.ID, Name: user.Name})
	}
	return userDTOs
}

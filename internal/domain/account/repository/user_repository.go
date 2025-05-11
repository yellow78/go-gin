package repository

import "go-gin/internal/domain/account/model/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetUserByID(id string) (*entity.User, error)
	GetByUserName(userName string) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Delete(id string) error
}

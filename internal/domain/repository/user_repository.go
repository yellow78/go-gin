package repository

import (
	"go-gin/internal/domain/model/entity"
)

type UserRepository interface {
	GetAll() []entity.User
}

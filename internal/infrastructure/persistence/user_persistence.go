package persistence

import "go-gin/internal/domain/model/entity"

type UserTestRepository struct{}

func NewUserTestRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAll() []entity.User {
	return []entity.User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Doe"},
	}
}

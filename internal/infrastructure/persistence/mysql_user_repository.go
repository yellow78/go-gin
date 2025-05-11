package persistence

import (
	"go-gin/internal/domain/account/model/entity"
	"go-gin/internal/domain/account/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// 檢查型別相容性
var _ repository.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *entity.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByUserName(userName string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("username = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Delete(id string) error {
	var user entity.User
	if err := r.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

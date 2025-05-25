package persistence

import (
	"go-gin/internal/domain/game/model/entity"
	"go-gin/internal/domain/game/repository"
	"gorm.io/gorm"
)

type mysqlDigimonMoveRepository struct {
	db *gorm.DB
}

func NewMySQLDigimonMoveRepository(db *gorm.DB) repository.DigimonMoveRepository {
	return &mysqlDigimonMoveRepository{db: db}
}

func (r *mysqlDigimonMoveRepository) FindByID(id string) (*entity.DigimonMove, error) {
	var move entity.DigimonMove
	if err := r.db.First(&move, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &move, nil
}

func (r *mysqlDigimonMoveRepository) FindByIDs(ids []string) ([]*entity.DigimonMove, error) {
	var moves []*entity.DigimonMove
	if err := r.db.Where("id IN ?", ids).Find(&moves).Error; err != nil {
		return nil, err
	}
	return moves, nil
}

func (r *mysqlDigimonMoveRepository) FindByName(name string) (*entity.DigimonMove, error) {
	var move entity.DigimonMove
	if err := r.db.First(&move, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &move, nil
}

func (r *mysqlDigimonMoveRepository) ListAll(offset int, limit int) ([]*entity.DigimonMove, int64, error) {
	var moves []*entity.DigimonMove
	var totalCount int64

	if err := r.db.Model(&entity.DigimonMove{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}
	
	if err := r.db.Offset(offset).Limit(limit).Order("name ASC").Find(&moves).Error; err != nil {
		return nil, totalCount, err // Return count even if list fetch fails, or handle differently
	}
	return moves, totalCount, nil
}

func (r *mysqlDigimonMoveRepository) Create(move *entity.DigimonMove) error {
	return r.db.Create(move).Error
}

func (r *mysqlDigimonMoveRepository) Update(move *entity.DigimonMove) error {
	return r.db.Save(move).Error
}

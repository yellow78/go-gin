package persistence

import (
	"go-gin/internal/domain/game/model/entity"
	"go-gin/internal/domain/game/repository"
	"gorm.io/gorm"
)

type mysqlDigimonEvolutionRepository struct {
	db *gorm.DB
}

func NewMySQLDigimonEvolutionRepository(db *gorm.DB) repository.DigimonEvolutionRepository {
	return &mysqlDigimonEvolutionRepository{db: db}
}

func (r *mysqlDigimonEvolutionRepository) FindByID(id string) (*entity.DigimonEvolution, error) {
	var evolution entity.DigimonEvolution
	if err := r.db.First(&evolution, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &evolution, nil
}

func (r *mysqlDigimonEvolutionRepository) FindByFromSpeciesID(fromSpeciesID string) ([]*entity.DigimonEvolution, error) {
	var evolutions []*entity.DigimonEvolution
	if err := r.db.Where("from_species_id = ?", fromSpeciesID).Find(&evolutions).Error; err != nil {
		return nil, err
	}
	return evolutions, nil
}

func (r *mysqlDigimonEvolutionRepository) FindByToSpeciesID(toSpeciesID string) ([]*entity.DigimonEvolution, error) {
	var evolutions []*entity.DigimonEvolution
	if err := r.db.Where("to_species_id = ?", toSpeciesID).Find(&evolutions).Error; err != nil {
		return nil, err
	}
	return evolutions, nil
}

func (r *mysqlDigimonEvolutionRepository) ListAll(offset int, limit int) ([]*entity.DigimonEvolution, int64, error) {
	var evolutions []*entity.DigimonEvolution
	var totalCount int64

	if err := r.db.Model(&entity.DigimonEvolution{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(limit).Find(&evolutions).Error; err != nil { // Add OrderBy if needed
		return nil, totalCount, err
	}
	return evolutions, totalCount, nil
}

func (r *mysqlDigimonEvolutionRepository) Create(evolution *entity.DigimonEvolution) error {
	return r.db.Create(evolution).Error
}

func (r *mysqlDigimonEvolutionRepository) Update(evolution *entity.DigimonEvolution) error {
	return r.db.Save(evolution).Error
}

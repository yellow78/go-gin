package persistence

import (
	"go-gin/internal/domain/game/model/entity"
	"go-gin/internal/domain/game/repository"
	"gorm.io/gorm"
)

type mysqlDigimonSpeciesRepository struct {
	db *gorm.DB
}

// NewMySQLDigimonSpeciesRepository creates a new GORM-based DigimonSpeciesRepository.
func NewMySQLDigimonSpeciesRepository(db *gorm.DB) repository.DigimonSpeciesRepository {
	return &mysqlDigimonSpeciesRepository{db: db}
}

func (r *mysqlDigimonSpeciesRepository) FindByID(id string) (*entity.DigimonSpecies, error) {
	var species entity.DigimonSpecies
	err := r.db.First(&species, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &species, nil
}

func (r *mysqlDigimonSpeciesRepository) FindByName(name string) (*entity.DigimonSpecies, error) {
	var species entity.DigimonSpecies
	err := r.db.First(&species, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return &species, nil
}

func (r *mysqlDigimonSpeciesRepository) ListAll(offset int, limit int) ([]*entity.DigimonSpecies, error) {
	var speciesList []*entity.DigimonSpecies
	// Add OrderBy if needed, e.g. .Order("name ASC")
	err := r.db.Offset(offset).Limit(limit).Find(&speciesList).Error
	if err != nil {
		return nil, err
	}
	return speciesList, nil
}

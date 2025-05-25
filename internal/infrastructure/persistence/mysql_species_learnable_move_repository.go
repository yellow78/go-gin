package persistence

import (
	"go-gin/internal/domain/game/model/entity"
	"go-gin/internal/domain/game/repository"
	"gorm.io/gorm"
)

type mysqlSpeciesLearnableMoveRepository struct {
	db *gorm.DB
}

func NewMySQLSpeciesLearnableMoveRepository(db *gorm.DB) repository.SpeciesLearnableMoveRepository {
	return &mysqlSpeciesLearnableMoveRepository{db: db}
}

func (r *mysqlSpeciesLearnableMoveRepository) ListBySpeciesIDAndMaxLevel(speciesID string, maxLevel int) ([]*entity.SpeciesLearnableMove, error) {
	var learnableMoves []*entity.SpeciesLearnableMove
	// Find moves where required_level is less than or equal to maxLevel.
	// Also handles cases where required_level is NULL (e.g. for non-level up learn methods, assuming they are learned by default)
	// or explicitly check for learn_method if only 'LevelUp' moves are desired here.
	// For initial moves, typically it's level 1.
	err := r.db.Where("species_id = ? AND (required_level <= ? OR required_level IS NULL)", speciesID, maxLevel).
				Order("required_level ASC"). // Optional: order by level
				Find(&learnableMoves).Error
	if err != nil {
		return nil, err
	}
	return learnableMoves, nil
}

// Example Create method if needed later:
// func (r *mysqlSpeciesLearnableMoveRepository) Create(slm *entity.SpeciesLearnableMove) error {
//    return r.db.Create(slm).Error
// }

package repository

import "go-gin/internal/domain/game/model/entity"

type SpeciesLearnableMoveRepository interface {
	ListBySpeciesIDAndMaxLevel(speciesID string, maxLevel int) ([]*entity.SpeciesLearnableMove, error)
	// Add Create/Delete methods if these are managed administratively
	// Create(slm *entity.SpeciesLearnableMove) error
}

package repository

import "go-gin/internal/domain/game/model/entity"

type DigimonEvolutionRepository interface {
	FindByID(id string) (*entity.DigimonEvolution, error)
	FindByFromSpeciesID(fromSpeciesID string) ([]*entity.DigimonEvolution, error)
	FindByToSpeciesID(toSpeciesID string) ([]*entity.DigimonEvolution, error)
	ListAll(offset int, limit int) ([]*entity.DigimonEvolution, int64, error) // Returns list and total count
	Create(evolution *entity.DigimonEvolution) error
	Update(evolution *entity.DigimonEvolution) error
	// Delete(id string) error // Optional: if evolution paths can be deleted
}

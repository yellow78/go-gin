package repository

import "go-gin/internal/domain/game/model/entity" // Adjust import path

type DigimonRepository interface {
	Create(digimon *entity.Digimon) error
	FindByID(id string) (*entity.Digimon, error)
	ListByPlayerID(playerID string, offset int, limit int) ([]*entity.Digimon, error) // Added pagination
	Update(digimon *entity.Digimon) error
	Delete(id string) error // Potentially soft delete
	// CountByPlayerID(playerID string) (int64, error) // Useful for checking slots
}

package repository

import "go-gin/internal/domain/game/model/entity"

type DigimonMoveRepository interface {
	FindByID(id string) (*entity.DigimonMove, error)
	FindByIDs(ids []string) ([]*entity.DigimonMove, error)
	FindByName(name string) (*entity.DigimonMove, error)
	ListAll(offset int, limit int) ([]*entity.DigimonMove, int64, error) // Returns list and total count
	Create(move *entity.DigimonMove) error
	Update(move *entity.DigimonMove) error
	// Delete(id string) error // Optional: if moves can be deleted
}

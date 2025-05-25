package repository

import "go-gin/internal/domain/game/model/entity"

type ItemRepository interface {
	FindByID(id string) (*entity.Item, error)
	FindByIDs(ids []string) ([]*entity.Item, error)
	FindByName(name string) (*entity.Item, error)
	ListAll(offset int, limit int) ([]*entity.Item, int64, error) // Returns list and total count
	ListByType(itemType string, offset int, limit int) ([]*entity.Item, int64, error) // Returns list and total count
	Create(item *entity.Item) error
	Update(item *entity.Item) error
	// Delete(id string) error // Optional: if items can be deleted
}

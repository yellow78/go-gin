package persistence

import (
	"go-gin/internal/domain/game/model/entity"
	"go-gin/internal/domain/game/repository"
	"gorm.io/gorm"
)

type mysqlItemRepository struct {
	db *gorm.DB
}

func NewMySQLItemRepository(db *gorm.DB) repository.ItemRepository {
	return &mysqlItemRepository{db: db}
}

func (r *mysqlItemRepository) FindByID(id string) (*entity.Item, error) {
	var item entity.Item
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *mysqlItemRepository) FindByIDs(ids []string) ([]*entity.Item, error) {
	var items []*entity.Item
	if err := r.db.Where("id IN ?", ids).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *mysqlItemRepository) FindByName(name string) (*entity.Item, error) {
	var item entity.Item
	if err := r.db.First(&item, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *mysqlItemRepository) ListAll(offset int, limit int) ([]*entity.Item, int64, error) {
	var items []*entity.Item
	var totalCount int64

	if err := r.db.Model(&entity.Item{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset(offset).Limit(limit).Order("name ASC").Find(&items).Error; err != nil {
		return nil, totalCount, err
	}
	return items, totalCount, nil
}

func (r *mysqlItemRepository) ListByType(itemType string, offset int, limit int) ([]*entity.Item, int64, error) {
	var items []*entity.Item
	var totalCount int64

	query := r.db.Model(&entity.Item{}).Where("item_type = ?", itemType)

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("name ASC").Find(&items).Error; err != nil {
		return nil, totalCount, err
	}
	return items, totalCount, nil
}

func (r *mysqlItemRepository) Create(item *entity.Item) error {
	return r.db.Create(item).Error
}

func (r *mysqlItemRepository) Update(item *entity.Item) error {
	return r.db.Save(item).Error
}

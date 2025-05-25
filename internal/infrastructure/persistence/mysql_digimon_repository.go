package persistence

import (
	"go-gin/internal/domain/game/model/entity"
	"go-gin/internal/domain/game/repository" // Import the repository interfaces
	"gorm.io/gorm"
	// "gorm.io/gorm/clause" // May be needed for Preload if used
)

type mysqlDigimonRepository struct {
	db *gorm.DB
}

// NewMySQLDigimonRepository creates a new GORM-based DigimonRepository.
func NewMySQLDigimonRepository(db *gorm.DB) repository.DigimonRepository {
	return &mysqlDigimonRepository{db: db}
}

func (r *mysqlDigimonRepository) Create(digimon *entity.Digimon) error {
	return r.db.Create(digimon).Error
}

func (r *mysqlDigimonRepository) FindByID(id string) (*entity.Digimon, error) {
	var digimon entity.Digimon
	// Consider preloading Species if DigimonResponse always needs it,
	// though this can also be handled at the usecase level by fetching species separately.
	// Example: err := r.db.Preload("Species").First(&digimon, "id = ?", id).Error
	err := r.db.First(&digimon, "id = ?", id).Error
	if err != nil {
		return nil, err // GORM handles ErrRecordNotFound
	}
	return &digimon, nil
}

func (r *mysqlDigimonRepository) ListByPlayerID(playerID string, offset int, limit int) ([]*entity.Digimon, error) {
	var digimonList []*entity.Digimon
	// Add OrderBy if needed, e.g. .Order("acquired_at DESC")
	err := r.db.Where("player_id = ?", playerID).Offset(offset).Limit(limit).Find(&digimonList).Error
	if err != nil {
		return nil, err
	}
	return digimonList, nil
}

func (r *mysqlDigimonRepository) Update(digimon *entity.Digimon) error {
	// Ensure 'digimon' has its ID set for GORM to update correctly.
	// .Save() updates all fields or creates if primary key is zero.
	// .Updates() updates specific fields if used with a struct or map.
	return r.db.Save(digimon).Error
}

func (r *mysqlDigimonRepository) Delete(id string) error {
	// For soft delete, GORM needs a gorm.DeletedAt field in the entity and db.Delete()
	// For hard delete:
	return r.db.Delete(&entity.Digimon{}, "id = ?", id).Error
}

// Optional: Implement CountByPlayerID if needed by usecases (e.g. for checking Digimon slots)
// func (r *mysqlDigimonRepository) CountByPlayerID(playerID string) (int64, error) {
//    var count int64
//    err := r.db.Model(&entity.Digimon{}).Where("player_id = ?", playerID).Count(&count).Error
//    return count, err
// }

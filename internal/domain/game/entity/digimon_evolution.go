package entity

import "time"

// DigimonEvolution 表示數碼寶貝的進化資訊。
// 包含從哪個物種進化成哪個物種、進化條件（如等級或道具）等。
type DigimonEvolution struct {
	ID string `gorm:"primaryKey;type:char(36)" json:"id"`
	// 唯一識別碼，使用 UUID（char(36)），也可改用自增整數（int auto_increment）

	FromSpeciesID string `gorm:"type:char(36);index;not null" json:"from_species_id"`
	// 起始物種的 ID（例如亞古獸）

	ToSpeciesID string `gorm:"type:char(36);index;not null" json:"to_species_id"`
	// 進化後的物種 ID（例如暴龍獸）

	Method string `gorm:"size:50;not null" json:"method"`
	// 進化方式，例如："LevelUp"（升級）、"ItemUse"（使用道具）、"DNA"（合體）

	RequiredLevel int `gorm:"null" json:"required_level,omitempty"`
	// 所需等級（若非升級方式則可為 NULL）

	RequiredItemID string `gorm:"type:char(36);null" json:"required_item_id,omitempty"`
	// 所需道具 ID（如果進化需要特定道具）

	Notes string `gorm:"type:text" json:"notes,omitempty"`
	// 備註或額外描述（可選）

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 建立時間，自動由 GORM 設定

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 最後更新時間，自動由 GORM 在資料變更時更新

	// 下列為可選的 GORM 關聯設計（通常在需要預載關聯資料時使用）
	// FromSpecies *DigimonSpecies `gorm:"foreignKey:FromSpeciesID" json:"-"`
	// ToSpecies   *DigimonSpecies `gorm:"foreignKey:ToSpeciesID" json:"-"`
	// RequiredItem *Item          `gorm:"foreignKey:RequiredItemID" json:"-"`
}

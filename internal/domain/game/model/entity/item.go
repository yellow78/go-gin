package entity

import "time"

type Item struct {
	ID                   string    `gorm:"primaryKey;type:char(36)" json:"id"` // Or int auto_increment
	Name                 string    `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Description          string    `gorm:"type:text" json:"description,omitempty"`
	ItemType             string    `gorm:"size:50;not null" json:"item_type"` // e.g., "Consumable", "KeyItem", "EvolutionMaterial", "Equipment"
	SpriteURL            string    `gorm:"size:255" json:"sprite_url,omitempty"`
	MaxStack             int       `gorm:"not null;default:1" json:"max_stack"`
	IsUsableInBattle     bool      `gorm:"default:false" json:"is_usable_in_battle"`
	IsUsableOutsideBattle bool      `gorm:"default:true" json:"is_usable_outside_battle"`
	EffectDetails        string    `gorm:"type:text" json:"effect_details,omitempty"` // Could be JSON string for complex effects
	SellPrice            int       `gorm:"not null;default:0" json:"sell_price"` // Price if sold to shop
	BuyPrice             int       `gorm:"not null;default:0" json:"buy_price"`   // Price if bought from shop
	CreatedAt            time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

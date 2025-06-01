package entity

import "time"

// Item 代表遊戲中的道具物件，可為消耗品、裝備、進化素材等
type Item struct {
	ID string `gorm:"primaryKey;type:char(36)" json:"id"`
	// 道具的唯一識別碼（UUID 格式）

	Name string `gorm:"uniqueIndex;size:100;not null" json:"name"`
	// 道具名稱，需唯一，例如：「回復藥」、「火焰寶石」

	Description string `gorm:"type:text" json:"description,omitempty"`
	// 道具描述，可提供玩家說明用途或背景故事，非必填

	ItemType string `gorm:"size:50;not null" json:"item_type"`
	// 道具類型，例如："Consumable"（消耗品）、"KeyItem"（關鍵道具）、
	// "EvolutionMaterial"（進化素材）、"Equipment"（裝備）

	SpriteURL string `gorm:"size:255" json:"sprite_url,omitempty"`
	// 道具的圖示 URL，可用於前端顯示

	MaxStack int `gorm:"not null;default:1" json:"max_stack"`
	// 最大堆疊數量，例如回復藥可以堆到 99，裝備通常為 1

	IsUsableInBattle bool `gorm:"default:false" json:"is_usable_in_battle"`
	// 是否可在戰鬥中使用，例如：治療藥水為 true，裝備則為 false

	IsUsableOutsideBattle bool `gorm:"default:true" json:"is_usable_outside_battle"`
	// 是否可在戰鬥外使用，例如：任務道具不可使用，回復藥水可使用

	EffectDetails string `gorm:"type:text" json:"effect_details,omitempty"`
	// 效果細節（建議以 JSON 字串存儲），可儲存複雜效果設定如：
	// {"heal_hp": 50}、{"evolve_to": "Greymon"}、{"buff": {"attack": 2}}

	SellPrice int `gorm:"not null;default:0" json:"sell_price"`
	// 出售價格（給 NPC 商店時能賣的價格）

	BuyPrice int `gorm:"not null;default:0" json:"buy_price"`
	// 購買價格（NPC 商店購買時的價格）

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 建立時間，由 GORM 自動產生

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 最後更新時間，由 GORM 自動維護
}

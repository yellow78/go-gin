package entity

import "time"

// Digimon 代表玩家所擁有的數碼寶貝個體資料（Player 的擁有物）
type Digimon struct {
	ID string `gorm:"primaryKey;type:char(36)" json:"id"`
	// 數碼寶貝個體的唯一識別碼（UUID 格式）

	PlayerID string `gorm:"type:char(36);index;not null" json:"player_id"`
	// 所屬玩家的 ID，對應到 users 表的主鍵（外鍵）

	SpeciesID string `gorm:"type:char(36);index;not null" json:"species_id"`
	// 所屬物種的 ID，對應 digimon_species 表的主鍵（外鍵）

	Nickname string `gorm:"size:100" json:"nickname,omitempty"`
	// 數碼寶貝的暱稱，非必填；可讓玩家自訂名字

	CurrentLevel int `gorm:"not null;default:1" json:"current_level"`
	// 目前等級，預設為 1；可根據經驗值或進化條件調整

	CurrentAttack int `gorm:"not null" json:"current_attack"`
	// 目前攻擊力，可能依據基礎值 + 等級成長 + 道具影響

	CurrentDefense int `gorm:"not null" json:"current_defense"`
	// 目前防禦力

	CurrentSpeed int `gorm:"not null" json:"current_speed"`
	// 目前速度，用於判定出手順序等戰鬥邏輯

	ExperiencePoints int64 `gorm:"not null;default:0" json:"experience_points"`
	// 經驗值，使用 int64 支援更高上限

	AcquiredAt time.Time `gorm:"autoCreateTime" json:"acquired_at"`
	// 獲得時間，GORM 會在插入時自動填入

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 最後更新時間，GORM 會在更新時自動維護

	// 關聯資料（可選）：實務中通常由 service 層做 Join 或 DTO 聚合處理
	// User    *accountEntity.User `gorm:"foreignKey:PlayerID" json:"-"`         // 所屬玩家
	// Species *DigimonSpecies     `gorm:"foreignKey:SpeciesID" json:"-"`        // 對應的物種資料
}

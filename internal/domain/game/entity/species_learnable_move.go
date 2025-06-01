package entity

import "time" // 引入 time 用於 CreatedAt/UpdatedAt 欄位

// SpeciesLearnableMove 定義某個數碼寶貝種類（Species）可學會的技能（Move），
// 並可以指定學習方法（如升級、導師教學）以及需求條件（如等級）。
type SpeciesLearnableMove struct {
	ID string `gorm:"primaryKey;type:char(36)" json:"id"`
	// 該紀錄的唯一識別碼（UUID）

	SpeciesID string `gorm:"type:char(36);index:idx_species_learn_move_keys,unique;not null" json:"species_id"`
	// 數碼寶貝種類 ID，對應 digimon_species 表主鍵（外鍵）
	// 與 MoveID 一起構成唯一索引，確保同一物種不會重複綁定相同技能

	MoveID string `gorm:"type:char(36);index:idx_species_learn_move_keys,unique;not null" json:"move_id"`
	// 招式 ID，對應 digimon_moves 表主鍵（外鍵）

	LearnMethod string `gorm:"size:50;not null" json:"learn_method"`
	// 技能學習方式，例如 "LevelUp", "Tutor", "EggMove"

	RequiredLevel int `gorm:"null" json:"required_level,omitempty"`
	// 若為升級學習，則指定等級。其他學習方式可為 NULL

	Notes string `gorm:"type:text" json:"notes,omitempty"`
	// 額外註解或條件，例如「需特定道具」或「僅限夜間進化」

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 建立時間（自動生成）

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 最後更新時間（自動生成）

	// 可選的 GORM 關聯，如果需要 preload 可啟用以下欄位：
	// Species DigimonSpecies `gorm:"foreignKey:SpeciesID" json:"-"`
	// Move    DigimonMove    `gorm:"foreignKey:MoveID" json:"-"`
}

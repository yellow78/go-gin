package entity

import "time"

// PlayerDigimonKnownMove 用於連結玩家擁有的某隻數碼寶貝實例與該實例已學會的招式。
// 每筆紀錄代表一隻數碼寶貝學會了一個招式。
type PlayerDigimonKnownMove struct {
	ID string `gorm:"primaryKey;type:char(36)" json:"id"`
	// 此 join 紀錄的唯一識別碼（UUID）

	PlayerDigimonID string `gorm:"type:char(36);index:idx_pdigimon_known_move_keys,unique;not null" json:"player_digimon_id"`
	// 玩家擁有的數碼寶貝實例 ID，對應 player_digimon 表的主鍵（外鍵）
	// 並與 MoveID 一起構成唯一索引，避免同一隻數碼寶貝重複學同一招式

	MoveID string `gorm:"type:char(36);index:idx_pdigimon_known_move_keys,unique;not null" json:"move_id"`
	// 招式 ID，對應 digimon_moves 表的主鍵（外鍵）

	SlotPosition int `gorm:"null" json:"slot_position,omitempty"`
	// 技能欄位位置，例如 1~4，如果該數碼寶貝僅能持有有限的技能欄
	// 如不使用此功能，可留空（NULL）

	LearnedAt time.Time `gorm:"autoCreateTime" json:"learned_at"`
	// 該招式由此數碼寶貝實例學會的時間（自動填入）

	// 以下為可選的 GORM 關聯，如果需要預載關聯資料可取消註解並使用 Preload
	// PlayerDigimon Digimon    `gorm:"foreignKey:PlayerDigimonID" json:"-"`
	// Move          DigimonMove `gorm:"foreignKey:MoveID" json:"-"`
}

package entity

import "time" // 時間欄位需要用到 time 套件

// DigimonMove 表示數碼寶貝可以學會的招式資料
type DigimonMove struct {
	ID string `gorm:"primaryKey;type:char(36)" json:"id"`
	// 招式的唯一識別碼，使用 UUID 格式（char(36)）

	Name string `gorm:"uniqueIndex;size:100;not null" json:"name"`
	// 招式名稱，必填且唯一，例如「火焰箭」

	Description string `gorm:"type:text" json:"description,omitempty"`
	// 招式說明，可為空，例如「對敵人造成火屬性傷害，有機率灼燒」

	Power int `gorm:"not null;default:0" json:"power"`
	// 攻擊力數值，預設為 0，非攻擊技能可設為 0

	Accuracy int `gorm:"not null;default:100" json:"accuracy"`
	// 命中率，範圍為 0~100，預設 100 表示必中

	MoveType string `gorm:"size:50;not null" json:"move_type"`
	// 招式類型，例如：「Fire（火）」、「Water（水）」、「Status（狀態技）」等

	MPCost int `gorm:"not null;default:0" json:"mp_cost"`
	// 消耗的魔力（MP），預設為 0

	Target string `gorm:"size:50;not null" json:"target"`
	// 技能目標，例如：「SingleEnemy（單體敵人）」、「AllEnemies（全體敵人）」、「Self（自身）」

	BaseCritChance int `gorm:"not null;default:5" json:"base_crit_chance"`
	// 暴擊機率（0~100），預設為 5%

	Effect string `gorm:"size:255" json:"effect,omitempty"`
	// 附加效果，例如：「Burn:10%（10%機率灼燒）」、「Heal:25%（恢復25% HP）」、「StatBoost:Attack:Self:1_stage（提升自身攻擊一階）」等

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 建立時間，自動生成

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 最後更新時間，自動更新
}

package entity

import "time"

// DigimonSpecies 表示數碼寶貝的物種資料，例如亞古獸、加布獸等。
type DigimonSpecies struct {
	ID string `gorm:"primaryKey;type:char(36)" json:"id"`
	// 數碼寶貝物種的唯一識別碼，通常使用 UUID 格式（長度 36）

	Name string `gorm:"uniqueIndex;size:100;not null" json:"name"`
	// 數碼寶貝名稱，必填且唯一，例如：「亞古獸」、「天女獸」

	Attribute string `gorm:"size:20;not null" json:"attribute"`
	// 屬性，例如：「Vaccine（疫苗）」、「Virus（病毒）」、「Data（資料）」等

	Stage string `gorm:"size:20;not null" json:"stage"`
	// 階段，例如：「Rookie（成長期）」、「Champion（成熟期）」、「Ultimate（完全體）」等

	BaseAttack int `gorm:"not null" json:"base_attack"`
	// 初始攻擊力數值，用於戰鬥或進化判斷

	BaseDefense int `gorm:"not null" json:"base_defense"`
	// 初始防禦力數值

	BaseSpeed int `gorm:"not null" json:"base_speed"`
	// 初始速度數值，影響出手順序或逃跑成功率

	SpriteURL string `gorm:"size:255" json:"sprite_url,omitempty"`
	// 數碼寶貝的圖像或頭像網址，可選填（可用於 UI 顯示）

	Description string `gorm:"type:text" json:"description,omitempty"`
	// 數碼寶貝描述或設定介紹，可選填

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// 資料建立時間，GORM 自動填入

	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	// 資料最後更新時間，GORM 自動更新

	// 後續可擴充：EvolutionInfo（可進化成哪些種）或被哪些 Digimon 當作進化條件
}

package game

// AcquireDigimonRequest 用於玩家取得新數碼寶貝時的請求。
// 玩家需指定物種 ID，可選擇給予暱稱。
type AcquireDigimonRequest struct {
	SpeciesID string `json:"species_id" binding:"required,uuid"` // 數碼寶貝的物種 ID（UUID 格式，必填）
	Nickname  string `json:"nickname,omitempty"`                 // 初始暱稱（可選）
}

// UpdateDigimonNicknameRequest 用於變更數碼寶貝的暱稱。
type UpdateDigimonNicknameRequest struct {
	Nickname string `json:"nickname" binding:"required,max=100"` // 新暱稱（必填，最多 100 字元）
}

// DigimonSpeciesResponse 表示單一數碼寶貝物種的詳細資訊。
type DigimonSpeciesResponse struct {
	ID          string `json:"id"`                    // 物種 ID
	Name        string `json:"name"`                  // 名稱
	Attribute   string `json:"attribute"`             // 屬性（如疫苗種、病毒種等）
	Stage       string `json:"stage"`                 // 成長階段（如幼年期、完全體等）
	BaseAttack  int    `json:"base_attack"`           // 基礎攻擊力
	BaseDefense int    `json:"base_defense"`          // 基礎防禦力
	BaseSpeed   int    `json:"base_speed"`            // 基礎速度
	SpriteURL   string `json:"sprite_url,omitempty"`  // 角色圖像網址（可選）
	Description string `json:"description,omitempty"` // 物種描述（可選）
}

// DigimonResponse 表示玩家所擁有的數碼寶貝資訊（包含實例狀態與物種資料）。
type DigimonResponse struct {
	ID               string                 `json:"id"`                 // 數碼寶貝實例 ID
	PlayerID         string                 `json:"player_id"`          // 擁有者（玩家）ID
	Nickname         string                 `json:"nickname,omitempty"` // 暱稱（可選）
	CurrentLevel     int                    `json:"current_level"`      // 當前等級
	CurrentAttack    int                    `json:"current_attack"`     // 當前攻擊力
	CurrentDefense   int                    `json:"current_defense"`    // 當前防禦力
	CurrentSpeed     int                    `json:"current_speed"`      // 當前速度
	ExperiencePoints int64                  `json:"experience_points"`  // 經驗值
	AcquiredAt       string                 `json:"acquired_at"`        // 獲得時間（RFC3339 格式的時間字串）
	Species          DigimonSpeciesResponse `json:"species"`            // 所屬物種的詳細資料
}

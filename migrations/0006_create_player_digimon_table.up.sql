-- 建立 player_digimon（玩家擁有的數碼獸實體）資料表
CREATE TABLE IF NOT EXISTS player_digimon (
    -- 唯一識別碼（UUID）
    id CHAR(36) PRIMARY KEY COMMENT 'UUID',
    -- 玩家 ID，對應 users 表
    player_id CHAR(36) NOT NULL COMMENT '玩家ID',
    -- 數碼獸物種 ID，對應 digimon_species 表
    species_id CHAR(36) NOT NULL COMMENT '數碼獸ID',
    -- 數碼獸名稱（可選，玩家自定義）
    nickname VARCHAR(100) COMMENT '數碼獸暱稱',
    -- 目前等級，預設為 1
    current_level INT NOT NULL DEFAULT 1 COMMENT '等級',
    -- 目前攻擊力（根據等級和成長計算）
    current_attack INT NOT NULL COMMENT '目前攻擊力',
    current_defense INT NOT NULL COMMENT '目前防禦力',
    current_speed INT NOT NULL COMMENT '目前速度',
    -- 目前經驗值，決定是否升級
    experience_points BIGINT NOT NULL DEFAULT 0 COMMENT '目前經驗值',
    -- 取得時間（加入隊伍時間）
    acquired_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '取得時間',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最後更新時間',

    -- 索引提高查詢效率
    INDEX idx_player_digimon_player_id (player_id),
    INDEX idx_player_digimon_species_id (species_id),

    -- 關聯約束
    FOREIGN KEY (player_id) REFERENCES users(id) ON DELETE CASCADE COMMENT '關聯玩家帳號',
    FOREIGN KEY (species_id) REFERENCES digimon_species(id) ON DELETE RESTRICT COMMENT '關聯數碼獸物種'
);

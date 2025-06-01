-- 建立 player_digimon_known_moves（玩家數碼獸已學會技能）資料表
CREATE TABLE IF NOT EXISTS player_digimon_known_moves (
    -- 唯一識別碼（UUID）
    id CHAR(36) PRIMARY KEY COMMENT 'UUID',
    -- 玩家數碼獸 ID（對應 player_digimon）
    player_digimon_id CHAR(36) NOT NULL COMMENT '數碼獸ID',
    -- 技能 ID（對應 digimon_moves）
    move_id CHAR(36) NOT NULL COMMENT '技能ID',
    -- 技能欄位位置（技能欄位序號，可為 NULL）
    slot_position INT NULL COMMENT '技能欄位位置',
    -- 技能學會時間
    learned_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '技能學會時間',
    
    UNIQUE INDEX uq_idx_pdigimon_known_move_keys (player_digimon_id, move_id),
    
    FOREIGN KEY (player_digimon_id) REFERENCES player_digimon(id) ON DELETE CASCADE COMMENT '關聯玩家數碼獸',
    FOREIGN KEY (move_id) REFERENCES digimon_moves(id) ON DELETE CASCADE COMMENT '關聯技能'
);

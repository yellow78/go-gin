-- 建立 species_learnable_moves（數碼獸物種可學習技能）資料表
CREATE TABLE IF NOT EXISTS species_learnable_moves (
    -- 唯一識別碼（UUID）
    id CHAR(36) PRIMARY KEY COMMENT 'UUID',
    -- 數碼獸物種 ID（對應 digimon_species）
    species_id CHAR(36) NOT NULL COMMENT '數碼獸ID',
    -- 技能 ID（對應 digimon_moves）
    move_id CHAR(36) NOT NULL COMMENT '技能ID',
    -- 學習方式（如：升級、自學、道具等）
    learn_method VARCHAR(50) NOT NULL COMMENT '學習方式',
    -- 學習此技能所需等級（若有）
    required_level INT NULL COMMENT '學習所需等級',
    -- 備註說明（額外條件等）
    notes TEXT COMMENT '備註說明',
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
    
    UNIQUE INDEX uq_idx_species_learn_move_keys (species_id, move_id),

    FOREIGN KEY (species_id) REFERENCES digimon_species(id) ON DELETE CASCADE COMMENT '關聯數碼獸物種',
    FOREIGN KEY (move_id) REFERENCES digimon_moves(id) ON DELETE CASCADE COMMENT '關聯技能'
);

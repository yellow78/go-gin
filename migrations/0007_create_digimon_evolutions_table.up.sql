-- 建立 digimon_evolutions（數碼獸進化路徑）資料表
CREATE TABLE IF NOT EXISTS digimon_evolutions (
    -- 唯一識別碼（UUID）
    id CHAR(36) PRIMARY KEY COMMENT 'UUID',
    -- 起始數碼獸物種 ID（來源）
    from_species_id CHAR(36) NOT NULL COMMENT '起始數碼獸物種ID',
    -- 進化後數碼獸物種 ID（目標）
    to_species_id CHAR(36) NOT NULL COMMENT '進化後數碼獸物種ID',
    -- 進化方式（如：等級、道具、任務）
    method VARCHAR(50) NOT NULL COMMENT '進化方式）',
    -- 需要達到的等級（若有）
    required_level INT NULL COMMENT '需要達到的等級',
    -- 需要消耗的道具 ID（若有）
    required_item_id CHAR(36) NULL COMMENT '需要消耗的道具ID',
    -- 備註說明（額外條件或說明）
    notes TEXT COMMENT '備註說明（額外條件或說明）',

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',

    -- 索引以加速查詢
    INDEX idx_evolutions_from_species_id (from_species_id),
    INDEX idx_evolutions_to_species_id (to_species_id),

    -- 外鍵約束
    FOREIGN KEY (from_species_id) REFERENCES digimon_species(id) ON DELETE CASCADE COMMENT '來源物種關聯',
    FOREIGN KEY (to_species_id) REFERENCES digimon_species(id) ON DELETE CASCADE COMMENT '目標物種關聯',
    FOREIGN KEY (required_item_id) REFERENCES items(id) ON DELETE SET NULL COMMENT '所需道具關聯'
);

-- 建立 digimon_moves（數碼獸技能）資料表
CREATE TABLE IF NOT EXISTS digimon_moves (
    -- 唯一識別碼（UUID）
    id CHAR(36) PRIMARY KEY COMMENT 'UUID',
    -- 技能名稱（唯一）
    name VARCHAR(100) NOT NULL UNIQUE COMMENT '技能名稱',
    description TEXT COMMENT '技能描述',
    -- 技能威力（數值越高傷害越高）
    power INT NOT NULL DEFAULT 0 COMMENT '技能威力',
    -- 命中率（百分比）
    accuracy INT NOT NULL DEFAULT 100 COMMENT '命中率',
    -- 技能類型（如：火、水、光、暗、物理、魔法等）
    move_type VARCHAR(50) NOT NULL COMMENT '技能類型',
    -- 魔力消耗（MP）
    mp_cost INT NOT NULL DEFAULT 0 COMMENT '魔力消耗',
    -- 目標類型（如：單體、全體、己方、敵方等）
    target VARCHAR(50) NOT NULL COMMENT '目標類型',
    -- 基礎暴擊率（百分比，預設 5%）
    base_crit_chance INT NOT NULL DEFAULT 5 COMMENT '基礎暴擊率',
    -- 附加效果（如：中毒、暈眩、恢復）
    effect VARCHAR(255) COMMENT '附加效果',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間'
);

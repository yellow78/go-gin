-- 建立 digimon_species（數碼獸物種）資料表
CREATE TABLE IF NOT EXISTS digimon_species (
    -- 唯一識別碼（UUID）
    id CHAR(36) PRIMARY KEY COMMENT 'UUID',
    -- 數碼獸名稱（唯一）
    name VARCHAR(100) NOT NULL UNIQUE COMMENT '數碼獸名稱',
    -- 屬性（如：疫苗、資料、病毒）
    attribute VARCHAR(20) NOT NULL COMMENT '屬性',
    -- 進化階段（如：幼年期、成熟期、完全體等）
    stage VARCHAR(20) NOT NULL COMMENT '進化階段',
    base_attack INT NOT NULL COMMENT '基礎攻擊力',
    base_defense INT NOT NULL COMMENT '基礎防禦力',
    base_speed INT NOT NULL COMMENT '基礎速度',
    -- 數碼獸圖示的 URL
    sprite_url VARCHAR(255) COMMENT '圖片 URL',
    description TEXT COMMENT '描述或介紹',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
    -- 最後更新時間（自動更新）
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最後更新時間'
);

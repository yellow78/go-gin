-- 建立 items（道具）資料表
CREATE TABLE IF NOT EXISTS items (
    -- 唯一識別碼（UUID）
    id CHAR(36) PRIMARY KEY COMMENT 'UUID',
    -- 道具名稱（唯一）
    name VARCHAR(100) NOT NULL UNIQUE COMMENT '道具名稱',
    description TEXT COMMENT '道具描述說明',
    -- 道具類型（例如：恢復、裝備、進化材料）
    item_type VARCHAR(50) NOT NULL COMMENT '道具類型',
    -- 圖片網址（顯示用圖示）
    sprite_url VARCHAR(255) COMMENT '圖片網址',
    -- 最大堆疊數量（預設為 1）
    max_stack INT NOT NULL DEFAULT 1 COMMENT '最大堆疊數量',
    -- 是否可在戰鬥中使用（預設不可）
    is_usable_in_battle BOOLEAN DEFAULT FALSE COMMENT '戰鬥中是否可使用',
    -- 是否可在非戰鬥中使用（預設可）
    is_usable_outside_battle BOOLEAN DEFAULT TRUE COMMENT '非戰鬥中是否可使用',
    -- 效果說明（如恢復量、加成屬性等）
    effect_details TEXT COMMENT '效果說明',
    sell_price INT NOT NULL DEFAULT 0 COMMENT '販售價格',
    buy_price INT NOT NULL DEFAULT 0 COMMENT '購買價格',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '建立時間',
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間'
);

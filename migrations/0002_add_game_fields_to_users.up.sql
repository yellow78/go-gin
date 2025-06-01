-- 向 users 表添加與遊戲相關的欄位
ALTER TABLE users
  -- 遊戲暱稱（不可為空），後續加上唯一限制
  ADD COLUMN in_game_name VARCHAR(50) NOT NULL COMMENT '暱稱' AFTER email,
  -- 玩家等級（預設為 1）
  ADD COLUMN level INT NOT NULL DEFAULT 1 COMMENT '玩家等級' AFTER in_game_name,
  -- 經驗值（預設為 0）
  ADD COLUMN experience_points BIGINT NOT NULL DEFAULT 0 COMMENT '經驗值' AFTER level,
  -- 最後登入時間（可為 NULL）
  ADD COLUMN last_login_at TIMESTAMP NULL DEFAULT NULL COMMENT '登入時間' AFTER experience_points,
  -- 更新時間（每次資料更新自動更新此欄位）
  ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
    ON UPDATE CURRENT_TIMESTAMP COMMENT '資料更新時間，自動更新' AFTER created_at,
  -- 為 in_game_name 加上唯一約束，避免重複名稱
  ADD CONSTRAINT uq_in_game_name UNIQUE (in_game_name);

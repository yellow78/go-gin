-- Add game-specific fields to users table
ALTER TABLE users
ADD COLUMN in_game_name VARCHAR(50) NOT NULL AFTER email,
ADD COLUMN level INT NOT NULL DEFAULT 1 AFTER in_game_name,
ADD COLUMN experience_points BIGINT NOT NULL DEFAULT 0 AFTER level,
ADD COLUMN last_login_at TIMESTAMP NULL DEFAULT NULL AFTER experience_points,
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP AFTER created_at;

ALTER TABLE users
ADD CONSTRAINT uq_in_game_name UNIQUE (in_game_name);

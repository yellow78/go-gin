CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,  -- 使用 CHAR(36) 儲存 UUID
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

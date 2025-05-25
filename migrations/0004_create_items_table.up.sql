-- Create items table
CREATE TABLE IF NOT EXISTS items (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    item_type VARCHAR(50) NOT NULL,
    sprite_url VARCHAR(255),
    max_stack INT NOT NULL DEFAULT 1,
    is_usable_in_battle BOOLEAN DEFAULT FALSE,
    is_usable_outside_battle BOOLEAN DEFAULT TRUE,
    effect_details TEXT,
    sell_price INT NOT NULL DEFAULT 0,
    buy_price INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

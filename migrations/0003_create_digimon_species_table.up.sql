-- Create digimon_species table
CREATE TABLE IF NOT EXISTS digimon_species (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    attribute VARCHAR(20) NOT NULL,
    stage VARCHAR(20) NOT NULL,
    base_attack INT NOT NULL,
    base_defense INT NOT NULL,
    base_speed INT NOT NULL,
    sprite_url VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

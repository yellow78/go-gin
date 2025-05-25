-- Create player_digimon table
CREATE TABLE IF NOT EXISTS player_digimon (
    id CHAR(36) PRIMARY KEY,
    player_id CHAR(36) NOT NULL,
    species_id CHAR(36) NOT NULL,
    nickname VARCHAR(100),
    current_level INT NOT NULL DEFAULT 1,
    current_attack INT NOT NULL,
    current_defense INT NOT NULL,
    current_speed INT NOT NULL,
    experience_points BIGINT NOT NULL DEFAULT 0,
    acquired_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_player_digimon_player_id (player_id),
    INDEX idx_player_digimon_species_id (species_id),
    FOREIGN KEY (player_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (species_id) REFERENCES digimon_species(id) ON DELETE RESTRICT
);

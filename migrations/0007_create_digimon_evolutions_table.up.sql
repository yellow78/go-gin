-- Create digimon_evolutions table
CREATE TABLE IF NOT EXISTS digimon_evolutions (
    id CHAR(36) PRIMARY KEY,
    from_species_id CHAR(36) NOT NULL,
    to_species_id CHAR(36) NOT NULL,
    method VARCHAR(50) NOT NULL,
    required_level INT NULL,
    required_item_id CHAR(36) NULL,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_evolutions_from_species_id (from_species_id),
    INDEX idx_evolutions_to_species_id (to_species_id),
    FOREIGN KEY (from_species_id) REFERENCES digimon_species(id) ON DELETE CASCADE,
    FOREIGN KEY (to_species_id) REFERENCES digimon_species(id) ON DELETE CASCADE,
    FOREIGN KEY (required_item_id) REFERENCES items(id) ON DELETE SET NULL
);

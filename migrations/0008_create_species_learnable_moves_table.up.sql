-- Create species_learnable_moves table
CREATE TABLE IF NOT EXISTS species_learnable_moves (
    id CHAR(36) PRIMARY KEY, 
    species_id CHAR(36) NOT NULL,
    move_id CHAR(36) NOT NULL,
    learn_method VARCHAR(50) NOT NULL,
    required_level INT NULL,
    notes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE INDEX uq_idx_species_learn_move_keys (species_id, move_id),
    FOREIGN KEY (species_id) REFERENCES digimon_species(id) ON DELETE CASCADE,
    FOREIGN KEY (move_id) REFERENCES digimon_moves(id) ON DELETE CASCADE
);

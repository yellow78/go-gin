-- Create player_digimon_known_moves table
CREATE TABLE IF NOT EXISTS player_digimon_known_moves (
    id CHAR(36) PRIMARY KEY,
    player_digimon_id CHAR(36) NOT NULL,
    move_id CHAR(36) NOT NULL,
    slot_position INT NULL,
    learned_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE INDEX uq_idx_pdigimon_known_move_keys (player_digimon_id, move_id),
    FOREIGN KEY (player_digimon_id) REFERENCES player_digimon(id) ON DELETE CASCADE,
    FOREIGN KEY (move_id) REFERENCES digimon_moves(id) ON DELETE CASCADE
);

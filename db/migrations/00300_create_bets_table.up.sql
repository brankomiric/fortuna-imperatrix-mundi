BEGIN;

CREATE TABLE bets (
    player_id INT NOT NULL,
    tournament_id INT NOT NULL,
    amount INT NOT NULL CHECK (amount > 0),
    PRIMARY KEY (player_id, tournament_id),
    FOREIGN KEY (player_id) REFERENCES players(player_id) ON DELETE CASCADE,
    FOREIGN KEY (tournament_id) REFERENCES tournaments(tournament_id) ON DELETE CASCADE
);

COMMIT;
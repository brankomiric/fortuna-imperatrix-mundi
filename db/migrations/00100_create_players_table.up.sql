BEGIN;

CREATE TABLE players (
    player_id INT AUTO_INCREMENT PRIMARY KEY,
    player_name VARCHAR(100) NOT NULL,
    player_email VARCHAR(255) UNIQUE NOT NULL,
    account_balance INT DEFAULT 0
);

COMMIT;
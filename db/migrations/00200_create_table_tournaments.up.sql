BEGIN;

CREATE TABLE tournaments (
    tournament_id INT AUTO_INCREMENT PRIMARY KEY,
    tournament_name VARCHAR(150) NOT NULL,
    prize_pool INT DEFAULT 0,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);

COMMIT;
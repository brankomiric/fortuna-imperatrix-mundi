BEGIN;

CREATE PROCEDURE DistributePrizes(IN tournamentId INT)
BEGIN
    DECLARE total_prize INT;
    DECLARE first_place_player INT;
    DECLARE second_place_player INT;
    DECLARE third_place_player INT;
    DECLARE first_place_prize INT;
    DECLARE second_place_prize INT;
    DECLARE third_place_prize INT;
    
    prize_check: BEGIN
        -- Get the total prize pool for the tournament
        SELECT prize_pool INTO total_prize FROM tournaments WHERE tournament_id = tournamentId;
        
        -- Exit if total_prize is 0 or NULL
        IF total_prize IS NULL OR total_prize = 0 THEN
            LEAVE prize_check;
        END IF;
    
        -- Get the top 3 players based on highest bet, breaking ties by earliest bet_time
        SELECT player_id INTO first_place_player FROM bets 
        WHERE tournament_id = tournamentId 
        ORDER BY amount DESC, bet_time ASC 
        LIMIT 1;
        
        SELECT player_id INTO second_place_player FROM bets 
        WHERE tournament_id = tournamentId 
        ORDER BY amount DESC, bet_time ASC 
        LIMIT 1 OFFSET 1;
    
        SELECT player_id INTO third_place_player FROM bets 
        WHERE tournament_id = tournamentId 
        ORDER BY amount DESC, bet_time ASC 
        LIMIT 1 OFFSET 2;
    
        -- Calculate prize amounts
        SET first_place_prize = total_prize * 0.50;
        SET second_place_prize = total_prize * 0.30;
        SET third_place_prize = total_prize * 0.20;
    
        -- Update the account balances of the top 3 players
        IF first_place_player IS NOT NULL THEN
            UPDATE players SET account_balance = account_balance + first_place_prize 
            WHERE player_id = first_place_player;
        END IF;
        
        IF second_place_player IS NOT NULL THEN
            UPDATE players SET account_balance = account_balance + second_place_prize 
            WHERE player_id = second_place_player;
        END IF;
        
        IF third_place_player IS NOT NULL THEN
            UPDATE players SET account_balance = account_balance + third_place_prize 
            WHERE player_id = third_place_player;
        END IF;

        -- Update the prize pool to 0
        UPDATE tournaments set prize_pool = 0 WHERE tournament_id = tournamentId;
    
    END prize_check;
    
END;

COMMIT;
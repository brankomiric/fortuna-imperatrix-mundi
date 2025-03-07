package database

import (
	"fmt"

	"github.com/brankomiric/fortuna-imperatrix-mundi/internal/dto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

func Initialize(connectionStr string) (*Database, error) {
	db, err := connect(connectionStr)
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}

func connect(connectionStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", connectionStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateConnectionString(host string, port string, user string, password string, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
}

func (dbObj *Database) TestConn() error {
	err := dbObj.DB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func (dbObj *Database) AddTournament(input dto.CreateTournament) (*int64, error) {
	query := "INSERT INTO tournaments (tournament_name, prize_pool, start_date, end_date) VALUES (?, ?, ?, ?)"
	result, err := dbObj.DB.Exec(query, input.Name, 0, input.StartDate, input.EndDate)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &lastInsertID, nil
}

func (dbObj *Database) PlaceBet(input dto.Bet) error {
	tx, err := dbObj.DB.Beginx()
	if err != nil {
		return err
	}

	// Deduct the amount from the player's account balance
	updateBalanceQuery := "UPDATE players SET account_balance = account_balance - ? WHERE player_id = ? AND account_balance >= ?"
	updateBalanceResult, err := tx.Exec(updateBalanceQuery, input.Amount, input.PlayerID, input.Amount)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	rowsAffected, err := updateBalanceResult.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("player does not have enough balance")
	}

	// Create the bet
	betQuery := "INSERT INTO bets (player_id, tournament_id, amount) VALUES (?, ?, ?)"
	_, err = tx.Exec(betQuery, input.PlayerID, input.TournamentID, input.Amount)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	// Update the tournament prize pool
	updateTournamentPrizePoolQuery := "UPDATE tournaments SET prize_pool = prize_pool + ? WHERE tournament_id = ?"
	_, err = tx.Exec(updateTournamentPrizePoolQuery, input.Amount, input.TournamentID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}

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

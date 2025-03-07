package database

import (
	"fmt"
	"os"
)

type DBConnParams struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func ReadConnectionStringParams() (*DBConnParams, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing DB connection parameters")

	}
	return &DBConnParams{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
	}, nil
}

#!/bin/sh

# This script is used to run migrations on the database.
# Example usage: ./runMigrations.sh "mysql://user:password@tcp(host:port)/dbname"

BD_CONN_STR=$1

go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

migrate -database $BD_CONN_STR -path db/migrations up
package database

import "github.com/brankomiric/fortuna-imperatrix-mundi/internal/dto"

type AbstractDB interface {
	TestConn() error
	AddTournament(input dto.CreateTournament) (*int64, error)
	PlaceBet(input dto.Bet) error
	InvokeDistributePrizesProcedure(tournamentID int) error
	GetPlayersRankedByBalance() ([]Player, error)
}

package dto

import "time"

type CreateTournament struct {
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Bet struct {
	PlayerID     int64 `json:"player_id"`
	TournamentID int64 `json:"tournament_id"`
	Amount       int64 `json:"amount"`
}

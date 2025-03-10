package database

type Player struct {
	ID             int64  `json:"id" db:"player_id"`
	Name           string `json:"name" db:"player_name"`
	Email          string `json:"email" db:"player_email"`
	AccountBalance int    `json:"account_balance" db:"account_balance"`
}

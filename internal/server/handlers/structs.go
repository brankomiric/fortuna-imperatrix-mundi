package handlers

import "github.com/brankomiric/fortuna-imperatrix-mundi/internal/database"

type Handler struct {
	DB database.AbstractDB
}

type HealthcheckResponse struct {
	Services Services `json:"services"`
}

type Services struct {
	MySQL HealthcheckMessage `json:"MySQL"`
}

type HealthcheckMessage struct {
	Status string `json:"status"`
}

type Response struct {
	Message string `json:"message"`
}

type IDResponse struct {
	ID int64 `json:"id"`
}

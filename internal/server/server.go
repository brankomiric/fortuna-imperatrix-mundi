package server

import (
	"github.com/brankomiric/fortuna-imperatrix-mundi/internal/database"
	"github.com/brankomiric/fortuna-imperatrix-mundi/internal/server/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRouter(db *database.Database) *fiber.App {
	app := fiber.New()

	h := handlers.Handler{DB: db}

	app.Get("/health", h.Health)

	app.Post("/tournaments/create", h.CreateTournament)

	app.Post("/players/bet", h.PlaceBet)

	app.Post("/tournaments/prizes/distribute/:tournament_id", h.DistributePrizes)

	return app
}

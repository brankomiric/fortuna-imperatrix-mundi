package server

import (
	"github.com/brankomiric/fortuna-imperatrix-mundi/internal/server/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	h := handlers.Handler{}

	app.Get("/health", h.Health)

	return app
}

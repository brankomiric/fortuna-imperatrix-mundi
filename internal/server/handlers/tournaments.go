package handlers

import (
	"net/http"

	"github.com/brankomiric/fortuna-imperatrix-mundi/internal/dto"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) CreateTournament(c fiber.Ctx) error {
	body := dto.CreateTournament{}

	if err := c.Bind().Body(&body); err != nil {
		errResp := ErrorResponse{Message: err.Error()}
		return c.Status(http.StatusBadRequest).JSON(errResp)
	}

	id, err := h.DB.AddTournament(body)
	if err != nil {
		errResp := ErrorResponse{Message: err.Error()}
		return c.Status(http.StatusInternalServerError).JSON(errResp)
	}

	return c.Status(http.StatusOK).JSON(IDResponse{ID: *id})
}

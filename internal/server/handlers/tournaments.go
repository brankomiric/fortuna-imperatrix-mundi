package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/brankomiric/fortuna-imperatrix-mundi/internal/dto"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) CreateTournament(c fiber.Ctx) error {
	body := dto.CreateTournament{}

	if err := c.Bind().Body(&body); err != nil {
		errResp := Response{Message: err.Error()}
		return c.Status(http.StatusBadRequest).JSON(errResp)
	}

	id, err := h.DB.AddTournament(body)
	if err != nil {
		errResp := Response{Message: err.Error()}
		return c.Status(http.StatusInternalServerError).JSON(errResp)
	}

	return c.Status(http.StatusOK).JSON(IDResponse{ID: *id})
}

func (h *Handler) PlaceBet(c fiber.Ctx) error {
	body := dto.Bet{}

	if err := c.Bind().Body(&body); err != nil {
		errResp := Response{Message: err.Error()}
		return c.Status(http.StatusBadRequest).JSON(errResp)
	}

	err := h.DB.PlaceBet(body)
	if err != nil {
		errResp := Response{Message: err.Error()}
		return c.Status(http.StatusInternalServerError).JSON(errResp)
	}

	return c.Status(http.StatusOK).JSON(Response{Message: "Bet placed"})
}

func (h *Handler) DistributePrizes(c fiber.Ctx) error {
	idStr := c.Params("tournament_id")

	log.Println("Distributing prizes for tournament number:", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errResp := Response{Message: err.Error()}
		return c.Status(http.StatusBadRequest).JSON(errResp)
	}

	err = h.DB.InvokeDistributePrizesProcedure(id)
	if err != nil {
		errResp := Response{Message: err.Error()}
		return c.Status(http.StatusInternalServerError).JSON(errResp)
	}

	return c.Status(http.StatusOK).JSON(Response{Message: "Prizes distributed"})
}

func (h *Handler) RankedPlayers(c fiber.Ctx) error {
	players, err := h.DB.GetPlayersRankedByBalance()
	if err != nil {
		errResp := Response{Message: err.Error()}
		return c.Status(http.StatusInternalServerError).JSON(errResp)
	}
	return c.Status(http.StatusOK).JSON(players)
}

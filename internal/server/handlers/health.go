package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func (h *Handler) Health(c fiber.Ctx) error {
	httpStatus := http.StatusOK
	resp := HealthcheckResponse{
		Services{
			MySQL: HealthcheckMessage{Status: "OK"},
		},
	}

	err := h.DB.TestConn()
	if err != nil {
		log.Printf("health check failed: MySQL: %s\n", err.Error())

		resp.Services.MySQL.Status = "Not OK"
		httpStatus = http.StatusInternalServerError
	}

	return c.Status(httpStatus).JSON(resp)
}

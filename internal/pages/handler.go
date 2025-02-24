package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router fiber.Router
	logger *slog.Logger
}

func NewHandler(router fiber.Router, logger *slog.Logger) {
	h := &PagesHandler{
		router: router,
		logger: logger,
	}
	api := h.router.Group("/api")
	api.Get("/error", h.error)
}

func (h *PagesHandler) error(c *fiber.Ctx) error {
	h.logger.Error("custom error message")
	return c.SendString("error, World!")
}

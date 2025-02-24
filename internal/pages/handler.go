package pages

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type PagesHandler struct {
	router fiber.Router
	logger *slog.Logger
}

type Categories struct {
	Categories []Category
}

type Category struct {
	Name string
	Link string
}

func NewHandler(router fiber.Router, logger *slog.Logger) {
	h := &PagesHandler{
		router: router,
		logger: logger,
	}
	api := h.router.Group("/api")
	h.router.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *PagesHandler) error(c *fiber.Ctx) error {
	h.logger.Error("custom error message")
	return c.SendString("error, World!")
}

func (h *PagesHandler) home(c *fiber.Ctx) error {
	data := Categories{
		Categories: []Category{
			{Name: "Category 1", Link: "/category/1"},
			{Name: "Category 2", Link: "/category/2"},
			{Name: "Category 3", Link: "/category/3"},
		},
	}
	return c.Render("home", data)
}

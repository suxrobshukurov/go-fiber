package main

import (
	"suxrobshukurov/go-fiber/config"
	"suxrobshukurov/go-fiber/internal/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()

	app := fiber.New()

	// middlewares
	app.Use(recover.New())

	// handlers
	pages.NewHandler(app)

	app.Listen(":3000")
}

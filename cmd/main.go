package main

import (
	"log"
	"suxrobshukurov/go-fiber/config"
	"suxrobshukurov/go-fiber/internal/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// configs
	config.LoadConfig()
	serverConfig := config.NewServerConfig()

	app := fiber.New()

	// middlewares
	app.Use(recover.New())

	// handlers
	pages.NewHandler(app)

	if err := app.Listen(serverConfig.Addr); err != nil {
		log.Fatal(err)
	}
}

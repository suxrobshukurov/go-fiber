package main

import (
	"log"
	"suxrobshukurov/go-fiber/config"
	"suxrobshukurov/go-fiber/internal/pages"
	"suxrobshukurov/go-fiber/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	// configs
	config.LoadConfig()
	serverConfig := config.NewServerConfig()
	loggerConfig := config.NewLoggerConfig()

	engine := html.New("./layouts", ".html")

	// logger
	customLogger := logger.NewLogger(loggerConfig)

	// application
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// middlewares
	app.Use(slogfiber.New(customLogger))
	app.Use(recover.New())

	// handlers
	pages.NewHandler(app, customLogger)

	if err := app.Listen(serverConfig.Addr); err != nil {
		log.Fatal(err)
	}
}

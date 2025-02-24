package logger

import (
	"log"
	"log/slog"
	"os"
	"suxrobshukurov/go-fiber/config"
)

const (
	jsonType    string = "json"
	consoleType string = "console"
)

func NewLogger(config *config.LoggerConfig) *slog.Logger {
	var logger *slog.Logger
	opts := slog.HandlerOptions{
		Level: slog.Level(config.Level),
	}
	switch config.Format {
	case jsonType:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &opts))
	case consoleType:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &opts))
	default:
		log.Fatalf("Incorrect logger type: %v. Check .env LOG_FORMAT. Only [json|console]", config.Format)
	}
	return logger
}

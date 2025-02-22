package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	ServerPortEnv     string = "SERVER_PORT"
	ServerHostEnv     string = "SERVER_HOST"
	DefaultServerPort int    = 3000
	DefaultServerHost string = "127.0.0.1"

	DatabaseURLEnv     string = "DATABASE_URL"
	DatabaseDefaultURL string = "postgres://postgres@localhost:5432/postgres?sslmode=disable"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
		return
	}
	log.Println("Config loaded")
}

type ServerConfig struct {
	Port int
	Host string
	Addr string
}

func NewServerConfig() *ServerConfig {
	sc := ServerConfig{
		Host: getString(ServerHostEnv, DefaultServerHost),
		Port: getInt(ServerPortEnv, DefaultServerPort),
	}
	sc.Addr = fmt.Sprintf("%s:%d", sc.Host, sc.Port)
	return &sc
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

// func getBool(key string, defaultValue bool) bool {
// 	value := os.Getenv(key)
// 	b, err := strconv.ParseBool(value)
// 	if err != nil {
// 		return defaultValue
// 	}
// 	return b
// }

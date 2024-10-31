package main

import (
	"amelia-sh-proxy/internal/http"
	"amelia-sh-proxy/internal/http/routes"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	_ = godotenv.Load(".env")
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Stack()
	log.Logger = logger.Logger()
	routes.Discover()
	http.Fiber.Init()
}

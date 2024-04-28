package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"

	"pavel-fokin/ai/apps/ai-bot/internal/server"
)

// Config is the server configuration.
type Config struct {
	Server server.Config
}

// NewConfig creates a new configuration.
func NewConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	return cfg
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config := NewConfig()

	router, err := server.NewRouter()
	if err != nil {
		log.Fatalf("Failed to create router: %v", err)
	}
	server := server.New(config.Server, router)

	log.Println("Starting LikeIt HTTP server... ", config.Server.Port)
	go server.Start()

	<-ctx.Done()

	log.Println("Shutting down the LikeIt HTTP server...")
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Failed to shutdown the server: %v", err)
	}
	log.Println("LikeIt HTTP server shutdown successfully")
}

package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"

	"ai-chats/internal/app"
	"ai-chats/internal/infra/api"
	"ai-chats/internal/infra/db"
	"ai-chats/internal/infra/db/sqlite"
	"ai-chats/internal/infra/ollama"
	"ai-chats/internal/infra/pubsub"
	"ai-chats/internal/server"
	"ai-chats/internal/worker"
)

// Config is the service configuration.
type Config struct {
	Server server.Config
	DB     db.Config
}

// NewConfig creates a new configuration for the service.
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

	db := sqlite.New(config.DB.DATABASE_URL)
	defer db.Close()

	sqlite.CreateTables(db)
	sqlite.LoadFixtures(db)

	pubsub := pubsub.New()
	// defer pubsub.CloseAll()

	app := app.New(
		sqlite.NewChats(db),
		sqlite.NewUsers(db),
		ollama.NewOllamaClient(),
		sqlite.NewOllamaModels(db),
		pubsub,
		sqlite.NewTx(db),
	)

	// Initialize the crypto package and the signing key.
	api.InitSigningKey(config.Server.TokenSigningKey)

	// Setup the server.
	sse := server.NewSSEConnections()
	// defer sse.CloseAll()

	router := api.NewRouter(app, sse, pubsub)
	server := server.New(config.Server, router)

	// Setup the worker.
	worker := worker.New(pubsub)
	worker.SetupHandlers(app)

	log.Println("Starting AIChats HTTP server... ", config.Server.Port)
	go server.Start()

	log.Println("Starting AIChats worker...")
	worker.Start()

	// Wait for the shutdown signal.
	<-ctx.Done()
	sse.CloseAll()

	log.Println("Shutting down the AIChats worker...")
	worker.Shutdown()

	log.Println("Shutting down the AIChats HTTP server...")
	server.Shutdown()
	log.Println("AIChats HTTP server shutdown successfully")
}

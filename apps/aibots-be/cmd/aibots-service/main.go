package main

import (
	"context"
	"io/fs"
	"log"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app"
	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/db"
	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/db/sqlite"
	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/ollama"
	"pavel-fokin/ai/apps/ai-bots-be/internal/infra/pubsub"
	"pavel-fokin/ai/apps/ai-bots-be/internal/pkg/crypto"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"
	"pavel-fokin/ai/apps/ai-bots-be/internal/worker"
	"pavel-fokin/ai/apps/ai-bots-be/web"
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

	db, err := sqlite.NewDB(config.DB.DATABASE_URL)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	defer db.Close()

	if err := sqlite.CreateTables(db); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	pubsub := pubsub.New()
	// defer pubsub.CloseAll()

	ollamamodels, err := ollama.NewOllamaModels()
	if err != nil {
		log.Fatalf("Failed to load Ollama models: %v", err)
	}

	app := app.New(
		sqlite.NewChats(db),
		sqlite.NewUsers(db),
		sqlite.NewMessages(db),
		ollamamodels,
		pubsub,
	)

	// Initialize the crypto package and the signing key.
	apiutil.InitSigningKey(config.Server.TokenSigningKey)
	crypto.InitBcryptCost(14)

	// Setup the server.
	server := server.New(config.Server, pubsub)
	server.SetupAuthAPI(app)
	server.SetupChatAPI(app)
	server.SetupOllamaAPI(app)
	staticFS, _ := fs.Sub(web.Dist, "dist")
	server.SetupStaticRoutes(staticFS)

	// Setup the worker.
	worker := worker.New(pubsub)
	worker.SetupHandlers(app)

	log.Println("Starting AIBots HTTP server... ", config.Server.Port)
	go server.Start()

	log.Println("Starting AIBots worker...")
	worker.Start()

	// Wait for the shutdown signal.
	<-ctx.Done()

	log.Println("Shutting down the AIBots worker...")
	worker.Shutdown()

	log.Println("Shutting down the AIBots HTTP server...")
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Failed to shutdown the server: %v", err)
	}
	log.Println("AIBots HTTP server shutdown successfully")
}

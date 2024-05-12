package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"

	"pavel-fokin/ai/apps/ai-bots-be/internal/app"
	"pavel-fokin/ai/apps/ai-bots-be/internal/db"
	"pavel-fokin/ai/apps/ai-bots-be/internal/db/sqlite"
	"pavel-fokin/ai/apps/ai-bots-be/internal/llm"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server"
)

// Config is the server configuration.
type Config struct {
	Server server.Config
	DB     db.Config
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

	chatBot, err := llm.NewChatModel("llama3")
	if err != nil {
		log.Fatalf("Failed to create chat bot: %v", err)
	}

	db, err := sqlite.NewDB(config.DB.DATABASE_URL)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	defer db.Close()

	if err := sqlite.CreateTables(db); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	app := app.New(
		chatBot,
		sqlite.NewChats(db),
		sqlite.NewUsers(db),
	)

	server := server.New(config.Server)
	server.SetupAuthAPI(app)
	server.SetupChatAPI(app)

	log.Println("Starting AIBots HTTP server... ", config.Server.Port)
	go server.Start()

	<-ctx.Done()

	log.Println("Shutting down the AIBots HTTP server...")
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Failed to shutdown the server: %v", err)
	}
	log.Println("AIBots HTTP server shutdown successfully")
}

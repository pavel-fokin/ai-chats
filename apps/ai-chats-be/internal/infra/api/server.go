package api

import (
	"context"
	"log"
	"net/http"
	"time"
)

const maxShutdownTimeout = 3

// Config is the server configuration.
type Config struct {
	Port            string `env:"AICHATS_SERVER_PORT" envDefault:"8080"`
	TokenSigningKey string `env:"AICHATS_TOKEN_SIGNING_KEY" envDefault:"secret"`
}

// Server is the main server struct.
type Server struct {
	config Config
	server *http.Server
	sse    *SSEConnections
}

// NewServer creates a new server.
func NewServer(config Config) *Server {
	server := &http.Server{
		Addr: ":" + config.Port,
	}

	return &Server{
		config: config,
		server: server,
		sse:    NewSSEConnections(),
	}
}

// Start starts the server.
func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(maxShutdownTimeout)*time.Second,
	)
	defer cancel()

	s.sse.CloseAll()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown the server: %v", err)
	}
}

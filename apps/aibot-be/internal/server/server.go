package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const maxShutdownTimeout = 10

// Config is the server configuration.
type Config struct {
	Port string `env:"LIKEIT_SERVER_PORT" envDefault:"8080"`
}

// Server is the main server struct.
type Server struct {
	server *http.Server
}

// NewServer creates a new server.
func New(config Config, router chi.Router) *Server {
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	return &Server{
		server: server,
	}
}

func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(maxShutdownTimeout)*time.Second,
	)
	defer cancel()

	return s.server.Shutdown(ctx)
}

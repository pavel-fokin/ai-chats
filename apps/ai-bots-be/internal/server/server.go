package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const maxShutdownTimeout = 10

// Config is the server configuration.
type Config struct {
	Port string `env:"LIKEIT_SERVER_PORT" envDefault:"8080"`
}

// Server is the main server struct.
type Server struct {
	router *chi.Mux
	server *http.Server
}

// NewServer creates a new server.
func New(config Config) *Server {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	return &Server{
		router: router,
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

func (s *Server) SetupChatAPI(chat Chat) {
	s.router.Post("/api/messages", PostChatMessages(chat))
	// s.router.Post("/api/chats", PostChats(chat))
	// s.router.Get("/api/chats/{chat_id}/messages", GetChatMessages(chat))
}

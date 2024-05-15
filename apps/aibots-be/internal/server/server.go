package server

import (
	"context"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"pavel-fokin/ai/apps/ai-bots-be/internal/server/api"
	"pavel-fokin/ai/apps/ai-bots-be/internal/server/apiutil"
)

const maxShutdownTimeout = 3

// Config is the server configuration.
type Config struct {
	Port            string `env:"AIBOTS_SERVER_PORT" envDefault:"8080"`
	tokenSigningKey string `env:"AIBOTS_TOKEN_SIGNING_KEY" envDefault:"secret"`
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

	// Initialize the token signing key and validator.
	apiutil.InitSigningKey(config.tokenSigningKey)

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	return &Server{
		router: router,
		server: server,
	}
}

// Start starts the server.
func (s *Server) Start() {
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Duration(maxShutdownTimeout)*time.Second,
	)
	defer cancel()

	return s.server.Shutdown(ctx)
}

func (s *Server) SetupAuthAPI(auth api.Auth) {
	s.router.Post("/api/auth/login", api.LogIn(auth))
	s.router.Post("/api/auth/signup", api.SignUp(auth))
}

// SetupChatAPI sets up the chat API.
func (s *Server) SetupChatAPI(chat api.ChatApp) {
	s.router.Group(func(r chi.Router) {
		r.Use(api.AuthToken)
		r.Post("/api/chats", api.PostChats(chat))
		r.Get("/api/chats", api.GetChats(chat))
		r.Post("/api/chats/{uuid}/messages", api.PostMessages(chat))
		r.Get("/api/chats/{uuid}/messages", api.GetMessages(chat))

	})

	s.router.Get("/api/chats/{uuid}/events", api.GetEvents(chat))
}

func (s *Server) SetupStaticRoutes(static fs.FS) {
	fs := http.FileServerFS(static)

	s.router.Get("/", fs.ServeHTTP)
	s.router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		fs.ServeHTTP(w, r)
	})
	s.router.Get(
		"/assets/*", fs.ServeHTTP,
	)
}

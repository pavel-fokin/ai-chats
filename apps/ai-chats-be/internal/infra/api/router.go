package api

import (
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"ai-chats/web"
)

type App interface {
	Auth
	Chats
	Ollama
}

func (s *Server) SetupRoutes(app App, pubsub Subscriber) {
	r := chi.NewRouter()
	s.server.Handler = r

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/auth/login", LogIn(app, s.config.TokenSigningKey))
	r.Post("/api/auth/signup", SignUp(app, s.config.TokenSigningKey))

	r.Group(func(r chi.Router) {
		r.Use(AuthHeader(s.config.TokenSigningKey))
		r.Post("/api/chats", PostChats(app))
		r.Get("/api/chats", GetChats(app))
		r.Get("/api/chats/{uuid}", GetChat(app))
		r.Delete("/api/chats/{uuid}", DeleteChat(app))
		r.Post("/api/chats/{uuid}/generate-title", PostGenerateChatTitle(app))
		r.Post("/api/chats/{uuid}/messages", PostMessages(app))
		r.Get("/api/chats/{uuid}/messages", GetMessages(app))
	})

	r.Group(func(r chi.Router) {
		r.Use(AuthParam(s.config.TokenSigningKey))
		r.Get("/api/events/app", GetAppEvents(app, s.sse, pubsub))
		r.Get("/api/chats/{uuid}/events", GetChatEvents(app, s.sse, pubsub))
	})

	r.Group(func(r chi.Router) {
		r.Use(AuthHeader(s.config.TokenSigningKey))
		r.Get("/api/ollama/models", GetOllamaModels(app))
		r.Post("/api/ollama/models", PostOllamaModels(app))
		r.Delete("/api/ollama/models/{model}", DeleteOllamaModel(app))
	})

	staticFS, _ := fs.Sub(web.Dist, "dist")
	fs := http.FileServerFS(staticFS)

	r.Get("/", fs.ServeHTTP)
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		fs.ServeHTTP(w, r)
	})
	r.Get(
		"/assets/*", fs.ServeHTTP,
	)
}

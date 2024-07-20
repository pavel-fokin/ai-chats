package api

import (
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"pavel-fokin/ai/apps/ai-bots-be/internal/server"
	"pavel-fokin/ai/apps/ai-bots-be/web"
)

type App interface {
	Auth
	ChatApp
	OllamaApp
	Models
}

func NewRouter(app App, sse *server.SSEConnections, pubsub Subscriber) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/auth/login", LogIn(app))
	r.Post("/api/auth/signup", SignUp(app))

	r.Group(func(r chi.Router) {
		r.Use(AuthHeader)
		r.Post("/api/chats", PostChats(app))
		r.Get("/api/chats", GetChats(app))
		r.Get("/api/chats/{uuid}", GetChat(app))
		r.Delete("/api/chats/{uuid}", DeleteChat(app))
		r.Post("/api/chats/{uuid}/messages", PostMessages(app))
		r.Get("/api/chats/{uuid}/messages", GetMessages(app))
	})

	r.Group(func(r chi.Router) {
		r.Use(AuthParam)
		r.Get("/api/chats/{uuid}/events", GetEvents(app, sse, pubsub))
	})

	r.Group(func(r chi.Router) {
		r.Use(AuthHeader)
		r.Get("/api/models/library", GetModelsLibrary(app))
		r.Get("/api/ollama-models", GetOllamaModels(app))
		r.Post("/api/ollama-models", PostOllamaModels(app))
		r.Delete("/api/ollama-models/{model}", DeleteOllamaModel(app))
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

	return r
}

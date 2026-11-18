package routes

import (
	"Golang_API_CRUD01/Internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupNewRouters(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/posts/{id}", app.PostHandler.HandlerGetPostById)
	r.Get("/posts", app.PostHandler.HandlerCreatePost)

	return r
}

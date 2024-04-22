package router

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"go-tavern/controllers"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/characters", controllers.GetAllCharacters)
		r.Post("/characters", controllers.CreateCharacter)
		r.Get("/characters/{name}", controllers.DownloadCharacter)

		r.Post("/tokenize/{name}", controllers.CountTokens)

		r.Post("/completion/{name}", controllers.Completion)

		r.Post("/backgrounds", controllers.CreateBackground)
		r.Get("/backgrounds", controllers.GetAllBackgrounds)
		r.Get("/backgrounds/{name}", controllers.DownloadBackground)

		r.Post("/presets", controllers.CreatePreset)
		r.Get("/presets", controllers.GetAllPresets)
		r.Get("/presets/{name}", controllers.DownloadPreset)

		r.Post("/chats", controllers.CreateChat)
		r.Get("/chats", controllers.GetAllChats)
		r.Get("/chats/{name}", controllers.GetChat)
	})

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"}, // Replace with the correct origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})

	// Apply the CORS middleware to our top-level router, with the defaults.
	r.Use(c.Handler)

	return r
}

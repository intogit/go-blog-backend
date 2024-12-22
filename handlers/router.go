package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func CreateRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/healthcheck", HandleHealthCheck)
			r.Post("/createUser", HandleCreateUser)
			r.Get("/getUserByApiKey", HandleGetUserByApiKey)
			r.Post("/createFeed", HandleCreateFeed)
			r.Post("/createFeedFollow", HandleCreateFeedFollow)
			r.Get("/getAllFeedFollow", HandleGetAllFeedFollow)
			r.Delete("/deleteFeedFollow/{feed_follow_id}", HandleDeleteFeedFollow)
		})
		r.Route("/v2", func(r chi.Router) {
			r.Get("/healthcheck", HandleHealthCheck)
		})
	})

	return r
}

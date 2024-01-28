package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initializeRoutes(router *chi.Mux) {
	router.Route("/api/v1", func(r chi.Router) {
		// TODO: Add a route validation, maybe with some env key, to prevent someone to add data to your elastic search cluster
		r.Post("/bulk-save", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/search", func(w http.ResponseWriter, r *http.Request) {})
	})
}

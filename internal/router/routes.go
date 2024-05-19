package router

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MarcusXavierr/search-engine-api/internal/handler"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-chi/chi/v5"
)

func initializeRoutes(router *chi.Mux, esclient *elasticsearch.Client) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
			request, err := parseSearchQueryString(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			handler.HandleSearch(esclient, w, request)
		})
	})
}

func parseSearchQueryString(r *http.Request) (handler.SearchRequest, error) {
	sentence := r.URL.Query().Get("sentence")
	index := r.URL.Query().Get("index")
	strSize := r.URL.Query().Get("size") // Default size is 10
	unique := r.URL.Query().Get("unique") == "true"

	if sentence == "" || index == "" || strSize == "" {
		return handler.SearchRequest{}, errors.New("Incorrect values, no index or sentence or size provided")
	}

	size, err := strconv.Atoi(strSize)
	if err != nil {
		return handler.SearchRequest{}, errors.New("Size should be a number")
	}

	return handler.SearchRequest{
		Sentence: sentence,
		Index:    index,
		Size:     size,
		Unique:   unique,
	}, nil
}

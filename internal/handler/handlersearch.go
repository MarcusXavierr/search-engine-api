package handler

import (
	"encoding/json"
	"net/http"

	es "github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch"
	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch/search"
	"github.com/elastic/go-elasticsearch/v7"
)

type SearchRequest struct {
	Sentence string
	Index    string
	Size     int
	Unique   bool
}

type SearchResponse struct {
	Paragraphs   []es.ParagraphDocument `json:"paragraphs"`
	UniqueTitles bool                   `json:"unique_titles"`
}

func HandleSearch(esclient *elasticsearch.Client, w http.ResponseWriter, request SearchRequest) {
	query := search.MountBlogSearchQuery(request.Sentence, request.Size, request.Unique)
	data, err := search.Search(esclient.Search, request.Index, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := SearchResponse{
		Paragraphs:   data,
		UniqueTitles: request.Unique,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

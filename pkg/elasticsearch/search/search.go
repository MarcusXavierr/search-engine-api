package search

import (
	"io"
	"strings"

	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Query interface {
	GetQuery() io.Reader
}

func Search(search esapi.Search, index string, query Query) ([]elasticsearch.ParagraphDocument, error) {
	search(
		search.WithIndex(index),
		search.WithBody(strings.NewReader(`{ "query": { "match_all": {} } }`)),
	)

	return nil, nil
}

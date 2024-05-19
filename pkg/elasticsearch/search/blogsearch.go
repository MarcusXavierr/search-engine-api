package search

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

func (a BlogQuery) GetQuery() io.Reader {
	data, err := json.Marshal(a)
	if err != nil {
		log.Printf("Error while marshal query: %s", err)
		return strings.NewReader("")
	}

	return strings.NewReader(string(data))
}

func MountBlogSearchQuery(sentence string, size int, uniqueTitles bool) BlogQuery {
	var empty struct{}

	query := BlogQuery{
		Size: size,
		Query: &query{
			MultiMatch: &multiMatch{
				Query:  sentence,
				Fields: []string{"title^3", "content"},
			},
		},
		Hightlight: &Highlight{
			Fields:       []genericField{genericField{"content": empty}},
			NumFragments: 1,
		},
	}

	collapse := &collapse{
		Field: "title.keyword",
		InnerHits: innerHits{
			Name: "top_hit_for_title",
			Size: 1,
			Sort: []sort{{Score: "desc"}},
		},
	}

	if uniqueTitles {
		query.Collapse = collapse
	}

	return query
}

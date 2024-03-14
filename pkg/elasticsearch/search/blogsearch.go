package search

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

type BlogQuery struct {
	Size  int              `json:"size"`
	Query *blogSearchQuery `json:"query"`
	Aggs  *nlogSearchAggs  `json:"aggs"`
}

type multiMatch struct {
	Query  string   `json:"query"`
	Fields []string `json:"fields"`
}
type blogSearchQuery struct {
	MultiMatch *multiMatch `json:"multi_match"`
}
type terms struct {
	Field string `json:"field"`
	Size  int    `json:"size"`
}
type topHits struct {
	Size int `json:"size"`
}
type topHitForTitle struct {
	TopHits *topHits `json:"top_hits"`
}
type nlogSearchAggs struct {
	UniqueTitles *blogUniqueTitles `json:"unique_titles"`
	Aggs         *blogInnerAggs    `json:"aggs"`
}
type blogUniqueTitles struct {
	Terms *terms         `json:"terms"`
	Aggs  *blogInnerAggs `json:"aggs"`
}
type blogInnerAggs struct {
	TopHitForTitle *topHitForTitle `json:"top_hit_for_title"`
}

func (a BlogQuery) GetQuery() io.Reader {
	data, err := json.Marshal(a)
	if err != nil {
		log.Printf("Error while marshal query: %s", err)
		return strings.NewReader("")
	}

	return strings.NewReader(string(data))
}

func MountBlogSearchQuery(sentence string, size int, uniqueTitles bool) BlogQuery {
	topHitsSize := 1
	if !uniqueTitles {
		topHitsSize = 15
	}

	return BlogQuery{
		Size: 0,
		Query: &blogSearchQuery{
			MultiMatch: &multiMatch{
				Query:  sentence,
				Fields: []string{"title^3", "content"},
			},
		},
		Aggs: &nlogSearchAggs{
			UniqueTitles: &blogUniqueTitles{
				Terms: &terms{
					Field: "title.keyword",
					Size:  size,
				},
			},
			Aggs: &blogInnerAggs{
				TopHitForTitle: &topHitForTitle{
					TopHits: &topHits{
						Size: topHitsSize,
					},
				},
			},
		},
	}
}

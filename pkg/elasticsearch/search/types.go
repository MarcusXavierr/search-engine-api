package search

type BlogQuery struct {
	Size     int       `json:"size"`
	Query    *query    `json:"query"`
	Collapse *collapse `json:"collapse,omitempty"`
}

type query struct {
	MultiMatch *multiMatch `json:"multi_match"`
}

type multiMatch struct {
	Query  string   `json:"query"`
	Fields []string `json:"fields"`
}

type collapse struct {
	Field     string    `json:"field"`
	InnerHits innerHits `json:"inner_hits"`
}

type innerHits struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	Sort []sort `json:"sort"`
}

type sort struct {
	Score string `json:"_score,omitempty"`
}

package elasticsearch

type ElasticDocumentData interface {
	GetID() string
}

type ParagraphDocument struct {
	Content  string   `json:"content"`
	Title    string   `json:"title"`
	Date     string   `json:"date"`
	ObjectID string   `json:"objectID"`
	Tags     []string `json:"tags"`
	Uri      string   `json:"uri"`
}

func (p ParagraphDocument) GetID() string {
	return p.ObjectID
}

package search

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Query interface {
	GetQuery() io.Reader
}

func Search(search esapi.Search, index string, query Query) ([]elasticsearch.ParagraphDocument, error) {
	if index == "" {
		return nil, errors.New("Index cannot be empty")
	}

	result, err := search(
		search.WithBody(query.GetQuery()),
		search.WithIndex(index),
	)

	if err != nil {
		return nil, errors.New("Error while searching: " + err.Error())
	}

	if result.IsError() {
		return nil, errors.New("Error while searching and err is nil: " + result.String())
	}

	return ParseBody(result.Body)
}

func ParseBody(body io.Reader) ([]elasticsearch.ParagraphDocument, error) {
	buffer, err := io.ReadAll(body)
	if err != nil {
		return nil, errors.New("Error while reading body: " + err.Error())
	}

	var rawData map[string]interface{}
	if err := json.Unmarshal(buffer, &rawData); err != nil {
		return nil, errors.New("Error while unmarshalling: " + err.Error())
	}

	paragraphs := make([]elasticsearch.ParagraphDocument, 0)
	hits, ok := rawData["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, errors.New("Error while asserting hits")
	}

	for _, hit := range hits {
		paragraph, err := parseHit(hit.(map[string]interface{}))
		if err != nil {
			return nil, errors.New("Error while parsing hit: " + err.Error())
		}
		paragraphs = append(paragraphs, paragraph)
	}

	return paragraphs, nil
}

func parseHit(hit map[string]interface{}) (elasticsearch.ParagraphDocument, error) {
	source, ok := hit["_source"].(map[string]interface{})
	if !ok {
		return elasticsearch.ParagraphDocument{}, errors.New("Error while asserting source")
	}

	bytes, err := json.Marshal(source)
	if err != nil {
		return elasticsearch.ParagraphDocument{}, errors.New("Error while marshalling paragraph")
	}

	var p elasticsearch.ParagraphDocument
	if err := json.Unmarshal(bytes, &p); err != nil {
		return elasticsearch.ParagraphDocument{}, errors.New("Error while unmarshalling paragraph")
	}

	p = useHighlightAsContent(p, hit)
	return p, nil
}

func useHighlightAsContent(paragraph elasticsearch.ParagraphDocument, hit map[string]interface{}) elasticsearch.ParagraphDocument {
	highlight, ok := hit["highlight"].(map[string]interface{})
	if !ok {
		return paragraph
	}

	content := highlight["content"].([]interface{})
	if len(content) == 0 {
		return paragraph
	}

	paragraph.Content = content[0].(string)
	return paragraph
}

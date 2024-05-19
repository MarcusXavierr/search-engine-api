package search

import (
	"reflect"
	"testing"

	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch"
)

func TestParseResponse(t *testing.T) {
	t.Run("Should parse the response correctly", func(t *testing.T) {

	})
}

func TestParseHit(t *testing.T) {
	t.Run("Should parse the paragraph correctly", func(t *testing.T) {
		mockParagraph := map[string]interface{}{
			"content":  "This is a test",
			"title":    "Test",
			"date":     "2021-01-01",
			"objectID": "1",
			"tags":     []string{"test"},
			"uri":      "http://test.com",
		}
		mockHit := map[string]interface{}{
			"_source": mockParagraph,
		}

		want := elasticsearch.ParagraphDocument{
			Content:  "This is a test",
			Title:    "Test",
			Date:     "2021-01-01",
			ObjectID: "1",
			Tags:     []string{"test"},
			Uri:      "http://test.com",
		}

		got, err := parseHit(mockHit)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}

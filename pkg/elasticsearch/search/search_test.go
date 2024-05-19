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
	mockParagraph := map[string]interface{}{
		"content":  "This is a test",
		"title":    "Test",
		"date":     "2021-01-01",
		"objectID": "1",
		"tags":     []string{"test"},
		"uri":      "http://test.com",
	}

	t.Run("Should parse the paragraph correctly", func(t *testing.T) {
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

		parseParagraphAndCompare(t, want, mockHit)
	})

	t.Run("Should use highlight if it exists", func(t *testing.T) {
		mockHit := map[string]interface{}{
			"_source": mockParagraph,
			"highlight": map[string]interface{}{
				"content": []interface{}{"resumed content", "other resume"},
			},
		}

		want := elasticsearch.ParagraphDocument{
			Content:  "resumed content",
			Title:    "Test",
			Date:     "2021-01-01",
			ObjectID: "1",
			Tags:     []string{"test"},
			Uri:      "http://test.com",
		}

		parseParagraphAndCompare(t, want, mockHit)
	})

	t.Run("Should return an error if the content is not a string", func(t *testing.T) {
		_, err := parseHit(mockParagraph)
		if err == nil {
			t.Error("expect error, found nil")
		}
	})
}

func parseParagraphAndCompare(t testing.TB, want elasticsearch.ParagraphDocument, mockHit map[string]interface{}) {
	t.Helper()
	got, err := parseHit(mockHit)
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected %v but got %v", want, got)
	}
}

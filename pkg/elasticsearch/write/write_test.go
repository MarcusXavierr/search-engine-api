package write_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"reflect"
	"testing"

	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch"
	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch/write"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

const indexName = "any_index"

var (
	mockDocuments = []elasticsearch.ParagraphDocument{
		{"doc1", "doc1", "doc1", "doc1", []string{}, "doc1"},
		{"doc2", "doc2", "doc2", "doc2", []string{"tag1", "tag2"}, "doc2"},
	}
	mockFailureFunction = func(_ context.Context, _ esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
		if err != nil {
			log.Printf("ERROR: %s", err)
		} else {
			log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
		}
	}

	mockedBulkIndexes = []esutil.BulkIndexerItem{
		{
			Index:      indexName,
			Action:     "create",
			DocumentID: mockDocuments[0].ObjectID,
			Body:       returnReaderFromDocument(mockDocuments[0]),
			OnFailure:  mockFailureFunction,
		},
		{
			Index:      indexName,
			Action:     "create",
			DocumentID: mockDocuments[1].ObjectID,
			Body:       returnReaderFromDocument(mockDocuments[1]),
			OnFailure:  mockFailureFunction,
		},
	}
)

func TestMountBulkIndexerData(t *testing.T) {
	t.Run("return an error when ParagraphDocument list is empty", func(t *testing.T) {
		_, err := write.MountBulkIndexerData(indexName, []elasticsearch.ParagraphDocument{})

		want := write.EmptyDocumentList

		if err != want {
			t.Errorf("expected %s but got %s", want, err)
		}
	})

	t.Run("Gets a list of documents and mounts ndJson as string", func(t *testing.T) {

		got, err := write.MountBulkIndexerData(indexName, mockDocuments)

		if err != nil {
			t.Errorf("error found: %s", err)
		}

		want := mockedBulkIndexes
		if !deepEqual(got, want) {
			t.Errorf("expected \n%v \nbut got \n%v", want, got)
		}

	})
}

func returnReaderFromDocument(doc elasticsearch.ParagraphDocument) io.ReadSeeker {
	data, err := json.Marshal(doc)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(data)
}

func deepEqual(got []esutil.BulkIndexerItem, want []esutil.BulkIndexerItem) bool {
	for i := 0; i < len(got); i++ {
		equals := bulkIndexerItemDeepEqual(got[i], want[i])
		if !equals {
			return false
		}
	}

	return true
}
func bulkIndexerItemDeepEqual(itemOne esutil.BulkIndexerItem, itemTwo esutil.BulkIndexerItem) bool {
	bodyOne := []byte{}
	bodyTwo := []byte{}

	itemTwo.Body.Read(bodyOne)
	itemTwo.Body.Read(bodyTwo)

	equal := itemOne.Action == itemTwo.Action &&
		itemOne.Index == itemTwo.Index &&
		itemOne.DocumentID == itemTwo.DocumentID &&
		reflect.DeepEqual(bodyOne, bodyTwo)

	return equal
}

package write

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type ndJson = string

var (
	EmptyDocumentList     = errors.New("The document list should not be empty")
	ParsingDocumentsError = errors.New("There was an error parsing documents: ")
	BulkInsertError       = errors.New("There was an error inserting documents on BulkIndexer.Add: ")
)

// This function gets a lot of data and insert it on elastic search at once
func BulkInsert[T elasticsearch.ElasticDocumentData](indexName string, documents []T, client esutil.BulkIndexer) error {
	items, err := MountBulkIndexerData(indexName, documents)

	if err != nil {
		return errors.Join(ParsingDocumentsError, err)
	}

	for _, item := range items {
		err := client.Add(context.Background(), item)
		if err != nil {
			return errors.Join(BulkInsertError, err)
		}
	}

	if err := client.Close(context.Background()); err != nil {
		return err
	}
	return nil
}

func MountBulkIndexerData[T elasticsearch.ElasticDocumentData](indexName string, documents []T) ([]esutil.BulkIndexerItem, error) {
	if len(documents) == 0 {
		return []esutil.BulkIndexerItem{}, EmptyDocumentList
	}

	slice := []esutil.BulkIndexerItem{}
	for _, doc := range documents {
		item, err := makeBulkIndexerItem(indexName, doc)
		if err != nil {
			return []esutil.BulkIndexerItem{}, err
		}

		slice = append(slice, item)
	}
	return slice, nil
}

func makeBulkIndexerItem[T elasticsearch.ElasticDocumentData](indexName string, doc T) (esutil.BulkIndexerItem, error) {
	data, err := json.Marshal(doc)
	if err != nil {
		return esutil.BulkIndexerItem{}, err
	}

	item := esutil.BulkIndexerItem{
		Index:      indexName,
		Action:     "create",
		DocumentID: doc.GetID(),
		Body:       bytes.NewReader(data),
		OnFailure: func(_ context.Context, _ esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
			if err != nil {
				fmt.Printf("ERROR: %s", err)
				log.Printf("ERROR: %s", err)
			} else {
				fmt.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
				log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
			}
		},
	}

	return item, nil
}

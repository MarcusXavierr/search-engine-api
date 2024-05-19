package main

import (
	"fmt"
	"os"
	"time"

	es "github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch"
	"github.com/MarcusXavierr/search-engine-api/pkg/elasticsearch/write"
	"github.com/MarcusXavierr/search-engine-api/pkg/reader"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/joho/godotenv"
)

func main() {
	data, err := reader.ReadJson[[]es.ParagraphDocument](os.DirFS("/home/marcus/Projects/personal-website/public"), "elastic_data.json")
	if err != nil {
		panic(err)
	}

	godotenv.Load()

	username := os.Getenv("ELASTIC_USERNAME")
	password := os.Getenv("ELASTIC_PASSWORD")
	url := os.Getenv("ELASTIC_URL")

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Username:  username,
		Password:  password,
		Addresses: []string{url},
	})

	if err != nil {
		panic(fmt.Sprintf("Error while connecting to client: %s", err))
	}

	bulkIndexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		NumWorkers:    2,
		FlushBytes:    int(5e+6),
		FlushInterval: 30 * time.Second,
		Client:        client,
		Index:         "paragraph",
	})

	if err != nil {
		panic(fmt.Sprintf("Error creating bulkIndexer: %s\n", err))
	}

	err = write.BulkInsert("paragraph", data, bulkIndexer)
	if err != nil {
		panic(fmt.Sprintf("Error while bulk insert: %s", err))
	}
}

package main

import (
	"os"

	"github.com/MarcusXavierr/search-engine-api/internal/router"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/joho/godotenv"
)

func main() {
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
		panic(err)
	}

	router.Initialize(client)
}

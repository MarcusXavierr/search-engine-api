package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Initialize() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	initializeRoutes(router)
	// Just an Hello world route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello friend\n"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := "0.0.0.0:" + port
	fmt.Println("Running on port", address)

	if err := http.ListenAndServe(address, router); err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pruser/allegro-feed-generator/config"
	"github.com/pruser/allegro-feed-generator/request"

	"github.com/gorilla/handlers"
)

func main() {
	log.Printf("Allegro Feed Generator is starting")
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	config.Print()

	handler := request.NewRequestHandler(config.WebAPIKey, config.UrlBase)
	http.HandleFunc("/feed", handler.CreateFeed)
	http.HandleFunc("/health", handler.Health)

	log.Print("Waiting for requests on port 8080")

	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

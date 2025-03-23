package main

import (
	"github.com/5bbclub/chatbot-character-service/api"
	"github.com/5bbclub/chatbot-character-service/cmd/crawler/config"
	"log"
)

func main() {
	// Load configuration
	if err := config.LoadConfig("config/config.toml"); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	conf := config.GetConfig()

	// Start the API server
	log.Printf("Starting API server on port %d", conf.Server.ApiPort)
	if err := api.StartServer(); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}

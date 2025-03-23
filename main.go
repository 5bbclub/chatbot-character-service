package main

import (
	"github.com/5bbclub/chatbot-character-service/api"
	"github.com/5bbclub/chatbot-character-service/config"
	"log"
)

func main() {
	// Load configuration
	if err := config.LoadConfig("config/config.toml"); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Start the API server
	serverConfig := config.GetConfig().Server
	log.Printf("Starting API server on port %d", serverConfig.ApiPort)
	if err := api.StartServer(); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}

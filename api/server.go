package api

import (
	"fmt"
	"github.com/5bbclub/chatbot-character-service/cmd/api/config"
	"log"
	"net/http"
)

func StartServer() error {
	port := config.GetConfig().Server.ApiPort
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "API server is running.")
	})

	log.Printf("Starting API server on port %d...", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

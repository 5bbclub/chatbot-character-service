package crawler

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() error {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Crawler server is running.")
	})

	log.Println("Starting Crawler server on port 9090...")
	return http.ListenAndServe(":9090", nil)
}

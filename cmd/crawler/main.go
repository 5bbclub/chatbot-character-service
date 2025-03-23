// cmd/crawler/main.go
package main

import (
	"github.com/5bbclub/chatbot-character-service/cmd/crawler/config"
	"github.com/5bbclub/chatbot-character-service/crawler"
	"log"
)

func main() {
	// 설정 파일 경로
	configPath := "cmd/crawler/config/config.toml"

	// 설정 로드
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Loaded configuration: %+v\n", conf)
	crawler.Run(conf)
}

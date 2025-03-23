// crawler/server.go
package crawler

import (
	"github.com/5bbclub/chatbot-character-service/utils/database"
	"log"
	"time"

	"github.com/5bbclub/chatbot-character-service/cmd/crawler/config"
	"github.com/5bbclub/chatbot-character-service/crawler/fetchers"
	"github.com/5bbclub/chatbot-character-service/crawler/processors"
	"github.com/5bbclub/chatbot-character-service/crawler/scheduler"
)

func Run(conf *config.Config) {

	// 데이터베이스 초기화
	dbConfig := conf.Database
	err := database.InitDB(dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	if err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	//defer database.CloseDB()

	// 스케줄러 초기화
	jobScheduler := scheduler.NewScheduler()

	// 각 서비스 설정을 처리
	for _, service := range conf.Services {
		log.Printf("Registering service: %s (Interval: %d seconds)", service.Name, service.Interval)

		var fetcher fetchers.DataFetcher

		// 서비스에 맞는 Fetcher 생성
		switch service.Name {
		case "Babechat":
			fetcher = fetchers.NewBabechatFetcher(service.Endpoint)
		case "Wrtn":
			//fetcher = fetchers.NewWrtnFetcher(service.Endpoint)
		case "Lofan":
			// Lofan용 Fetcher를 추가할 때 필요
			log.Printf("Lofan fetcher is not implemented yet.")
			continue
		default:
			log.Printf("No fetcher available for service: %s", service.Name)
			continue
		}

		// 작업 추가
		jobScheduler.AddJob(
			service.Name,
			time.Duration(service.Interval)*time.Second,
			func() {
				log.Printf("Running fetch and process for service: %s", service.Name)
				err := processors.ProcessData(fetcher)
				if err != nil {
					log.Printf("⛔ Error processing %s data: %v", service.Name, err)
				}
			},
		)
	}

	// 스케줄러 시작
	jobScheduler.Start()

	// 프로그램 실행 중 유지
	select {}
}

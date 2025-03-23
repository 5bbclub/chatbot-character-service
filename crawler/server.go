// crawler/server.go
package crawler

import (
	"github.com/5bbclub/chatbot-character-service/cmd/crawler/config"
	"github.com/5bbclub/chatbot-character-service/crawler/fetchers"
	"github.com/5bbclub/chatbot-character-service/crawler/processors"
	"github.com/5bbclub/chatbot-character-service/crawler/scheduler"
	"github.com/5bbclub/chatbot-character-service/utils/database"
	"log"
)

func Run(conf *config.Config) {

	// 데이터베이스 초기화
	dbConfig := conf.Database
	err := database.InitDB(dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	if err != nil {
		log.Fatalf("❌ Failed to initialize database: %v", err)
	}
	if database.DB == nil {
		log.Fatalf("❌ Database is not initialized")
	}
	//defer database.CloseDB()

	// 각 서비스 설정을 처리
	for _, service := range conf.Services {
		log.Printf("Registering service: %s (Interval: %d seconds)", service.Name, service.Interval)

		// 서비스에 맞는 Fetcher 생성
		switch service.Name {
		case "Babechat":
			babechatScheduler := &scheduler.JobScheduler{
				Name:          "babechat",
				FetcherImpl:   fetchers.NewBabechatFetcher(conf),
				ProcessorImpl: processors.NewBabeChatProcessor(conf, database.DB),
			}
			go babechatScheduler.Start()
		case "Wrtn":
			wrtnScheduler := &scheduler.JobScheduler{
				Name:          "wrtn",
				FetcherImpl:   fetchers.NewWrtnFetcher(conf),
				ProcessorImpl: processors.NewWrtnProcessor(conf, database.DB),
			}
			go wrtnScheduler.Start()
		case "Lofan":
			// Lofan용 Fetcher를 추가할 때 필요
			log.Printf("Lofan fetcher is not implemented yet.")
			continue
		default:
			log.Printf("No fetcher available for service: %s", service.Name)
			continue
		}
	}

	// 프로그램 실행 중 유지
	select {}
}

// crawler/processors/processors.go
package processors

import (
	"encoding/json"
	"github.com/5bbclub/chatbot-character-service/crawler/fetchers"
	"github.com/5bbclub/chatbot-character-service/utils/database"
	"log"
)

// ProcessData는 공통적인 데이터를 처리하고 저장합니다.
func ProcessData(fetcher fetchers.DataFetcher) error {
	log.Println("🛠️ Processing data...")

	// 데이터 가져오기
	data, err := fetcher.FetchData()
	if err != nil {
		return err
	}

	// 데이터가 없을 경우 메시지 출력
	if len(data) == 0 {
		log.Println("⚠️ No data to process")
		return nil
	}

	// Data 저장 리팩터링
	for _, item := range data {
		jsonData, err := json.Marshal(item)
		if err != nil {
			log.Printf("⛔ Failed to marshal data: %v", err)
			continue
		}

		// GORM을 통해 데이터 저장
		err = database.SaveServiceData(fetcher.GetServiceName(), string(jsonData))
		if err != nil {
			log.Printf("⛔ Failed to save data to database: %v", err)
		}
	}

	log.Println("✅ Successfully processed and saved data.")
	return nil
}

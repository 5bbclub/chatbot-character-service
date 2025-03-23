package scheduler

import (
	"github.com/5bbclub/chatbot-character-service/crawler/fetchers"
	"github.com/5bbclub/chatbot-character-service/crawler/processors"
	"log"
)

type JobScheduler struct {
	Name          string               // Scheduler 이름 (서비스 명칭)
	FetcherImpl   fetchers.Fetcher     // Fetcher 구현체
	ProcessorImpl processors.Processor // Processor 구현체
}

// Start: Fetcher와 Processor를 실행
func (s *JobScheduler) Start() {
	log.Printf("[Scheduler: %s] Starting...", s.Name)

	// Fetcher와 Processor 간 데이터 채널 공유
	dataChannel := make(chan interface{}, 500)

	// Fetcher 초기화 및 실행
	s.FetcherImpl.SetOutputChannel(dataChannel)
	go s.FetcherImpl.Start()

	// Processor 초기화 및 실행
	s.ProcessorImpl.SetInputChannel(dataChannel)
	go s.ProcessorImpl.Start()

	log.Printf("[Scheduler: %s] All components started", s.Name)
}

// crawler/scheduler/scheduler.go
package scheduler

import (
	"log"
	"time"
)

type Job struct {
	Name     string
	Interval time.Duration
	Action   func()
}

type Scheduler struct {
	jobs []Job
	done chan bool
}

// NewScheduler는 새로운 스케줄러 인스턴스를 반환합니다.
func NewScheduler() *Scheduler {
	return &Scheduler{
		jobs: []Job{},
		done: make(chan bool),
	}
}

// AddJob은 스케줄러에 새로운 작업을 추가합니다.
func (s *Scheduler) AddJob(name string, interval time.Duration, action func()) {
	job := Job{
		Name:     name,
		Interval: interval,
		Action:   action,
	}
	s.jobs = append(s.jobs, job)
}

// Start는 스케줄러가 작업을 주기적으로 실행하도록 시작합니다.
func (s *Scheduler) Start() {
	log.Println("⏱️ Starting the job scheduler...")
	for _, job := range s.jobs {
		go func(j Job) {
			ticker := time.NewTicker(j.Interval)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					log.Printf("🚀 Executing job: %s\n", j.Name)
					j.Action()
				case <-s.done:
					return
				}
			}
		}(job)
	}
}

// Stop은 스케줄러를 종료합니다.
func (s *Scheduler) Stop() {
	log.Println("🛑 Stopping the job scheduler...")
	close(s.done)
}
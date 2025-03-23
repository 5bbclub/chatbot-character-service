package fetchers

import (
	"time"
)

type BaseFetcher struct {
	Name          string           // 서비스 이름
	Interval      time.Duration    // 데이터 fetch 주기
	OutputChannel chan interface{} // 처리기로 전달될 채널
}

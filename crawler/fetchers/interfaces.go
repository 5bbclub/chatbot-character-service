// crawler/fetchers/interfaces.go
package fetchers

type Fetcher interface {
	Start()                            // 데이터를 가져와서 전달
	SetOutputChannel(chan interface{}) // 데이터를 전달할 채널 설정
	Fetch() ([]byte, error)            // 데이터를 가져옴
}

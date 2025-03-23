package processors

type Processor interface {
	Start()
	SetInputChannel(chan interface{}) // 데이터를 받을 채널 설정
	Process(data interface{}) error   // 데이터를 처리
}

package processors

import (
	"gorm.io/gorm"
)

type BaseProcessor struct {
	Name         string           // 서비스 이름
	InputChannel chan interface{} // 데이터를 받는 채널
	DB           *gorm.DB         // 데이터베이스 연결
}

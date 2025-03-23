package models

import (
	"time"

	"gorm.io/gorm"
)

type ServiceProvider struct {
	ID          uint           `gorm:"primaryKey"`                             // 기본 키
	Name        string         `gorm:"uniqueIndex;type:varchar(255);not null"` // 서비스 제공자 이름 (유니크)
	Description string         `gorm:"type:text"`                              // 서비스 제공자 설명
	WebsiteURL  string         `gorm:"type:varchar(255)"`                      // 웹사이트 URL
	CreatedAt   time.Time      `gorm:"autoCreateTime"`                         // 생성 시간
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`                         // 수정 시간
	DeletedAt   gorm.DeletedAt `gorm:"index"`                                  // 소프트 삭제를 위한 필드
	Characters  []Character    `gorm:"foreignKey:ServiceProviderID"`           // 캐릭터들과 1:N 관계
}

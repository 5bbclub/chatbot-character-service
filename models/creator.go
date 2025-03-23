package models

import (
	"time"

	"gorm.io/gorm"
)

type Creator struct {
	ID          uint           `gorm:"primaryKey"`                             // 기본 키
	Name        string         `gorm:"type:varchar(255);not null"`             // 제작자 이름
	Email       string         `gorm:"uniqueIndex;type:varchar(255);not null"` // 제작자 이메일 (유니크)
	LinkedInURL string         `gorm:"type:varchar(255)"`                      // LinkedIn URL
	CreatedAt   time.Time      `gorm:"autoCreateTime"`                         // 생성 시간
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`                         // 수정 시간
	DeletedAt   gorm.DeletedAt `gorm:"index"`                                  // 소프트 삭제 필드
	Characters  []Character    `gorm:"many2many:character_creators"`           // 다대다 관계 (캐릭터)
}

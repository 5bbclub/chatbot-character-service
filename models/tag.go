package models

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID          uint           `gorm:"primaryKey"`                             // 기본 키
	Name        string         `gorm:"uniqueIndex;type:varchar(255);not null"` // 태그 이름 (유니크)
	Description string         `gorm:"type:text"`                              // 태그 설명
	CreatedAt   time.Time      `gorm:"autoCreateTime"`                         // 생성 시간
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`                         // 수정 시간
	DeletedAt   gorm.DeletedAt `gorm:"index"`                                  // 소프트 삭제 필드
	Characters  []Character    `gorm:"many2many:character_tags"`               // 다대다 관계 (캐릭터)
}

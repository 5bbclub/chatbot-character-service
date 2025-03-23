package models

import (
	"time"
)

type CharacterTag struct {
	ID          uint      `gorm:"primaryKey"`     // 기본 키
	CharacterID uint      `gorm:"not null"`       // 외래 키: 캐릭터 ID
	TagID       uint      `gorm:"not null"`       // 외래 키: 태그 ID
	CreatedAt   time.Time `gorm:"autoCreateTime"` // 생성 시간
}

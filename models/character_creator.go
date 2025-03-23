package models

import (
	"time"
)

type CharacterCreator struct {
	ID          uint      `gorm:"primaryKey"`        // 기본 키
	CharacterID uint      `gorm:"not null"`          // 외래 키: 캐릭터 ID
	CreatorID   uint      `gorm:"not null"`          // 외래 키: 제작자 ID
	Role        string    `gorm:"type:varchar(255)"` // 제작자 역할 (예: 디자이너, 스크립트 작성자)
	CreatedAt   time.Time `gorm:"autoCreateTime"`    // 생성 시간
}

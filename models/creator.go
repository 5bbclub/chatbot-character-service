package models

import (
	"time"

	"gorm.io/gorm"
)

type Creator struct {
	ID                uint           `gorm:"primaryKey"`                   // 기본 키
	ServiceProviderID uint           `gorm:"not null;index"`               // 외래 키: 서비스 제공자 ID
	Name              string         `gorm:"type:varchar(255);not null"`   // 제작자 이름
	Email             string         `gorm:"type:varchar(255);not null"`   // 제작자 이메일
	CreatedAt         time.Time      `gorm:"autoCreateTime"`               // 생성 시간
	UpdatedAt         time.Time      `gorm:"autoUpdateTime"`               // 수정 시간
	DeletedAt         gorm.DeletedAt `gorm:"index"`                        // 소프트 삭제 필드
	Characters        []Character    `gorm:"many2many:character_creators"` // 다대다 관계 (캐릭터)
}

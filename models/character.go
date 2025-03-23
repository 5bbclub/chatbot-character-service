package models

import (
	"time"

	"gorm.io/gorm"
)

type Character struct {
	ID                uint            `gorm:"primaryKey"`                   // 기본 키
	ServiceProviderID uint            `gorm:"not null;index"`               // 외래 키: 서비스 제공자 ID
	Name              string          `gorm:"type:varchar(255);not null"`   // 캐릭터 이름
	Description       string          `gorm:"type:text"`                    // 캐릭터 설명
	ProfileImageURL   string          `gorm:"type:varchar(255)"`            // 프로필 이미지 URL
	CreatedAt         time.Time       `gorm:"autoCreateTime"`               // 생성 시간
	UpdatedAt         time.Time       `gorm:"autoUpdateTime"`               // 수정 시간
	DeletedAt         gorm.DeletedAt  `gorm:"index"`                        // 소프트 삭제 필드
	ServiceProvider   ServiceProvider `gorm:"foreignKey:ServiceProviderID"` // 1:N 관계에서 서비스 제공자 참조
	Creators          []Creator       `gorm:"many2many:character_creators"` // 다대다 관계 (제작자)
	Tags              []Tag           `gorm:"many2many:character_tags"`     // 다대다 관계 (태그)
}

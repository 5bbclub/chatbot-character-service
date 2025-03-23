package models

import (
	"time"

	"gorm.io/gorm"
)

type Character struct {
	ID                uint            `gorm:"primaryKey"`                   // 기본 키
	ServiceProviderID uint            `gorm:"not null;index"`               // 외래 키: 서비스 제공자 ID
	InternalID        string          `gorm:"type:varchar(255);unique"`     // 내부 ID
	Name              string          `gorm:"type:varchar(255);not null"`   // 캐릭터 이름
	Description       string          `gorm:"type:text"`                    // 캐릭터 설명
	ProfileImageURL   string          `gorm:"type:varchar(255)"`            // 프로필 이미지 URL
	ChatCount         uint            `gorm:"default:0"`                    // 채팅 횟수
	ChatUserCount     uint            `gorm:"default:0"`                    // 채팅 사용자 수
	LikeCount         uint            `gorm:"default:0"`                    // 좋아요 횟수
	CountryCode       string          `gorm:"type:varchar(2)"`              // 국가 코드
	IsAdult           bool            `gorm:"default:false"`                // 성인 캐릭터 여부
	IsFeatured        bool            `gorm:"default:false"`                // 추천 캐릭터 여부
	CreatedAt         time.Time       `gorm:"autoCreateTime"`               // 생성 시간
	UpdatedAt         time.Time       `gorm:"autoUpdateTime"`               // 수정 시간
	DeletedAt         gorm.DeletedAt  `gorm:"index"`                        // 소프트 삭제 필드
	ServiceProvider   ServiceProvider `gorm:"foreignKey:ServiceProviderID"` // 1:N 관계에서 서비스 제공자 참조
	Creators          []Creator       `gorm:"many2many:character_creators"` // 다대다 관계 (제작자)
	Tags              []Tag           `gorm:"many2many:character_tags"`     // 다대다 관계 (태그)
}

type CharacterInput struct {
	ServiceProviderName string   `json:"service_provider_name"` // 서비스 제공자 이름
	InternalID          string   `json:"internal_id"`           // 내부 ID
	Name                string   `json:"name"`                  // 캐릭터 이름
	Description         string   `json:"description"`           // 캐릭터 설명
	ProfileImageURL     string   `json:"profile_image_url"`     // 캐릭터 프로필 이미지 URL
	TagNames            []string `json:"tags"`                  // 태그 이름 목록
	Creator             Creator  `json:"creator"`               // 제작자 정보
}

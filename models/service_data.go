package models

import "time"

// ServiceData는 GORM 모델로 `service_data` 테이블에 매핑됩니다.
type ServiceData struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	ServiceName string    `gorm:"size:255;not null" json:"service_name"`
	Data        string    `gorm:"type:json;not null" json:"data"` // JSON 데이터를 저장
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

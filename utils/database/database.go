package database

import (
	"fmt"
	"gorm.io/gorm"
	"log"

	"github.com/5bbclub/chatbot-character-service/models"
	"gorm.io/driver/mysql"
)

// DB는 GORM 데이터베이스 핸들을 전역적으로 저장합니다.
var DB *gorm.DB

// InitDB는 GORM을 사용해 MySQL 데이터베이스에 연결 및 모델 마이그레이션을 수행합니다.
func InitDB(user, password, host, port, database string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	// GORM 연결 초기화
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// GORM 자동 마이그레이션 (테이블 생성 또는 업데이트)
	if err := db.AutoMigrate(&models.ServiceData{}); err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}

	log.Println("✅ Successfully connected to the database and migrated models.")
	DB = db
	return nil
}

// SaveServiceData는 데이터를 GORM을 통해 저장합니다.
func SaveServiceData(serviceName string, jsonData string) error {
	data := models.ServiceData{
		ServiceName: serviceName,
		Data:        jsonData,
	}

	// DB에 데이터 저장
	if err := DB.Create(&data).Error; err != nil {
		return fmt.Errorf("failed to insert service data: %w", err)
	}
	return nil
}

package main

import (
	"log"

	"github.com/5bbclub/chatbot-character-service/models" // 모델 경로
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql" // MySQL 드라이버
)

func main() {
	// MySQL 연결 정보
	dsn := "crawler_user:crawler_password@tcp(127.0.0.1:3306)/dipping?charset=utf8mb4&parseTime=True&loc=Local"

	// GORM 데이터베이스 연결
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to the database: %v", err)
	}

	// 데이터베이스 마이그레이션 실행
	if err := db.AutoMigrate(
		&models.ServiceProvider{},
		&models.Character{},
		&models.Creator{},
		&models.CharacterCreator{},
		&models.Tag{},
		&models.CharacterTag{},
	); err != nil {
		log.Fatalf("❌ Failed to migrate database schema: %v", err)
	}
	log.Println("✅ Database migration completed!")
}

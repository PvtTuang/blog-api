package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase(maxRetries int, delay time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 1; i <= maxRetries; i++ {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Printf("พยายามเชื่อมต่อกับฐานข้อมูลครั้งที่: %d\n ", i)
			return db, nil
		}

		log.Printf("พยายามครั้งที่ %d ล้มเหลว: %v", i, err)

		if i < maxRetries {
			time.Sleep(delay)
		}
	}

	return nil, fmt.Errorf("ไม่สามารถเชื่อมต่อกับฐานข้อมูลได้หลังจากพยายาม %d ครั้ง: %v", maxRetries, err)
}

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("ข้อผิดพลาดในการโหลดไฟล์ .env: %v", err)
	}

	db, err := connectDatabase(5, 5*time.Second)
	if err != nil {
		log.Fatalf("เชื่อมต่อฐานข้อมูลไม่สำเร็จ: %v", err)
	}

	db.AutoMigrate(&Post{})

	DB = db
}

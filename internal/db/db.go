package db

import (
	"Basic/internal/TaskService"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=eourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&TaskService.Task{}); err != nil {
		log.Fatalf("Failed to migrate tasks: %v", err)
	}
	return db, nil
}

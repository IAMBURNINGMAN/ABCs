package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Исправлена опечатка в пароле
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Database connected successfully!")
	return db, nil
}

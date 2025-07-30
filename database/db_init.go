package database

import (
	"fmt"
	"log"

	"github.com/myustanzah/GO-Fiber.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=root dbname=goDb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	// Migrate the schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	fmt.Println("Database connection established and schema migrated successfully")
	return db, nil

}

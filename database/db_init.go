package database

import (
	"fmt"
	"log"

	"github.com/myustanzah/GO-Fiber.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	
	var DB_HOST := "db"
       	var DB_PORT := "5432"
      	var DB_USER := "postgres"
      	var DB_PASSWORD := "root"
      	var DB_NAME = "goDb"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	// Migrate the schema
	if err := db.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Product{},
		&models.OrderItem{},
		&models.Category{},
	); err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	fmt.Println("Database connection established and schema migrated successfully")
	return db, nil

}

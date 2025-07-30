package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model              // This will automatically include ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name        string      `json:"name" gorm:"type:varchar(100);not null"`
	Description string      `json:"description" gorm:"type:text;not null"`
	Price       float64     `json:"price" gorm:"not null"`
	Stock       int         `json:"stock" gorm:"not null"`                   // Number of items in stock
	CategoryID  uint        `json:"category_id" gorm:"not null"`             // Foreign key to Category
	Category    Category    `json:"category" gorm:"foreignKey:CategoryID"`   // Relationship with Category
	OrderItems  []OrderItem `json:"order_items" gorm:"foreignKey:ProductID"` // One-to-many relationship with OrderItem
	ImageURL    string      `json:"image_url" gorm:"type:varchar(255);NULL"` // URL of the product image
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   time.Time   `json:"deleted_at" gorm:"autoDeleteTime"`
}

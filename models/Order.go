package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model             // This will automatically include ID, CreatedAt, UpdatedAt, DeletedAt fields
	CustomerID uint        `json:"customer_id" gorm:"not null"`           // Foreign key to User
	Customer   User        `json:"customer" gorm:"foreignKey:CustomerID"` // Relationship with User
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"` // One-to-many relationship with OrderItem
	TotalPrice float64     `json:"total_price" gorm:"not null"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

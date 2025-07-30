package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model            // This will automatically include ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	Products    []Product `json:"products" gorm:"foreignKey:CategoryID"` // One-to-many relationship with Product
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"autoDeleteTime"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // This will automatically include ID, CreatedAt, UpdatedAt, DeletedAt fields
	Name       string    `gorm:"type:varchar(100);not null" json:"name"`
	Email      string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Address    string    `gorm:"type:varchar(255);not null" json:"address"`
	Password   string    `gorm:"type:varchar(255);NULL" json:"password"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  time.Time `gorm:"autoDeleteTime" json:"deleted_at"`
	Orders     []Order   `gorm:"foreignKey:CustomerID" json:"orders"` // One-to-many relationship with Order
}

// func (User) TableName() string {
// 	return "users"
// }
// func (u *User) BeforeCreate() error {
// 	// You can add any logic here before creating a user, like hashing the password
// 	return nil
// }
// func (u *User) BeforeUpdate() error {
// 	// You can add any logic here before updating a user, like hashing the password
// 	return nil
// }
// func (u *User) BeforeDelete() error {
// 	// You can add any logic here before deleting a user, like logging the action
// 	return nil
// }
// func (u *User) AfterCreate() error {
// 	// You can add any logic here after creating a user, like sending a welcome email
// 	return nil
// }
// func (u *User) AfterUpdate() error {
// 	// You can add any logic here after updating a user, like sending a notification
// 	return nil
// }
// func (u *User) AfterDelete() error {
// 	// You can add any logic here after deleting a user, like logging the action
// 	return nil
// }

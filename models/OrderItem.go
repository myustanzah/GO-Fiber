package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model         // This will automatically include ID, CreatedAt, UpdatedAt, DeletedAt fields
	OrderID    uint    `json:"order_id" gorm:"not null"`            // Foreign key to Order
	Order      Order   `json:"order" gorm:"foreignKey:OrderID"`     // Relationship with Order
	ProductID  uint    `json:"product_id" gorm:"not null"`          // Foreign key to Product
	Product    Product `json:"product" gorm:"foreignKey:ProductID"` // Relationship with Product
	Quantity   int     `json:"quantity" gorm:"not null"`            // Quantity of the product in the order
	Price      float64 `json:"price" gorm:"not null"`               // Price of the product at the time of order
	Total      float64 `json:"total" gorm:"not null"`               // Total price for this item (Quantity * Price)
}

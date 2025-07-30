package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/controllers"
	"gorm.io/gorm"
)

func SetupOrderItemsRoutes(r fiber.Router, db *gorm.DB) {
	controller := controllers.NewOrderItemController(db)

	r.Get("/all", controller.GetAllOrderItems)          // Get all order items
	r.Get("/by-id/:id", controller.GetOrderItemByID)    // Get order item by ID
	r.Post("/create", controller.CreateOrderItem)       // Create a new order item
	r.Put("/update/:id", controller.UpdateOrderItem)    // Update order item by ID
	r.Delete("/delete/:id", controller.DeleteOrderItem) // Delete order item by ID
	// Add more order item-related routes as needed
}

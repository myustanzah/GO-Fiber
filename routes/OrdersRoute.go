package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/controllers"
	"gorm.io/gorm"
)

func SetupOrderRoutes(r fiber.Router, db *gorm.DB) {
	controller := controllers.NewOrderController(db)

	r.Get("/all", controller.GetAllOrders)          // Get all orders
	r.Get("/by-id/:id", controller.GetOrderByID)    // Get order by ID
	r.Post("/create", controller.CreateOrder)       // Create a new order
	r.Put("/update/:id", controller.UpdateOrder)    // Update order by ID
	r.Delete("/delete/:id", controller.DeleteOrder) // Delete order by ID
	// Add more order-related routes as needed
}

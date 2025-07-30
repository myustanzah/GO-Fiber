package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/controllers"
	"gorm.io/gorm"
)

func SetupProductRoutes(r fiber.Router, db *gorm.DB) {
	controller := controllers.NewProductController(db)

	r.Get("/all", controller.GetAllProducts)          // Get all products
	r.Get("/by-id/:id", controller.GetProductByID)    // Get product by ID
	r.Post("/create", controller.CreateProduct)       // Create a new product
	r.Put("/update/:id", controller.UpdateProduct)    // Update product by ID
	r.Delete("/delete/:id", controller.DeleteProduct) // Delete product by ID
	// Add more product-related routes as needed
}

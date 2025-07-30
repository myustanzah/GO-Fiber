package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/controllers"
	"gorm.io/gorm"
)

func SetupCategoryRoutes(r fiber.Router, db *gorm.DB) {
	controller := controllers.NewCategoryController(db)

	r.Get("/all", controller.GetAllCategories)         // Get all categories
	r.Get("/by-id/:id", controller.GetCategoryByID)    // Get category by ID
	r.Post("/create", controller.CreateCategory)       // Create a new category
	r.Put("/update/:id", controller.UpdateCategory)    // Update category by ID
	r.Delete("/delete/:id", controller.DeleteCategory) // Delete category by ID
	// Add more category-related routes as needed
}

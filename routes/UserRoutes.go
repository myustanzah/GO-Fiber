package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/controllers"
	"gorm.io/gorm"
)

func SetupUserRoutes(r fiber.Router, db *gorm.DB) {
	controller := controllers.NewUserController(db)

	r.Get("/all", controller.GetAllUsers)          // Get all users
	r.Get("/by-id/:id", controller.GetUserByID)    // Get user by ID
	r.Post("/create", controller.CreateUser)       // Create a new user
	r.Put("/update/:id", controller.UpdateUser)    // Update user by ID
	r.Delete("/delete/:id", controller.DeleteUser) // Delete user by ID
	// Add more user-related routes as needed
}

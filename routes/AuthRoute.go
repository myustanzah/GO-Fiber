package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/myustanzah/GO-Fiber.git/controllers"
	"gorm.io/gorm"
)

func SetupAuthRoutes(r fiber.Router, db *gorm.DB) {
	controller := controllers.NewAuthController(db)

	r.Post("/login", controller.Login)       // User login
	r.Post("/register", controller.Register) // User registration
	// Add more authentication-related routes as needed
}

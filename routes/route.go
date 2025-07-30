package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	r.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "Page not found")
	})

	r.Static("/static", "./public/src")
	r.Static("/error-page", "./public", fiber.Static{
		Index:         "error404.html",
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 3600, // 1 hour
	})
}

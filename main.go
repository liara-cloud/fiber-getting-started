package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Serve static files (CSS, JS, images)
	app.Static("/css", "static/css")
	app.Static("/js", "static/js")
	app.Static("/images", "static/images")

	// Serve the HTML file at the root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("static/index.html")
	})

	// Start the server on port 8080
	app.Listen(":8080")
}

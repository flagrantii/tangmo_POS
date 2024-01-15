// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define your routes and handlers here

	app.Listen(":3000")
}

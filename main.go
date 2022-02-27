package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/imanaspaul/googleouth/handler"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	// routes
	app.Get("/", handler.GoogleOauth)

	// app listen
	app.Listen(":3000")
}

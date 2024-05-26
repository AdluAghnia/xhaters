package main

import (
	"log"

	"github.com/AdluAghnia/xhater/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(logger.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":6969"))
}

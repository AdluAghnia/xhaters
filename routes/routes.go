package routes

import (
	"github.com/AdluAghnia/xhater/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// View
	app.Get("/", handler.Frontpagehandler)
	app.Get("/login", handler.LoginViewhandler)
	app.Get("/register", handler.RegisterViewhandler)
	app.Get("/info", handler.Myinformationhandler)

	// Login/Register & Session
	app.Get("/api/login", handler.HandlerLogin)
	app.Get("/api/register", handler.HandlerRegister)
	app.Get("/api/logout", handler.HandlerLogout)
}

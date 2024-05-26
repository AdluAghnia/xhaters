package handler

import "github.com/gofiber/fiber/v2"

func Frontpagehandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Page": "Home",
	}, "layouts/main")
}

func LoginViewhandler(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Page": "Login",
	}, "layouts/main")
}

func RegisterViewhandler(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Page": "Register",
	}, "layouts/main")
}

func Myinformationhandler(c *fiber.Ctx) error {
	return c.Render("myinfo", fiber.Map{
		"Page": "Ingfo",
	})
}

func HandlerLogin(c *fiber.Ctx) error {
	return c.SendString("LOGIN")
}

func HandlerRegister(c *fiber.Ctx) error {
	return c.SendString("REGISTER")
}

func HandlerLogout(c *fiber.Ctx) error {
	return c.SendString("Logout")
}

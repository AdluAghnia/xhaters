package handler

import (
	"log"
	"strings"

	"github.com/AdluAghnia/xhater/models"
	"github.com/AdluAghnia/xhater/session"
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

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

func HandlerRegister(c *fiber.Ctx) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	newUser := new(models.User)
	newUser.Email = email
	newUser.Username = username
	newUser.Password = password

	fm := fiber.Map{
		"error":        true,
		"errorMessage": "",
	}

	u, err := newUser.CreateUser()
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constrait") {
			fm["errorMessage"] = "This email already in use"
		} else {
			fm["errorMessage"] = "Something wrong"
		}

		return flash.WithError(c, fm).Redirect("/register")
	}

	err = session.CreateUserSession(c, u.ID)
	if err != nil {
		log.Println(err)
		fm["errorMessage"] = "Somethings has gone wrong: unable to register"

		return flash.WithError(c, fm).Redirect("/register")
	}
	fm = fiber.Map{
		"success":        true,
		"successMessage": "Successfully Registered !!!",
	}

	return flash.WithSuccess(c, fm).Redirect("/info")
}

func HandlerLogin(c *fiber.Ctx) error {
	return c.SendString("LOGIN")
}

func HandlerLogout(c *fiber.Ctx) error {
	return c.SendString("Logout")
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
	"log"
)

func HandlersRun() {
	engine := html.New("./Pages", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Use(requestid.New())
	app.Static("/", "./Pages")

	login := app.Group("/login")
	login.Get("", LoginPage)
	login.Post("", Login)

	signUp := app.Group("/signup")
	signUp.Get("", SignUpPage)
	signUp.Post("", SignUpHandler)

	log.Fatal(app.Listen(":8080"))
}

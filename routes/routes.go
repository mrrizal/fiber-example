package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/book"
	"github.com/mrrizal/fiber-example/user"
)

func SetupUserRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Post("/user/sign-up", user.SignUpHandler)
	v1.Post("/user/login", user.LoginHandler)
}
func SetupRoutes(app *fiber.App, s book.Service) {
	SetupUserRoutes(app)
	book.SetupBookRoutes(app, s)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/book"
	"github.com/mrrizal/fiber-example/user"
)

func SetupRoutes(app *fiber.App, bookService book.Service, userService user.Service) {
	book.SetupBookRoutes(app, bookService)
	user.SetupUserRoutes(app, userService)
}

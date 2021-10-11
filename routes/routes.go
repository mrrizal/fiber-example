package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/book"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/book", book.GetBooksHandler)
	v1.Get("/book/:id", book.GetBookHandler)
	v1.Post("/book", book.NewBookHandler)
	v1.Delete("/book/:id", book.DeleteBookHandler)
}

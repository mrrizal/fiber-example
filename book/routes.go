package book

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/mrrizal/fiber-example/configs"
)

func SetupBookRoutes(app *fiber.App, s Service) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	book := v1.Group("/book")

	book.Get("/", func(c *fiber.Ctx) error {
		return GetBooksHandler(c, s)
	})

	book.Get("/:id", func(c *fiber.Ctx) error {
		return GetBookHandler(c, s)
	})

	book.Use(jwtware.New(jwtware.Config{SigningKey: []byte(configs.Configs.SecretKey)}))

	book.Post("/", func(c *fiber.Ctx) error {
		return NewBookHandler(c, s)
	})

	book.Delete("/:id", func(c *fiber.Ctx) error {
		return DeleteBookHandler(c, s)
	})
}

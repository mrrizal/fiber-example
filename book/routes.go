package book

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/mrrizal/fiber-example/configs"
)

func SetupBookRoutes(app *fiber.App, s Service) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/book", func(c *fiber.Ctx) error {
		return GetBooksHandler(c, s)
	})

	v1.Get("/book/:id", func(c *fiber.Ctx) error {
		return GetBookHandler(c, s)
	})

	v1.Use(jwtware.New(jwtware.Config{SigningKey: []byte(configs.Configs.SecretKey)}))

	v1.Post("/book", func(c *fiber.Ctx) error {
		return NewBookHandler(c, s)
	})

	v1.Delete("/book/:id", func(c *fiber.Ctx) error {
		return DeleteBookHandler(c, s)
	})
}

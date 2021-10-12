package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/mrrizal/fiber-example/book"
	"github.com/mrrizal/fiber-example/configs"
	"github.com/mrrizal/fiber-example/user"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/book", book.GetBooksHandler)
	v1.Get("/book/:id", book.GetBookHandler)

	v1.Post("/user/sign-up", user.SignUpHandler)
	v1.Post("/user/login", user.LoginHandler)

	v1.Use(jwtware.New(jwtware.Config{SigningKey: []byte(configs.Configs.SecretKey)}))

	// you need jwt
	v1.Post("/book", book.NewBookHandler)
	v1.Delete("/book/:id", book.DeleteBookHandler)
}

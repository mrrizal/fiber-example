package user

import "github.com/gofiber/fiber/v2"

func SetupUserRoutes(app *fiber.App, service Service) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	user := v1.Group("/user")

	user.Post("/sign-up", func(c *fiber.Ctx) error {
		return SignUpHandler(c, service)
	})

	user.Post("/login", func(c *fiber.Ctx) error {
		return LoginHandler(c, service)
	})
}

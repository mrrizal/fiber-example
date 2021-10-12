package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/utils"
)

func SignUpHandler(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	if err := SingUp(&user); err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	c.Status(201)
	return c.JSON(UserResponse(user))
}

func LoginHandler(c *fiber.Ctx) error {
	userCredentials := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&userCredentials); err != nil {
		return utils.ErrorResponse(c, 400, err)
	}

	if err := Login(userCredentials.Username, userCredentials.Password); err != nil {
		return utils.ErrorResponse(c, 400, err)
	}

	response := struct {
		Message string `json:"message"`
	}{}
	response.Message = "Login success"
	c.Status(200)
	return c.JSON(response)
}

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

	if err := singUp(&user); err != nil {
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

	user, err := login(userCredentials.Username, userCredentials.Password)
	if err != nil {
		return utils.ErrorResponse(c, 400, err)
	}

	token, err := generateJWTToken(user)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	response := struct {
		Token string `json:"token"`
	}{}
	response.Token = token
	c.Status(200)
	return c.JSON(response)
}

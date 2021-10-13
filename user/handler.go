package user

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/utils"
)

func SignUpHandler(c *fiber.Ctx, s Service) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return utils.ErrorResponse(c, 400, err)
	}

	if user.Password == "" {
		return utils.ErrorResponse(c, 400, errors.New("password cannot be empty"))
	}

	if err := s.singUp(&user); err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	c.Status(201)
	return c.JSON(UserResponse(user))
}

func LoginHandler(c *fiber.Ctx, s Service) error {
	userCredentials := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&userCredentials); err != nil {
		return utils.ErrorResponse(c, 400, err)
	}

	if userCredentials.Username == "" || userCredentials.Password == "" {
		return utils.ErrorResponse(c, 400, errors.New("username and password is required"))
	}

	user, err := s.login(userCredentials.Username, userCredentials.Password)
	if err != nil {
		return utils.ErrorResponse(c, 400, err)
	}

	token, err := GenerateJWTToken(user)
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

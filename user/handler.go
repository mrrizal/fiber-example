package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizal/fiber-example/database"
	"github.com/mrrizal/fiber-example/utils"
)

func SignUpHandler(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		c.Status(400)
		return c.JSON(err.Error())
	}

	hash := utils.Hash{}
	generatedPassword, err := hash.Generate(user.Password)
	if err != nil {
		c.Status(500)
		return c.JSON(err.Error())
	}
	user.Password = generatedPassword

	if err := database.DBConn.Create(&user).Error; err != nil {
		c.Status(500)
		return c.JSON(err.Error())
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
		c.Status(400)
		return c.JSON(err.Error())
	}

	if err := Login(userCredentials.Username, userCredentials.Password); err != nil {
		c.Status(400)
		return c.JSON(err.Error())
	}

	return c.JSON(userCredentials)
}

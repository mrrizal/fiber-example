package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/database"
	"go.elastic.co/apm"
	"strings"
)

type Service interface {
	singUp(c *fiber.Ctx, user *User) error
	login(c *fiber.Ctx, username, password string) (User, error)
}

type ServiceStruct struct{}

func (s *ServiceStruct) singUp(c *fiber.Ctx, user *User) error {
	span, _ := apm.StartSpan(c.Context(), "singUp", "postgres")
	defer span.End()
	hash := hash{}
	generatedPassword, err := hash.generatePassword(c, user.Password)
	if err != nil {
		return err
	}
	user.Password = generatedPassword

	user.Username = strings.ToLower(user.Username)
	if err := database.DBConn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s *ServiceStruct) login(c *fiber.Ctx, username, password string) (User, error) {
	span, _ := apm.StartSpan(c.Context(), "login", "postgres")
	defer span.End()
	var user User
	if err := database.DBConn.Where("username = ?", strings.ToLower(username)).First(&user).Error; err != nil {
		return User{}, err
	}

	hash := hash{}
	err := hash.comparePassword(c, user.Password, password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

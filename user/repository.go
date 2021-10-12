package user

import (
	"github.com/mrrizal/fiber-example/database"
	"github.com/mrrizal/fiber-example/utils"
)

func SingUp(user *User) error {
	hash := utils.Hash{}
	generatedPassword, err := hash.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = generatedPassword

	if err := database.DBConn.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

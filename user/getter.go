package user

import (
	"github.com/mrrizal/fiber-example/database"
	"github.com/mrrizal/fiber-example/utils"
)

func Login(username, password string) error {
	var user User
	if err := database.DBConn.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}

	hash := utils.Hash{}
	err := hash.Compare(user.Password, password)
	if err != nil {
		return err
	}
	return nil
}

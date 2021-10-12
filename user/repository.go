package user

import (
	"github.com/mrrizal/fiber-example/database"
	"strings"
)

func singUp(user *User) error {
	hash := hash{}
	generatedPassword, err := hash.generate(user.Password)
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

func login(username, password string) (User, error) {
	var user User
	if err := database.DBConn.Where("username = ?", strings.ToLower(username)).First(&user).Error; err != nil {
		return User{}, err
	}

	hash := hash{}
	err := hash.compare(user.Password, password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

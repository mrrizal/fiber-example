package user

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/mrrizal/fiber-example/configs"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type hash struct{}

func (c *hash) generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes)
	return hash, nil
}

func (c *hash) compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func GenerateJWTToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, err := token.SignedString([]byte(configs.Configs.SecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

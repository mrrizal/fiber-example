package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mrrizal/fiber-example/configs"
	"go.elastic.co/apm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type hash struct{}

func (h *hash) generatePassword(c *fiber.Ctx, s string) (string, error) {
	span, _ := apm.StartSpan(c.Context(), "generatePassword", "utils")
	defer span.End()
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes)
	return hash, nil
}

func (h *hash) comparePassword(c *fiber.Ctx, hash string, s string) error {
	span, _ := apm.StartSpan(c.Context(), "comparePassword", "utils")
	defer span.End()
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func GenerateJWTToken(c *fiber.Ctx, user User) (string, error) {
	span, _ := apm.StartSpan(c.Context(), "GenerateJWTToken", "utils")
	defer span.End()
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

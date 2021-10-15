package utils

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.elastic.co/apm"
)

func ErrorResponse(c *fiber.Ctx, statusCode int, err error) error {
	handleAPMError(err, c.Context())
	c.Status(statusCode)
	resp := struct {
		Message string `json:"message"`
	}{Message: err.Error()}
	return c.JSON(resp)
}

func ValidateJWTToken(c *fiber.Ctx) (float64, error) {
	err := errors.New("Invalid jwt token")
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return 0, err
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	if _, ok := claims["id"].(float64); !ok {
		return 0, err
	}

	return claims["id"].(float64), nil
}

func handleAPMError(err error, ctx context.Context) {
	e := apm.CaptureError(ctx, err)
	e.Send()
}

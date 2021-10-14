package user

import (
	"bytes"
	"encoding/json"
	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/configs"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

type userServiceStructFake struct {
}

func (s *userServiceStructFake) singUp(c *fiber.Ctx, user *User) error {
	hash := hash{}
	generatedPassword, err := hash.generatePassword(c, user.Password)
	if err != nil {
		return err
	}
	user.Password = generatedPassword

	user.Username = strings.ToLower(user.Username)
	return nil
}

func (s *userServiceStructFake) login(c *fiber.Ctx, username, password string) (User, error) {
	var user User
	err := faker.FakeData(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func setup(service Service) *fiber.App {
	app := fiber.New()
	SetupUserRoutes(app, service)
	return app
}

func TestMain(m *testing.M) {
	os.Setenv("PAGE_SIZE", "10")
	os.Setenv("SECRET_KEY", "secret")
	configs.GetSettings()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestSignUpHandler(t *testing.T) {
	userService := &userServiceStructFake{}
	app := setup(userService)

	tests := []struct {
		name           string
		requestPayload []byte
		statusCode     int
	}{
		{name: "Invalid Payload", requestPayload: []byte(``), statusCode: 400},
		{name: "Without Password", requestPayload: []byte(`{"username": "test"}`), statusCode: 400},
		{name: "Correct Payload", requestPayload:
		[]byte(`{"username": "test", "password": "test", "first_name": "test", "last_name": "test"}`), statusCode: 201},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/api/v1/user/sign-up",
				bytes.NewBuffer(test.requestPayload))
			request.Header.Set("Content-Type", "application/json; charset=UTF-8")
			response, err := app.Test(request)
			if err != nil {
				t.Error(err.Error())
			}
			assert.Equal(t, test.statusCode, response.StatusCode)
			if test.statusCode == 201 {
				respBody, _ := ioutil.ReadAll(response.Body)
				userResp := User{}
				if err := json.Unmarshal(respBody, &userResp); err != nil {
					t.Error(err.Error())
				}
				assert.Equal(t, User{Username: "test", FirstName: "test", LastName: "test"}, userResp)
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	userService := &userServiceStructFake{}
	app := setup(userService)

	tests := []struct {
		name           string
		requestPayload []byte
		statusCode     int
	}{
		{name: "Invalid Payload", requestPayload: []byte(``), statusCode: 400},
		{name: "Without Password", requestPayload: []byte(`{"username": "test"}`), statusCode: 400},
		{name: "Without Username", requestPayload: []byte(`{"password": "test"}`), statusCode: 400},
		{name: "Correct Payload", requestPayload: []byte(`{"username": "test", "password": "test"}`), statusCode: 200},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/api/v1/user/login",
				bytes.NewBuffer(test.requestPayload))
			request.Header.Set("Content-Type", "application/json; charset=UTF-8")
			response, err := app.Test(request)
			if err != nil {
				t.Error(err.Error())
			}
			assert.Equal(t, test.statusCode, response.StatusCode)
		})
	}
}

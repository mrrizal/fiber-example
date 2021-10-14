package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mrrizal/fiber-example/user"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/configs"
	"github.com/stretchr/testify/assert"
)

type bookServiceStructFake struct {
}

func (s *bookServiceStructFake) getBooks(c *fiber.Ctx, id int, previousPage bool) ([]Book, error) {
	var books []Book
	for i := 0; i < configs.Configs.PageSize; i++ {
		var book Book
		err := faker.FakeData(&book)
		if err != nil {
			return books, err
		}
		book.ID = uint(i)
		books = append(books, book)
	}
	return books, nil
}

func (s *bookServiceStructFake) getBook(c *fiber.Ctx, id int) (Book, error) {
	var book Book
	err := faker.FakeData(&book)
	book.AuthorID = 1
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s *bookServiceStructFake) newBook(c *fiber.Ctx, book *Book) error {
	return nil
}

func (s *bookServiceStructFake) deleteBook(c *fiber.Ctx, book *Book) error {
	return nil
}

func setup(service Service) *fiber.App {
	app := fiber.New()
	SetupBookRoutes(app, service)
	return app
}

func TestGetBooksHandler(t *testing.T) {
	bookService := &bookServiceStructFake{}
	app := setup(bookService)

	request := httptest.NewRequest(http.MethodGet, "/api/v1/book/", bytes.NewReader(nil))
	response, err := app.Test(request)
	if err != nil {
		t.Error(err.Error())
	}
	body, _ := ioutil.ReadAll(response.Body)
	respJSON := struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []Book `json:"results"`
	}{}

	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, configs.Configs.PageSize, len(respJSON.Results))
	assert.NotEqual(t, "", respJSON.Next)
	assert.NotEqual(t, "", respJSON.Previous)
	nextUrl := fmt.Sprintf("http://example.com/api/v1/book/?next=%d", respJSON.Results[len(respJSON.Results)-1].ID)
	previousURL := fmt.Sprintf("http://example.com/api/v1/book/?previous=%d", respJSON.Results[0].ID)
	assert.Equal(t, nextUrl, respJSON.Next)
	assert.Equal(t, previousURL, respJSON.Previous)
}

func TestGetBookHandler(t *testing.T) {
	bookService := &bookServiceStructFake{}
	app := setup(bookService)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/book/%d", 1), bytes.NewReader(nil))
	response, err := app.Test(request)
	if err != nil {
		t.Error(err.Error())
	}
	body, _ := ioutil.ReadAll(response.Body)
	var respJSON Book

	err = json.Unmarshal(body, &respJSON)
	if err != nil {
		t.Error(err.Error())
	}
	assert.Equal(t, 200, response.StatusCode)
}

func TestNewBookHandler(t *testing.T) {
	bookService := &bookServiceStructFake{}
	app := setup(bookService)

	requestBody := []byte(`{"title": "test"}`)
	t.Run("without auth", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/api/v1/book/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		response, err := app.Test(request)
		if err != nil {
			t.Error(err.Error())
		}
		assert.NotEqual(t, 200, response.StatusCode)
	})

	t.Run("with auth", func(t *testing.T) {
		var testUser user.User
		testUser.ID = 1
		testUser.Username = "test"

		token, _ := user.GenerateJWTToken(&fiber.Ctx{}, testUser)
		request := httptest.NewRequest(http.MethodPost, "/api/v1/book/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
		response, err := app.Test(request)
		if err != nil {
			t.Error(err.Error())
		}

		responseJSON, _ := ioutil.ReadAll(response.Body)
		bookResponse := Book{}
		err = json.Unmarshal(responseJSON, &bookResponse)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Equal(t, 201, response.StatusCode)
		assert.Equal(t, uint(0), bookResponse.ID)
		assert.Equal(t, testUser.ID, bookResponse.AuthorID)
	})

	t.Run("invalid payload", func(t *testing.T) {
		var testUser user.User
		testUser.ID = 1
		testUser.Username = "test"

		token, _ := user.GenerateJWTToken(&fiber.Ctx{}, testUser)
		requestBody := []byte(`{"test": "test"}`)
		request := httptest.NewRequest(http.MethodPost, "/api/v1/book/", bytes.NewBuffer(requestBody))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		request.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
		response, err := app.Test(request)
		if err != nil {
			t.Error(err.Error())
		}

		responseJSON, _ := ioutil.ReadAll(response.Body)
		bookResponse := Book{}
		err = json.Unmarshal(responseJSON, &bookResponse)
		if err != nil {
			t.Error(err.Error())
		}

		assert.NotEqual(t, 200, response.StatusCode)
	})
}

func TestDeleteBookHandler(t *testing.T) {
	bookService := &bookServiceStructFake{}
	app := setup(bookService)

	t.Run("without auth", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/book/1", nil)
		response, err := app.Test(request)
		if err != nil {
			t.Error(err.Error())
		}
		assert.NotEqual(t, 200, response.StatusCode)
	})

	t.Run("with auth", func(t *testing.T) {
		var testUser user.User
		testUser.ID = 1
		testUser.Username = "test"

		token, _ := user.GenerateJWTToken(&fiber.Ctx{}, testUser)
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/book/1", nil)
		request.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
		response, err := app.Test(request)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Equal(t, 204, response.StatusCode)
	})

	t.Run("author id != user id", func(t *testing.T) {
		var testUser user.User
		testUser.ID = 2
		testUser.Username = "test"

		token, _ := user.GenerateJWTToken(&fiber.Ctx{}, testUser)
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/book/1", nil)
		request.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
		response, err := app.Test(request)
		if err != nil {
			t.Error(err.Error())
		}

		assert.Equal(t, 403, response.StatusCode)
	})
}

func TestMain(m *testing.M) {
	os.Setenv("PAGE_SIZE", "10")
	os.Setenv("SECRET_KEY", "secret")
	configs.GetSettings()
	exitVal := m.Run()
	os.Exit(exitVal)
}

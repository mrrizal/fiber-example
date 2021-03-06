package book

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/configs"
	"go.elastic.co/apm"
)

type booksResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Book `json:"results"`
}

func getNextURL(c *fiber.Ctx, books []Book) string {
	span, _ := apm.StartSpan(c.Context(), "getNextURL", "parser")
	defer span.End()
	if len(books) == 0 || len(books) < configs.Configs.PageSize {
		return ""
	}
	nextURL := fmt.Sprintf("%s%s?next=%d", c.BaseURL(), c.Path(), books[len(books)-1].ID)
	return nextURL
}

func getPreviousURL(c *fiber.Ctx, books []Book) string {
	span, _ := apm.StartSpan(c.Context(), "getPreviousURL", "parser")
	defer span.End()
	if len(books) == 0 {
		return ""
	}
	nextURL := fmt.Sprintf("%s%s?previous=%d", c.BaseURL(), c.Path(), books[0].ID)
	return nextURL
}

func booksToResponse(c *fiber.Ctx, books []Book) booksResponse {
	span, _ := apm.StartSpan(c.Context(), "booksToResponse", "parser")
	defer span.End()
	resp := booksResponse{Results: books, Next: getNextURL(c, books), Previous: getPreviousURL(c, books)}
	return resp
}

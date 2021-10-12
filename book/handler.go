package book

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/utils"
)

func GetBooksHandler(c *fiber.Ctx) error {
	var id int
	next := c.Query("next")
	previous := c.Query("previous")
	previousPage := false

	id, err := utils.GetIDFromURLQuery(next, previous, &previousPage)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	books, err := getBooks(id, previousPage)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	booksResponse := booksToResponse(c, books)
	return c.JSON(booksResponse)
}

func GetBookHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	book, err := getBook(id)
	if err != nil {
		return utils.ErrorResponse(c, 404, err)
	}
	return c.JSON(book)
}

func NewBookHandler(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return utils.ErrorResponse(c, 400, err)
	}

	if err := newBook(book); err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	c.Status(201)
	return c.JSON(book)
}

func DeleteBookHandler(c *fiber.Ctx) error {
	var book Book
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	if book, err = getBook(id); err != nil {
		return utils.ErrorResponse(c, 404, err)
	}

	if err := deleteBook(&book); err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	c.Status(204)
	return c.JSON(nil)
}

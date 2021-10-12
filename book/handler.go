package book

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/utils"
)

func GetBooksHandler(c *fiber.Ctx) error {
	var id int
	next := c.Query("next")
	previous := c.Query("previous")
	previousPage := false

	id, err := getIDFromURLQuery(next, previous, &previousPage)
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

	authorID, err := utils.ValidateJWTToken(c)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	book.AuthorID = uint(authorID)
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

	authorID, err := utils.ValidateJWTToken(c)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	if uint(authorID) != book.AuthorID {
		return utils.ErrorResponse(c, 403, errors.New("You don't have access to do this action."))
	}

	if err := deleteBook(&book); err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	c.Status(204)
	return c.JSON(nil)
}

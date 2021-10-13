package book

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/utils"
)

type BookService interface {
}

func GetBooksHandler(c *fiber.Ctx, s Service) error {
	var id int
	next := c.Query("next")
	previous := c.Query("previous")
	previousPage := false

	id, err := getIDFromURLQuery(next, previous, &previousPage)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	books, err := s.getBooks(id, previousPage)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	booksResponse := booksToResponse(c, books)
	return c.JSON(booksResponse)
}

func GetBookHandler(c *fiber.Ctx, s Service) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	book, err := s.getBook(id)
	if err != nil {
		return utils.ErrorResponse(c, 404, err)
	}
	return c.JSON(book)
}

func NewBookHandler(c *fiber.Ctx, s Service) error {
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		return utils.ErrorResponse(c, 400, err)
	}

	authorID, err := utils.ValidateJWTToken(c)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	book.AuthorID = uint(authorID)
	if err := s.newBook(book); err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	c.Status(201)
	return c.JSON(book)
}

func DeleteBookHandler(c *fiber.Ctx, s Service) error {
	var book Book
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	if book, err = s.getBook(id); err != nil {
		return utils.ErrorResponse(c, 404, err)
	}

	authorID, err := utils.ValidateJWTToken(c)
	if err != nil {
		return utils.ErrorResponse(c, 500, err)
	}

	if uint(authorID) != book.AuthorID {
		return utils.ErrorResponse(c, 403, errors.New("You don't have access to do this action."))
	}

	if err := s.deleteBook(&book); err != nil {
		return utils.ErrorResponse(c, 500, err)
	}
	c.Status(204)
	return c.JSON(nil)
}

package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/database"
	"github.com/mrrizal/fiber-example/utils"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func GetBooksHandler(c *fiber.Ctx) error {
	var id int
	next := c.Query("next")
	previous := c.Query("previous")
	previousPage := false

	id, err := utils.GetIDFromURLQuery(next, previous, &previousPage)
	if err != nil {
		log.Error(err.Error())
		c.Status(500)
		return nil
	}

	books, err := getBooks(id, previousPage)
	if err != nil {
		log.Error(err)
		return c.JSON([]Book{})
	}
	booksResponse := booksToResponse(c, books)
	return c.JSON(booksResponse)
}

func GetBookHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error(err.Error())
		c.Status(500)
		return c.JSON(nil)
	}

	book, err := getBook(id)
	if err != nil {
		log.Error(err)
		c.Status(404)
		return c.JSON(nil)
	}
	return c.JSON(book)
}

func NewBookHandler(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(400)
		return c.JSON(err.Error())
	}

	if err := database.DBConn.Create(&book).Error; err != nil {
		c.Status(500)
		return c.JSON(err.Error())
	}
	c.Status(201)
	return c.JSON(book)
}

func DeleteBookHandler(c *fiber.Ctx) error {
	var book Book
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Error(err.Error())
		c.Status(500)
		return c.JSON(nil)
	}

	if err := database.DBConn.First(&book, id).Error; err != nil {
		log.Error(err.Error())
		c.Status(404)
		return c.JSON(nil)
	}

	if err := database.DBConn.Delete(&book).Error; err != nil {
		log.Error(err.Error())
		c.Status(500)
		return c.JSON(nil)
	}
	c.Status(204)
	return c.JSON(nil)
}

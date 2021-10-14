package book

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizal/fiber-example/configs"
	"github.com/mrrizal/fiber-example/database"
	"go.elastic.co/apm"
)

type Service interface {
	getBooks(c *fiber.Ctx, id int, previousPage bool) ([]Book, error)
	getBook(c *fiber.Ctx, id int) (Book, error)
	newBook(c *fiber.Ctx, book *Book) error
	deleteBook(c *fiber.Ctx, book *Book) error
}

type ServiceStruct struct{}

func booksQueryBuilder(c *fiber.Ctx, id int, previousPage bool, tableName string) string {
	span, _ := apm.StartSpan(c.Context(), "booksQueryBuilder", "postgres")
	defer span.End()
	query := ""
	operator := "<"
	order := "desc"
	if id == 0 {
		query = fmt.Sprintf("select * from %s where deleted_at is null order by id desc limit %d", tableName,
			configs.Configs.PageSize)
	} else {
		if previousPage {
			operator = ">"
			order = "asc"
		}
		query = fmt.Sprintf("select * from %s where id %s %d and deleted_at is null order by id %s limit %d",
			tableName, operator, id, order, configs.Configs.PageSize)
	}
	return query
}

func (s *ServiceStruct) getBooks(c *fiber.Ctx, id int, previousPage bool) ([]Book, error) {
	span, _ := apm.StartSpan(c.Context(), "getBooks", "postgres")
	defer span.End()
	var books []Book

	query := booksQueryBuilder(c, id, previousPage, "books")
	if err := database.DBConn.Raw(query).Scan(&books).Error; err != nil {
		return books, err
	}

	if previousPage {
		var previousBooks []Book
		for i := len(books) - 1; i >= 0; i-- {
			previousBooks = append(previousBooks, books[i])
		}
		return previousBooks, nil
	}
	return books, nil
}

func (s *ServiceStruct) getBook(c *fiber.Ctx, id int) (Book, error) {
	span, _ := apm.StartSpan(c.Context(), "getBook", "postgres")
	defer span.End()
	var book Book
	if err := database.DBConn.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (s *ServiceStruct) newBook(c *fiber.Ctx, book *Book) error {
	span, _ := apm.StartSpan(c.Context(), "newBook", "postgres")
	defer span.End()
	if err := database.DBConn.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func (s *ServiceStruct) deleteBook(c *fiber.Ctx, book *Book) error {
	span, _ := apm.StartSpan(c.Context(), "deleteBook", "postgres")
	defer span.End()
	if err := database.DBConn.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

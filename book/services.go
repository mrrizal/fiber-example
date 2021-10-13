package book

import (
	"fmt"
	"github.com/mrrizal/fiber-example/configs"
	"github.com/mrrizal/fiber-example/database"
)

type Service interface {
	getBooks(id int, previousPage bool) ([]Book, error)
	getBook(id int) (Book, error)
	newBook(book *Book) error
	deleteBook(book *Book) error
}

type ServiceStruct struct{}

func booksQueryBuilder(id int, previousPage bool, tableName string) string {
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

func (s *ServiceStruct) getBooks(id int, previousPage bool) ([]Book, error) {
	var books []Book

	query := booksQueryBuilder(id, previousPage, "books")
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

func (s *ServiceStruct) getBook(id int) (Book, error) {
	var book Book
	if err := database.DBConn.Where("id = ?", id).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (s *ServiceStruct) newBook(book *Book) error {
	if err := database.DBConn.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func (s *ServiceStruct) deleteBook(book *Book) error {
	if err := database.DBConn.Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

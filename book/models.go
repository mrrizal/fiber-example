package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" gorm:"uniqueIndex:books_title;not null;default:null"`
	Author string `json:"author" gorm:"uniqueIndex:books_title;not null;default:null"`
	Rating int    `json:"rating" gorm:"default=0"`
}

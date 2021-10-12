package book

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title     string         `json:"title" gorm:"uniqueIndex:books_title;not null;default:null"`
	Author    string         `json:"author" gorm:"uniqueIndex:books_title;not null;default:null"`
	Rating    int            `json:"rating" gorm:"default=0"`
}

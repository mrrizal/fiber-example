package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username      string         `json:"username" gorm:"unique;not null;default:null"`
	Password      string         `json:"password" gorm:"not null;default:null"`
	FirstName     string         `json:"first_name" gorm:"not null;default:null"`
	LastName      string         `json:"last_name" gorm:"not null;default:null"`
	ProfilePicURL string         `json:"profile_pic_url"`
}

type UserResponse struct {
	ID            uint           `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username      string         `json:"username" gorm:"unique;not null;default:null"`
	Password      string         `json:"-" gorm:"not null;default:null"`
	FirstName     string         `json:"first_name" gorm:"not null;default:null"`
	LastName      string         `json:"last_name" gorm:"not null;default:null"`
	ProfilePicURL string         `json:"profile_pic_url"`
}

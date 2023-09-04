package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var CustomValidator *validator.Validate

func init() {
	CustomValidator = validator.New()
}

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"not null" validate:"required"`
	Email       string `gorm:"unique;not null" validate:"required,email"`
	Password    string `gorm:"not null" validate:"required,min=6"`
	Age         int    `gorm:"not null" validate:"required,gte=9"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Photos      []Photo       `gorm:"foreignKey:UserID"` // One-to-Many relationship with Photo
	Comments    []Comment     `gorm:"foreignKey:UserID"` // One-to-Many relationship with Comment
	SocialMedia []SocialMedia `gorm:"foreignKey:UserID"` // One-to-Many relationship with SocialMedia
}

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null" validate:"required"`
	Caption   string
	PhotoURL  string `gorm:"not null" validate:"required"`
	UserID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	PhotoID   uint   `gorm:"not null"`
	Message   string `gorm:"not null" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"not null" validate:"required"`
	SocialMediaURL string `gorm:"not null" validate:"required"`
	UserID         uint   `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

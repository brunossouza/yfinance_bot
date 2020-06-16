package models

import (
	"github.com/jinzhu/gorm"
)

// Usuario dados do usuario do telegram
type Usuario struct {
	gorm.Model
	UserID   int
	Nome     string
	Username string
}

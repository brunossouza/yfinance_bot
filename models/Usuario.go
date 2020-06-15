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
	Acoes    []Acoes `gorm:"foreignkey:UsuarioRefer"`
}

// Acoes de cada usuario
type Acoes struct {
	gorm.Model
	Ticker       string
	Variacao     float64
	UsuarioRefer string
}

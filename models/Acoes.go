package models

import "github.com/jinzhu/gorm"

// Acoes de cada usuario
type Acoes struct {
	gorm.Model
	Ticker      string
	Variacao    float64
	ValorMinimo float64
	ValorMaximo float64
	UserID      int
}

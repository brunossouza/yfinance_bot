package models

import (
	"github.com/jinzhu/gorm"
)

// TelegramBotConfig configuração do bot
type TelegramBotConfig struct {
	gorm.Model
	Token string
}

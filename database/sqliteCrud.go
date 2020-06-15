package database

import (
	"yfinance/models"
	"yfinance/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// GetConexao return a conexão do banco
func GetConexao() (db *gorm.DB) {
	db, err := gorm.Open("sqlite3", "database.db")
	utils.CheckError(err)
	return
}

// SaveBotConfig func para salvar a configuração do bot
func SaveBotConfig(config *models.TelegramBotConfig) {
	db := GetConexao()
	defer db.Close()

	// Delete table
	db.DropTableIfExists(&models.TelegramBotConfig{})

	// Migrate the schema
	db.AutoMigrate(&models.TelegramBotConfig{})

	// Create
	db.Create(config)
}

// SaveUsuario func para salvar os dados do usuário
func SaveUsuario(user *models.Usuario) {
	db := GetConexao()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Usuario{})

	// Create
	db.Create(user)
}

// GetUsuario func para salvar os dados do usuário
func GetUsuario(userID int) (user models.Usuario) {
	db := GetConexao()
	defer db.Close()

	db.Where("user_id = ?", userID).First(&user)

	return
}

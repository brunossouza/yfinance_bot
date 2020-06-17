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

// MigrateDatabaseSchema func to create the database
func MigrateDatabaseSchema() {
	db := GetConexao()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.TelegramBotConfig{})
	db.AutoMigrate(&models.Usuario{})
	db.AutoMigrate(&models.Acoes{})
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

// GetTelegramBotConfig pega a configuração do token
func GetTelegramBotConfig() (config models.TelegramBotConfig) {
	db := GetConexao()
	defer db.Close()

	db.First(&config)

	return
}

// SaveUsuario func para salvar os dados do usuário
func SaveUsuario(user *models.Usuario) {
	db := GetConexao()
	defer db.Close()

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

// HasUsuarioRegistred func para salvar os dados do usuário
func HasUsuarioRegistred(userID int) bool {
	db := GetConexao()
	defer db.Close()

	var user models.Usuario
	db.Where("user_id = ?", userID).First(&user)

	return user.UserID != 0
}

// SaveAcao func para salvar os dados da ação
func SaveAcao(acao *models.Acoes) {
	db := GetConexao()
	defer db.Close()

	// Create
	db.Create(acao)
}

// GetAllAcoes func para salvar os dados do usuário
func GetAllAcoes() (acoes []models.Acoes) {
	db := GetConexao()
	defer db.Close()

	db.Find(&acoes)

	return
}

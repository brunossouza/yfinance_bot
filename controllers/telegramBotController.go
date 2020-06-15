package controllers

import (
	"strings"
	"time"
	"yfinance/database"
	"yfinance/models"
	"yfinance/utils"

	tgbot "gopkg.in/tucnak/telebot.v2"
)

// TelegramBotHandler controlle de comandos do bot
func TelegramBotHandler() {

	config := database.GetTelegramBotConfig()

	bot, err := tgbot.NewBot(tgbot.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		// URL: "http://195.129.111.17:8012",

		Token:  config.Token,
		Poller: &tgbot.LongPoller{Timeout: 10 * time.Second},
	})
	utils.CheckError(err)

	bot.Handle("/help", func(m *tgbot.Message) {
		bot.Send(m.Sender, "/help ajuda\n\n/start iniciar bot\n/list lista sua ações")
	})

	bot.Handle("/list", func(m *tgbot.Message) {
		var user = database.GetUsuario(m.Sender.ID)
		bot.Send(m.Sender, user.UserID)
	})

	bot.Handle("/check", func(m *tgbot.Message) {
		if m.Payload != "" {
			payload := m.Payload
			var tickers = strings.Split(payload, " ")
			for _, ticker := range tickers {
				var msg = GetQuoteDataAsMessage(strings.ToUpper(ticker))
				bot.Send(m.Sender, msg)
			}
		} else {
			var msg = "usage: /check ticker ticker1 ... tickerN"
			bot.Send(m.Sender, msg)
		}
	})

	bot.Handle("/start", func(m *tgbot.Message) {
		var user = models.Usuario{Nome: m.Sender.FirstName, Username: m.Sender.Username, UserID: m.Sender.ID}
		database.SaveUsuario(&user)
		bot.Send(m.Sender, "Usuário cadastrado")
	})

	bot.Start()
}

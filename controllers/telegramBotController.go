package controllers

import (
	"strconv"
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
		var helpMsg = "/help - ajuda\n\n" +
			"/start - registra seu usuário no bot\n" +
			"/status - verificar uma cotação\n"

		bot.Send(m.Sender, helpMsg)
	})

	bot.Handle("/start", func(m *tgbot.Message) {
		if !database.HasUsuarioRegistred(m.Sender.ID) {
			var user = models.Usuario{Nome: m.Sender.FirstName, Username: m.Sender.Username, UserID: m.Sender.ID}
			database.SaveUsuario(&user)
			bot.Send(m.Sender, "Usuário cadastrado.")
		} else {
			bot.Send(m.Sender, "Usuário já cadastrado")
		}
		bot.Send(m.Sender, "Para informações de utilização envie:\n\n/help")
	})

	bot.Handle("/status", func(m *tgbot.Message) {
		if database.HasUsuarioRegistred(m.Sender.ID) {
			if m.Payload != "" {
				payload := m.Payload
				var tickers = strings.Split(payload, " ")
				bot.Send(m.Sender, "Buscando os dados das cotações....")
				for _, ticker := range tickers {
					var msg = GetQuoteDataAsMessage(strings.ToUpper(ticker))
					bot.Send(m.Sender, msg)
				}
			} else {
				var msg = "Utilização: /status ticker ticker1 ... tickerN"
				bot.Send(m.Sender, msg)
			}
		} else {
			bot.Send(m.Sender, "Usuário não registrado!!!\nEnvie o cammnado abaixo para registrar:\n\n/start")
		}
	})

	bot.Handle("/check", func(m *tgbot.Message) {
		if database.HasUsuarioRegistred(m.Sender.ID) {
			if m.Payload != "" {
				payload := strings.Split(m.Payload, " ")

				if payload[0] == "variacao" {
					percent, err := strconv.ParseFloat(strings.ReplaceAll(payload[2], ",", "."), 8)
					if err != nil {
						bot.Send(m.Sender, "Valor inválido: "+payload[2])
					} else {
						var acao = models.Acoes{UserID: m.Sender.ID, Ticker: strings.ToUpper(payload[1]), Variacao: percent}
						database.SaveAcao(&acao)
						bot.Send(m.Sender, "Ação salva com sucesso: \n\nTicker: "+strings.ToUpper(payload[1])+"\nVariação: "+payload[2]+"%")
					}
				} else if payload[0] == "valor" {
					var ok = true
					valorMinimo, err := strconv.ParseFloat(strings.ReplaceAll(payload[2], ",", "."), 8)
					if err != nil {
						ok = false
						bot.Send(m.Sender, "Valor inválido: "+payload[2])
					}
					valorMaximo, err := strconv.ParseFloat(strings.ReplaceAll(payload[3], ",", "."), 8)
					if err != nil {
						ok = false
						bot.Send(m.Sender, "Valor inválido: "+payload[3])
					}

					if ok {
						var acao = models.Acoes{UserID: m.Sender.ID, Ticker: strings.ToUpper(payload[1]), ValorMinimo: valorMinimo, ValorMaximo: valorMaximo}
						database.SaveAcao(&acao)
						bot.Send(m.Sender, "Ação salva com sucesso: \n\nTicker: "+strings.ToUpper(payload[1])+"\nValor maxímo (R$): "+payload[3]+"\nValor mínimo (R$): "+payload[2])
					}
				}
			}
		} else {
			bot.Send(m.Sender, "Usuário não registrado!!!\nEnvie o cammnado abaixo para registrar:\n\n/start")
		}
	})

	bot.Start()
}

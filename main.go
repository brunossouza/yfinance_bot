package main

import (
	"flag"
	"fmt"
	"os"
	"yfinance/controllers"
	"yfinance/database"
	"yfinance/models"
)

// variaveis de configuração
var (
	configCmd   *flag.FlagSet
	tgBotToken  string
	configReset bool
)

// variaveis de utilização
var (
	actionCmd    *flag.FlagSet
	ticker       string
	listTickers  bool
	statusTicker bool
)

func init() {
	// Create the batabase
	database.MigrateDatabaseSchema()

	configCmd = flag.NewFlagSet("config", flag.ExitOnError)
	configCmd.StringVar(&tgBotToken, "token", "", "Token do seu Telegram Bot")

	actionCmd = flag.NewFlagSet("acao", flag.ExitOnError)
	actionCmd.StringVar(&ticker, "t", "", "Ticker da ação")
	actionCmd.BoolVar(&listTickers, "l", false, "Lista acções cadastradas")
}

func main() {

	// var tickers = [4]string{"WEGE3", "ITSA4", "ROMI3", "JHSF3"}

	// fmt.Println("")

	// for _, ticker := range tickers {

	// 	var quote = controllers.GetQuoteData(ticker)
	// 	fmt.Println("Ticker:\t\t", ticker)
	// 	fmt.Println("Empresa:\t", quote.QuoteResponse.Result[0].LongName)
	// 	fmt.Println("Fechamento:\t", quote.QuoteResponse.Result[0].RegularMarketPreviousClose)
	// 	fmt.Println("Abertura:\t", quote.QuoteResponse.Result[0].RegularMarketOpen)
	// 	fmt.Println("Preço Atual:\t", quote.QuoteResponse.Result[0].RegularMarketPrice)
	// 	fmt.Println("Variação (R$):\t", quote.QuoteResponse.Result[0].RegularMarketChange)
	// 	fmt.Println("Variação (%):\t", quote.QuoteResponse.Result[0].RegularMarketChangePercent)

	// 	fmt.Println("")
	// }

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "config":
		configCmd.Parse(os.Args[2:])
		if tgBotToken != "" {
			fmt.Println("Salvando token api do telegram.")
			database.SaveBotConfig(&models.TelegramBotConfig{Token: tgBotToken})
			fmt.Println("Token registrado com sucesso.")
		}

	case "start":
		controllers.TelegramBotHandler()

	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}

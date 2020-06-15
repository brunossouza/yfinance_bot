package main

import (
	"flag"
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
	configCmd = flag.NewFlagSet("config", flag.ExitOnError)
	configCmd.StringVar(&tgBotToken, "token", "", "Token do seu Telegram Bot")

	actionCmd = flag.NewFlagSet("acao", flag.ExitOnError)
	actionCmd.StringVar(&ticker, "t", "", "Ticker da ação")
	actionCmd.BoolVar(&listTickers, "l", false, "Lista acções cadastradas")
}

func main() {

	// var tickers = [4]string{"WEGE3", "ITSA4", "ROMI3", "JHSF3"}

	// controllers.TelegramBotHandler("token")

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

	// if len(os.Args) < 2 {
	// 	fmt.Println("expected 'foo' or 'bar' subcommands")
	// 	os.Exit(1)
	// }

	// switch os.Args[1] {
	// case "config":
	// 	configCmd.Parse(os.Args[2:])
	// 	if tgBotToken != "" {
	// 		database.SaveBotConfig(&models.TelegramBotConfig{Token: tgBotToken})
	// 	}

	// case "acao":
	// 	actionCmd.Parse(os.Args[2:])
	// 	fmt.Println("subcommand 'acao'")
	// 	fmt.Println("  tgBotToken:", ticker)

	// default:
	// 	fmt.Println("expected 'foo' or 'bar' subcommands")
	// 	os.Exit(1)
	// }
}

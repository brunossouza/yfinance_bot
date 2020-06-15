package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"yfinance/models"
	"yfinance/utils"
)

// GetQuoteData retorna as cotações
func GetQuoteData(ticker string) (quote models.YahooFinanceQuote) {

	const URLYahooFinanceQuote = "https://query1.finance.yahoo.com/v7/finance/quote?symbols="

	response, err := http.Get(URLYahooFinanceQuote + ticker + ".SA")
	utils.CheckError(err)

	body, err := ioutil.ReadAll(response.Body)
	utils.CheckError(err)

	quote, err = models.UnmarshalYahooFinanceQuote(body)
	utils.CheckError(err)

	return
}

// GetQuoteDataAsMessage retorna as cotações
func GetQuoteDataAsMessage(ticker string) (message string) {

	var quote = GetQuoteData(ticker)

	if len(quote.QuoteResponse.Result) != 0 {
		message = fmt.Sprintf("Ticker: %s \n\nEmpresa: %s\nFechamento: %.2f\nAbertura: %.2f\nPreço Atual: %.2f\nVariação (R$): %.2f\nVariação (%%): %.2f",
			ticker, quote.QuoteResponse.Result[0].LongName, quote.QuoteResponse.Result[0].RegularMarketPreviousClose, quote.QuoteResponse.Result[0].RegularMarketOpen,
			quote.QuoteResponse.Result[0].RegularMarketPrice, quote.QuoteResponse.Result[0].RegularMarketChange, quote.QuoteResponse.Result[0].RegularMarketChangePercent)
	} else {
		message = fmt.Sprintf("Ticker: %s\n\nNão encontrado.", ticker)
	}

	return
}

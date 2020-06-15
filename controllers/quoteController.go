package controllers

import (
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

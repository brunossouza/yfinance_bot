package controllers

import (
	"time"
	"yfinance/database"
)

// VerificarQuotacoes all
func VerificarQuotacoes() {
	for {
		for _, q := range database.GetAllAcoes() {
			var quote = GetQuoteData(q.Ticker)
			if q.Variacao != 0 && q.Variacao <= quote.QuoteResponse.Result[0].RegularMarketChangePercent {
				var msg = GetQuoteDataAsMessage(q.Ticker)
				SendMessage(q.UserID, msg)
			}

			if q.Variacao == 0 && (q.ValorMinimo <= quote.QuoteResponse.Result[0].RegularMarketDayLow || q.ValorMaximo >= quote.QuoteResponse.Result[0].RegularMarketDayHigh) {
				var msg = GetQuoteDataAsMessage(q.Ticker)
				SendMessage(q.UserID, msg)
			}
		}
		time.Sleep(time.Duration(5 * time.Second))
	}
}
package utill

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	openExchangeRatesAPIURL = "https://open.er-api.com/v6/latest/USD"
)

type OpenExchangeRatesResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func FetchExchangeRates() (float64, error) {
	// Send HTTP GET request to the Open Exchange Rates API
	resp, err := http.Get(openExchangeRatesAPIURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var data OpenExchangeRatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	// Extract the exchange rate for IDR from the response
	exchangeRate, ok := data.Rates["IDR"]
	if !ok {
		return 0, fmt.Errorf("exchange rate for IDR not found in response")
	}

	return exchangeRate, nil
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetRate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	currentRate, err := GetCurrentUAHRate()
	if err != nil {
		Log(r).WithError(err).Error("failed to get current rate")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(currentRate)
	if err != nil {
		Log(r).WithError(err).Error("failed to encode response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetCurrentUAHRate() (float64, error) {
	response, err := http.Get(exchangeAPI)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get current UAH rate: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0.0, fmt.Errorf("failed to read response body: %w", err)
	}

	var currentUAHRateResponse CurrentUAHRateResponse
	err = json.Unmarshal(body, &currentUAHRateResponse)
	if err != nil {
		return 0.0, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	var currentRate float64
	for _, rate := range currentUAHRateResponse {
		if rate.Currency == USD {
			currentRate, err = strconv.ParseFloat(rate.Rate, 64)
			if err != nil {
				return 0.0, fmt.Errorf("failed to parse rate: %w", err)
			}
		}
	}

	return currentRate, nil
}

const (
	USD = "USD"

	// hardcoded because another API needs a different JSON schema to parse
	exchangeAPI = "https://api.privatbank.ua/p24api/pubinfo"
)

type (
	CurrentUAHRate struct {
		Currency string `json:"ccy"`
		Rate     string `json:"buy"`
	}

	CurrentUAHRateResponse []CurrentUAHRate
)

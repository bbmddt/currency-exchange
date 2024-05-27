package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ExchangeRates struct {
	Currencies map[string]map[string]float64 `json:"currencies"`
}

type CurrencyExchangeService struct {
	rates *ExchangeRates
}

// NewCurrencyExchangeService return a new CurrencyExchangeService with provided rates
func NewCurrencyExchangeService(filePath string) (*CurrencyExchangeService, error) {
	rates, err := loadRatesFromFile(filePath)
	if err != nil {
		return nil, err
	}

	return &CurrencyExchangeService{rates: rates}, nil
}

// loadRatesFromFile load exchange rates from a JSON file
func loadRatesFromFile(filePath string) (*ExchangeRates, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var rates ExchangeRates
	if err := json.Unmarshal(data, &rates); err != nil {
		return nil, err
	}

	return &rates, nil
}

// convert an amount from one currency to another
func (s *CurrencyExchangeService) Convert(source, target string, amount string) (string, error) {
	source = strings.ToUpper(source)
	target = strings.ToUpper(target)

	if _, ok := s.rates.Currencies[source]; !ok {
		return "", errors.New("unsupported source currency")
	}
	if _, ok := s.rates.Currencies[source][target]; !ok {
		return "", errors.New("unsupported target currency")
	}

	cleanedAmount := strings.ReplaceAll(amount, ",", "")
	parsedAmount, err := strconv.ParseFloat(cleanedAmount, 64)
	if err != nil {
		return "", errors.New("invalid amount")
	}

	convertedAmount := parsedAmount * s.rates.Currencies[source][target]
	fmt.Printf("convertedAmount:%v \n", convertedAmount)
	roundedAmount := fmt.Sprintf("%.2f", convertedAmount)

	finalAmount := addCommas(roundedAmount)
	return finalAmount, nil
}

// add commas as thousands separator in number string
func addCommas(numStr string) string {
	parts := strings.Split(numStr, ".")
	integerPart := parts[0]

	var result strings.Builder
	count := 0

	for i := len(integerPart) - 1; i >= 0; i-- {
		if count == 3 {
			result.WriteString(",")
			count = 0
		}
		result.WriteByte(integerPart[i])
		count++
	}

	// Reverse the string
	finalIntegerPart := reverse(result.String())

	if len(parts) > 1 {
		return finalIntegerPart + "." + parts[1]
	}
	return finalIntegerPart
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

package exchangerate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ExchangeRate структура для зберігання обмінного курсу
type ExchangeRate struct {
	R030         int     `json:"r030"`
	Txt          string  `json:"txt"`
	Rate         float64 `json:"rate"`
	Cc           string  `json:"cc"`
	ExchangeDate string  `json:"exchangedate"`
}

// GetRates виконує запит до API та повертає слайс обмінних курсів
func GetRates(valcode string, date ...string) ([]ExchangeRate, error) {
	const urlNbuRate = "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?"

	url := urlNbuRate
	if date[0] != "" {
		url += "&" + "date=" + date[0]
	}
	if valcode != "" {
		url += "&" + "valcode=" + valcode
	}
	url += "&json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rates []ExchangeRate
	err = json.Unmarshal(body, &rates)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func GetRate(valcode string) (float64, error) {
	rates, err := GetRates(valcode, "")
	if err != nil {
		return 0, err
	}
	if len(rates) == 0 {
		return 0, fmt.Errorf("not found rate for %s", valcode)
	}
	return rates[0].Rate, err
}

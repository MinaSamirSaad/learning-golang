package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"test.com/go/cryptomasters/datatypes"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	currency = strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiURL, currency))
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		// convert the body to string
		// bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		response := CEXResponse{}
		jsonErr := json.Unmarshal(bodyBytes, &response)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return response.GetRate(), nil
	} else {
		return nil, fmt.Errorf("API request failed with status code %d", res.StatusCode)
	}
}

func JsonParse(data []byte) (*CEXResponse, error) {
	response := CEXResponse{}
	jsonErr := json.Unmarshal(data, &response)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &response, jsonErr
}

func GetCurrencyData(currency string) {
	rate, err := GetRate(currency)

	if err == nil {
		fmt.Printf("the currency is %s and price is %f \n", rate.Currency, rate.Price)
		return
	}
}

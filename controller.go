package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetCurrencyPrice(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	symbol := p.ByName("symbol")
	if symbol != "all" {

		var currencyInput Currency
		// var currencyOutput CurrencyResponse
		url := "https://api.hitbtc.com/api/3/public/currency/" + symbol
		requestData := ""
		headers := map[string]string{}

		content, _, err := RequestAPIData("GET", url, requestData, headers)
		if err != nil {
			panic("Something went wrong")
		} else {
			err := json.Unmarshal(content, &currencyInput)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		fmt.Println(currencyInput)
	} else {

	}

	// SuccessResponse(w, r, http.StatusOK, sports, "Success")
	return
}

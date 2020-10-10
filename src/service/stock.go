package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func LoadStockPrice(ticker string) (float64, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.google.com/search?q=%s", ticker))
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	text := string(bodyBytes)
	priceReg, err := regexp.Compile("<div class=\"BNeawe iBp4i AP7Wnd\">([\\d.]+) <span")
	if err != nil {
		return 0, err
	}

	price := ""
	priceMatch := priceReg.FindStringSubmatch(text)
	if len(priceMatch) > 1 {
		price = priceMatch[1]
	} else {
		return 0, errors.New("invalid stock")
	}

	log.Println(fmt.Sprintf("Ticker: %s, Price: %s", ticker, price))
	return strconv.ParseFloat(price, 64)
}

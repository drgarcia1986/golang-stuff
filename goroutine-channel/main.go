package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type CurrencyResult struct {
	currency string
	value    string
}

var currencyRegex = regexp.MustCompile("<input type=\"text\" id=\"nacional\" value=\"([^\"]+)\"/>")
var currencyUrls = map[string]string{
	"euro":   "http://eurohoje.com",
	"dollar": "http://dolarhoje.com",
	"libra":  "http://librahoje.com/",
	"peso":   "http://pesohoje.com/",
}

func getPageBody(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func getCurrency(currency, url string, ch chan<- CurrencyResult) {
	body, _ := getPageBody(url)
	match := currencyRegex.FindSubmatch(body)
	ch <- CurrencyResult{currency: currency, value: string(match[1])}
}

func main() {
	ch := make(chan CurrencyResult, len(currencyUrls))
	for k, v := range currencyUrls {
		go getCurrency(k, v, ch)
	}
	defer close(ch)

	for range currencyUrls {
		r := <-ch
		fmt.Printf("%s - R$ %s\n", r.currency, r.value)
	}
}

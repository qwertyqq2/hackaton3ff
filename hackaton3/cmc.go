package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"encoding/json"
)

var (
	baseURL = "https://api1.binance.com"
)

type ParamsReq struct {
	symbols string
}

type PriceSymbol struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func ReqPrice(p *ParamsReq) *http.Request {
	resource := "/api/v3/ticker/price"
	data := url.Values{}
	data.Set("symbols", p.symbols)

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource

	u.RawQuery = data.Encode()

	urlStr := fmt.Sprintf("%v", u)

	fmt.Println(urlStr)
	r, _ := http.NewRequest("GET", urlStr, nil)

	return r
}

func main() {

	client := &http.Client{}

	p := &ParamsReq{
		symbols: `["BTCUSDT","BNBUSDT"]`,
	}

	req := ReqPrice(p)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var prices []*PriceSymbol
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &prices); err != nil {
		log.Fatal(err)
	}

	fmt.Println(prices)

}

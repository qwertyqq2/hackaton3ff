package apiexchange

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"encoding/json"
)

const (
	baseURL = "https://api1.binance.com"
)

type ParamsReq struct {
	symbol string
}

func NewParamsReq(symbol string) *ParamsReq {
	return &ParamsReq{symbol: symbol}
}

type PriceSymbol struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func ReqPrice(p *ParamsReq) *http.Request {
	resource := "/api/v3/ticker/price"
	data := url.Values{}
	data.Set("symbol", p.symbol)

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = resource

	u.RawQuery = data.Encode()

	urlStr := fmt.Sprintf("%v", u)

	fmt.Println(urlStr)
	r, _ := http.NewRequest("GET", urlStr, nil)
	return r
}

func SendReq(p *ParamsReq) (*PriceSymbol, error) {

	client := &http.Client{}

	req := ReqPrice(p)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var price PriceSymbol
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &price); err != nil {
		return nil, err
	}

	return &price, nil

}

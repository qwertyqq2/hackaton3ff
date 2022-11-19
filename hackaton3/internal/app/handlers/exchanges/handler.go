package exchanges

import (
	"encoding/json"
	"fmt"
	"log"
	"test_go/internal/app/handlers"
	"test_go/internal/app/handlers/exchanges/apiexchange"
	"test_go/internal/app/store"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	tickerUrl = "/ticker"
)

type handler struct {
	store store.Store
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(tickerUrl, h.GetTicker)
}

func New(s store.Store) handlers.Handler {
	return &handler{s}
}

func (h *handler) GetTicker(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	symbol := r.URL.Query().Get("symbol")
	paramsReq := apiexchange.NewParamsReq(symbol)
	price, err := apiexchange.SendReq(paramsReq)
	fmt.Println(price.Price)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json_price, err := json.Marshal(price)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json_price)
}

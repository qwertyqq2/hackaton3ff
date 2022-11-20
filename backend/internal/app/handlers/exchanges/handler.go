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
	fmt.Println("Here")
	symbol := r.URL.Query().Get("symbol")
	paramsReq := apiexchange.NewParamsReq(symbol)
	price, err := apiexchange.SendReq(paramsReq)
	if err != nil {
		log.Fatal(err)
	}
	json_price, err := json.Marshal(price)
	if err != nil {
		log.Fatal(err)
	}
	setupCORS(&w)
	w.Header().Set("Content-Type", "application/json")
	if (*r).Method == "OPTIONS" {
		log.Println("OPT")
		return
	}
	w.Write(json_price)
}

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

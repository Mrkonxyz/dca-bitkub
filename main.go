package main

import (
	"Mrkonxyz/github.com/api"
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/handler"
	"net/http"
)

func main() {
	cfg := config.LoadConfig(".")
	apiService := api.NewApiService(&cfg)
	bk := bitkub.NewBitkubService(apiService)
	h := handler.NewHandler(bk)
	http.HandleFunc("/wallet", h.GetWallet)
	http.HandleFunc("/dca-btc", h.BuyBitCion)
	http.ListenAndServe(":8080", nil)
}

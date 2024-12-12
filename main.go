package main

import (
	"Mrkonxyz/github.com/api"
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/config"
	"fmt"
)

func main() {
	cfg := config.LoadConfig(".")

	apiService := api.NewApiService(&cfg)

	bk := bitkub.NewBitkubService(apiService)

	res := bk.GetWallet()
	wallet := make(map[string]float64)
	for k, v := range res.Result {
		if val, ok := v.(float64); ok {
			wallet[k] = val

		}

	}

	fmt.Printf("wallet: %v\n", wallet)

	buyRs, _ := bk.BuyBitCion()

	fmt.Printf("buyRs: %v\n", buyRs)

}

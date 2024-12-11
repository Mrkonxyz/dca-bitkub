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
		if v > 0 {
			wallet[k] = v
		}
	}

	fmt.Printf("wallet: %v\n", wallet)

}

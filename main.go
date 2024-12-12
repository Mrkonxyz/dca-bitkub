package main

import (
	"Mrkonxyz/github.com/config"
	"fmt"
	"time"
)

func main() {
	cfg := config.LoadConfig(".")

	// apiService := api.NewApiService(&cfg)
	for {
		fmt.Printf("cfg.BaseUrl: %v\n", cfg.BaseUrl)
		time.Sleep(2 * time.Second)
	}
	// bk := bitkub.NewBitkubService(apiService)

	// res := bk.GetWallet()
	// wallet := make(map[string]float64)
	// for k, v := range res.Result {
	// 	if v > 0 {
	// 		wallet[k] = v
	// 	}
	// }

	// fmt.Printf("wallet: %v\n", wallet)

}

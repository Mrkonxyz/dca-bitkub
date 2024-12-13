package main

import (
	"Mrkonxyz/github.com/api"
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig(".")
	if cfg.ApiKey == "" || cfg.ApiSecret == "" || cfg.BaseUrl == "" {
		log.Panic("error can't read env")
	}
	apiService := api.NewApiService(&cfg)
	bk := bitkub.NewBitkubService(apiService)
	h := handler.NewHandler(bk)
	r := gin.Default()
	r.GET("/", h.Health)
	r.POST("/dca-bitcoin", h.BuyBitCion)
	r.Run()
}

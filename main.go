package main

import (
	"Mrkonxyz/github.com/api"
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig(".")
	apiService := api.NewApiService(&cfg)
	bk := bitkub.NewBitkubService(apiService)
	h := handler.NewHandler(bk)
	r := gin.Default()
	r.GET("/", h.BuyBitCion)
	r.Run()
}

package main

import (
	"Mrkonxyz/github.com/api"
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/discord"
	"Mrkonxyz/github.com/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig(".")
	cfg.Validate()
	apiService := api.NewApiService(&cfg)
	bk := bitkub.NewBitkubService(apiService)
	ds := discord.NewDiscordService(apiService)
	h := handler.NewHandler(bk, ds)

	r := gin.Default()
	r.GET("/", h.Health)
	r.GET("/wallet", h.GetBTC)
	r.POST("/dca-bitcoin", h.BuyBitCion)
	r.Run()
}

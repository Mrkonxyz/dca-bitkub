package main

import (
	"Mrkonxyz/github.com/api"
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/discord"
	"Mrkonxyz/github.com/handler"
	"Mrkonxyz/github.com/middlewere"

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
	r.Use()
	r.GET("/", h.Health)
	protected := r.Group("")
	protected.Use(middlewere.AuthMiddleware(cfg))
	protected.GET("/wallet", h.GetWallet)
	protected.POST("/dca-bitcoin", h.DcaBTC)
	r.Run()
}

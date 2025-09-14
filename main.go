package main

import (
	"Mrkonxyz/github.com/api"
	"Mrkonxyz/github.com/bitkub"
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/discord"
	"Mrkonxyz/github.com/handler"
	"Mrkonxyz/github.com/middlewere"
	"Mrkonxyz/github.com/service"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig(".")
	cfg.Validate()

	// Set a 10-second timeout for connecting
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a client
	client, err := config.ConnectMongoDB(cfg)
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer client.Disconnect(ctx)

	apiService := api.NewApiService(&cfg)
	bk := bitkub.NewBitkubService(apiService, client)
	ds := discord.NewDiscordService(apiService)

	service := service.NewService(client, &cfg, ctx)

	h := handler.NewHandler(bk, ds, service)

	r := gin.Default()
	r.Use()
	r.GET("/health", h.Health)
	// Public routes
	user := r.Group("user")
	{
		user.POST("/", h.CreateUser)
		user.GET("/:username", h.GetUserByUsername)
	}

	// private routes
	private := r.Group("dca")
	{
		private.Use(middlewere.AuthMiddleware(cfg))
		private.GET("/wallet", h.GetWallet)
		private.POST("/dca-bitcoin", h.DcaBTC)
	}

	r.Run()
}

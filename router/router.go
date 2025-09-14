package router

import (
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg config.Config, dcaHandler *handler.DcaHandler) *gin.Engine {

	r := gin.Default()
	r.Use()

	DcaRoutes(r.Group(""), dcaHandler, cfg)
	return r
}

// apiService := api.NewApiService(&cfg)
// bk := bitkub.NewBitkubService(apiService, client)
// ds := discord.NewDiscordService(apiService)

// service := service.NewService(client, &cfg, context.Background())

// h := handler.NewHandler(bk, ds, service)
// r.GET("/health", h.Health)
// // Public routes
// user := r.Group("dca")
// {
// 	user.POST("/", h.CreateUser)
// 	user.GET("/:username", h.GetUserByUsername)
// }

// // private routes

// private := r.Group("")
// private.Use(middlewere.AuthMiddleware(cfg))
// dca := private.Group("dca")
// {
// 	dca.POST("/", h.DcaBTC)
// }
// settings := private.Group("settings")
// {
// 	settings.GET("/", h.GetSettings)
// 	settings.POST("/", h.UpdateSettings)
// }
// wallet := private.Group("wallet")
// {
// 	wallet.GET("/", h.GetWallet)
// }

// 	return
// }

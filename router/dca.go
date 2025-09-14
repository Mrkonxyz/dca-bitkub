package router

import (
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/handler"

	"github.com/gin-gonic/gin"
)

func DcaRoutes(r *gin.RouterGroup, h *handler.DcaHandler, cfg config.Config) {
	dca := r.Group("/dca")
	{
		dca.POST("/", h.CreateDca)
		dca.GET("/", h.GetDca)
		dca.GET("/wallet", h.GetWallet)
		dca.POST("/trigger", h.Trigger)
	}
}

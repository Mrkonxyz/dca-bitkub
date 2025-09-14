package router

import (
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/handler"
	"Mrkonxyz/github.com/middlewere"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg config.Config, dcaHandler *handler.DcaHandler) *gin.Engine {
	r := gin.Default()
	// public routes
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	// private routes
	r.Use(middlewere.AuthMiddleware(cfg))
	DcaRoutes(r.Group(""), dcaHandler, cfg)
	return r
}

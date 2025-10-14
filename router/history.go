package router

import (
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/handler"

	"github.com/gin-gonic/gin"
)

func HistoryRoutes(r *gin.RouterGroup, h *handler.HistoryHandler, cfg config.Config) {
	history := r.Group("/history")
	deposit := history.Group("/deposit")
	{
		deposit.GET("sync", h.SyncDepositHistory)
	}
}

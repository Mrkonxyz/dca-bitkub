package handler

import (
	"Mrkonxyz/github.com/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HistoryHandler struct {
	service *service.HistoryService
}

func NewHistoryHandler(service *service.HistoryService) *HistoryHandler {
	return &HistoryHandler{service: service}
}

func (h *HistoryHandler) SyncDepositHistory(c *gin.Context) {
	ctx := c.Request.Context()
	res, err := h.service.SyncDepositHistory(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

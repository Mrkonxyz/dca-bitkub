package handler

import (
	"Mrkonxyz/github.com/bitkub"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service bitkub.Bitkub
}

func NewHandler(service *bitkub.Bitkub) *Handler {
	return &Handler{Service: *service}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "service up."})
}
func (h *Handler) GetWallet(c *gin.Context) {
	res := h.Service.GetWallet()
	wallet := make(map[string]float64)
	for k, v := range res.Result {
		if val, ok := v.(float64); ok {
			if val > 0 {
				wallet[k] = val
			}
		}

	}

	c.JSON(http.StatusOK, wallet)
}

func (h *Handler) BuyBitCion(c *gin.Context) {
	res, _ := h.Service.BuyBitCion(200)
	c.JSON(http.StatusOK, res)
}

package handler

import (
	"Mrkonxyz/github.com/bitkub"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service bitkub.Bitkub
}

func NewHandler(service *bitkub.Bitkub) *Handler {
	return &Handler{Service: *service}
}
func (h *Handler) GetWallet(w http.ResponseWriter, r *http.Request) {
	res := h.Service.GetWallet()
	wallet := make(map[string]float64)
	for k, v := range res.Result {
		if val, ok := v.(float64); ok {
			if val > 0 {
				wallet[k] = val
			}
		}

	}
	json.NewEncoder(w).Encode(wallet)
}

func (h *Handler) BuyBitCion(w http.ResponseWriter, r *http.Request) {
	res, _ := h.Service.BuyBitCion(200)
	json.NewEncoder(w).Encode(res)
}

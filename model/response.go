package model

type Response struct {
	Error  uint               `json:"error"`
	Result map[string]float64 `json:"result"`
}

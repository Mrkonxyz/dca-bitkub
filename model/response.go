package model

type Response struct {
	Error  uint                   `json:"error"`
	Result map[string]interface{} `json:"result"`
}

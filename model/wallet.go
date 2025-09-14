package model

type GeWalletResponse struct {
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"Amount"`
	AmountTHB string  `json:"AmountTHB"`
}

type WalletResponse struct {
	Error  int                `json:"error"`
	Result map[string]float64 `json:"result"`
}

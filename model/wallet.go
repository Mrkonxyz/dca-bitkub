package model

type Wallet struct {
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"Amount"`
	AmountTHB string  `json:"AmountTHB"`
}

type WalletInfoResponse struct {
	Wallet    []Wallet `json:"wallet"`
	Principle float64  `json:"principle"`
	Profit    float64  `json:"profit"`
}

type WalletResponse struct {
	Error  int                `json:"error"`
	Result map[string]float64 `json:"result"`
}

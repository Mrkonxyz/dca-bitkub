package model

type Wallet struct {
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"Amount"`
	AmountTHB string  `json:"AmountTHB"`
}

type WalletInfoResponse struct {
	Wallet    []Wallet `json:"wallet"`
	Principle string   `json:"principle"`
	Profit    string   `json:"profit"`
}

type WalletResponse struct {
	Error  int                `json:"error"`
	Result map[string]float64 `json:"result"`
}

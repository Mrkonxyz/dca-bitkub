package model

type DcaOrder struct {
	ID         string  `json:"id" bson:"_id,unique"`
	UserID     string  `json:"user_id" bson:"user_id"`
	Amount     float64 `json:"amount" bson:"amount"`
	Symbol     string  `json:"symbol" bson:"symbol"`
	SymbolInfo string  `json:"symbol_info" bson:"symbol_info"`
}

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dca struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID string             `json:"user_id" bson:"user_id"`
	Amount float64            `json:"amount" bson:"amount"`
	Symbol string             `json:"symbol" bson:"symbol"`
}

type BuyCryptoRequest struct {
	Symbol    string  `json:"sym"`
	Amount    float64 `json:"amt"`
	Rate      float64 `json:"rat"`
	OrderType string  `json:"typ"`
	ClientId  string  `json:"client_id"`
	PostOnly  bool    `json:"post_only"`
}

type BuyCryptoResponse struct {
	Error  int    `json:"error"`
	Result Result `json:"result"`
}

type Result struct {
	ID       string  `json:"id"`   // order id
	Hash     string  `json:"hash"` // order hash
	Typ      string  `json:"typ"`  // order type
	Amt      float64 `json:"amt"`  // spending amount
	Rat      float64 `json:"rat"`  // rate
	Fee      float64 `json:"fee"`  // fee
	Cre      float64 `json:"cre"`  // fee credit used
	Rec      float64 `json:"rec"`  // amount to receive
	Ts       string  `json:"ts"`   // timestamp
	ClientID string  `json:"ci"`   // input id for reference
}

type MarketData struct {
	Markets map[string]MarketDetails `json:"-"` // Map key corresponds to THB_1INCH, THB_AAVE, etc.
}

type MarketDetails struct {
	ID            int     `json:"id"`
	Last          float64 `json:"last"`
	LowestAsk     float64 `json:"lowestAsk"`
	HighestBid    float64 `json:"highestBid"`
	PercentChange float64 `json:"percentChange"`
	BaseVolume    float64 `json:"baseVolume"`
	QuoteVolume   float64 `json:"quoteVolume"`
	IsFrozen      int     `json:"isFrozen"`
	High24Hr      float64 `json:"high24hr"`
	Low24Hr       float64 `json:"low24hr"`
	Change        float64 `json:"change"`
	PrevClose     float64 `json:"prevClose"`
	PrevOpen      float64 `json:"prevOpen"`
}

type DcaRequest struct {
	Amount float64 `json:"amount"`
}

package bitkub

import (
	"bytes"
	"encoding/json"
)

// sym string The symbol you want to trade (e.g. btc_thb).
// amt float Amount you want to spend with no trailing zero (e.g. 1000.00 is invalid, 1000 is ok)
// rat float Rate you want for the order with no trailing zero (e.g. 1000.00 is invalid, 1000 is ok)
// typ string Order type: limit or market (for market order, please specify rat as 0)
// client_id string your id for reference ( not required )
// post_only bool Postonly flag: true or false ( not required )
type BuyBitCionRequest struct {
	Symbol    string  `json:"sym"`
	Amount    float64 `json:"amt"`
	Rate      float64 `json:"rat"`
	OrderType string  `json:"typ"`
	ClientId  string  `json:"client_id"`
	PostOnly  bool    `json:"post_only"`
}

type BuyBitCionResponse struct {
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

func (bk *Bitkub) BuyBitCion(amount float64) (response *BuyBitCionResponse, err error) {
	path := "/api/v3/market/place-bid"
	body := BuyBitCionRequest{
		Symbol:    "btc_thb",
		Amount:    amount,
		Rate:      0,
		OrderType: "market",
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	res, err := bk.ApiService.PostWithSig(path, bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res, &response); err != nil {
		return nil, err
	}
	return response, nil
}

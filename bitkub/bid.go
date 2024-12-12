package bitkub

import (
	"Mrkonxyz/github.com/model"
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

func (bk *Bitkub) BuyBitCion(amount float64) (response model.Response, err error) {
	path := "/api/v3/market/place-bid"
	body := BuyBitCionRequest{
		Symbol:    "btc_thb",
		Amount:    amount,
		Rate:      0,
		OrderType: "market",
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return
	}
	response = bk.ApiService.Post(path, bytes.NewBuffer(jsonData))
	return
}

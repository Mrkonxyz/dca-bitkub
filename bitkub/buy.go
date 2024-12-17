package bitkub

import (
	"bytes"
	"encoding/json"
)

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

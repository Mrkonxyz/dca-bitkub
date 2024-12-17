package bitkub

import (
	"encoding/json"
)

func (bk *Bitkub) GetPrice(sym string) (response map[string]MarketDetails, err error) {
	res, err := bk.ApiService.Get("https://api.bitkub.com/api/market/ticker?sym=" + sym)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	return response, nil
}

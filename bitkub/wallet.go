package bitkub

import (
	"Mrkonxyz/github.com/util"
	"encoding/json"
	"math"
)

type WalletResponse struct {
	Error  int                `json:"error"`
	Result map[string]float64 `json:"result"`
}

type GeWalletResponse struct {
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"Amount"`
	AmountTHB string  `json:"AmountTHB"`
}

func (bk *Bitkub) RoundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}

func (bk *Bitkub) GetWallet() (response []GeWalletResponse, err error) {
	path := "/api/v3/market/wallet"

	res, err := bk.ApiService.PostWithSig(path, nil)
	if err != nil {
		return nil, err
	}
	var temp WalletResponse
	if err = json.Unmarshal(res, &temp); err != nil {
		return nil, err
	}
	prices, _ := bk.GetPrice("")
	for k, v := range temp.Result {
		if v > 0 {
			price := prices["THB_"+k]
			var toThb string
			if price.Last == 0 {
				toThb = util.FormatMoney(v)
			} else {
				toThb = util.FormatMoney(v * price.Last)
			}
			response = append(response, GeWalletResponse{
				Symbol:    k,
				Amount:    v,
				AmountTHB: toThb,
			})
		}
	}

	return response, nil
}

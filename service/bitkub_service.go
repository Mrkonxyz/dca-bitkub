package service

import (
	"Mrkonxyz/github.com/model"
	"Mrkonxyz/github.com/utils"
	"bytes"
	"encoding/json"
	"math"
)

type BitKubService struct {
	ApiService *utils.ApiService
}

func NewBitKubService(apiService *utils.ApiService) *BitKubService {
	return &BitKubService{
		ApiService: apiService,
	}
}

func (bk *BitKubService) BuyBitCion(amount float64) (response *model.BuyBitCionResponse, err error) {
	path := "/api/v3/market/place-bid"
	body := model.BuyBitCionRequest{
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

func (bk *BitKubService) GetPrice(sym string) (response map[string]model.MarketDetails, err error) {
	res, err := bk.ApiService.Get("https://api.bitkub.com/api/market/ticker?sym=" + sym)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(res, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (bk *BitKubService) RoundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}

func (bk *BitKubService) GetWallet() (response []model.GeWalletResponse, err error) {
	path := "/api/v3/market/wallet"

	res, err := bk.ApiService.PostWithSig(path, nil)
	if err != nil {
		return nil, err
	}
	var temp model.WalletResponse
	if err = json.Unmarshal(res, &temp); err != nil {
		return nil, err
	}
	prices, _ := bk.GetPrice("")
	for k, v := range temp.Result {
		if v > 0 {
			price := prices["THB_"+k]
			var toThb string
			if price.Last == 0 {
				toThb = utils.FormatMoney(v)
			} else {
				toThb = utils.FormatMoney(v * price.Last)
			}
			response = append(response, model.GeWalletResponse{
				Symbol:    k,
				Amount:    v,
				AmountTHB: toThb,
			})
		}
	}

	return response, nil
}

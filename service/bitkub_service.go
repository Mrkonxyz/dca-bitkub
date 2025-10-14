package service

import (
	"Mrkonxyz/github.com/model"
	"Mrkonxyz/github.com/repository"
	"Mrkonxyz/github.com/utils"
	"bytes"
	"context"
	"encoding/json"
	"math"
)

type BitKubService struct {
	topUpRepository *repository.TopUpRepository
	ApiService      *utils.ApiService
}

func NewBitKubService(apiService *utils.ApiService, topUpRepository *repository.TopUpRepository) *BitKubService {
	return &BitKubService{
		ApiService:      apiService,
		topUpRepository: topUpRepository,
	}
}

func (bk *BitKubService) BuyCrypto(amount float64, symbol string) (response *model.BuyCryptoResponse, err error) {
	path := "/api/v3/market/place-bid"
	body := model.BuyCryptoRequest{
		Symbol:    symbol,
		Amount:    amount,
		Rate:      0,
		OrderType: "market",
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	res, err := bk.ApiService.PostWithSig(path, bytes.NewBuffer(jsonData), nil)

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

func (bk *BitKubService) GetWallet(ctx context.Context) (response *model.WalletInfoResponse, err error) {
	path := "/api/v3/market/wallet"

	res, err := bk.ApiService.PostWithSig(path, nil, nil)
	if err != nil {
		return nil, err
	}
	var temp model.WalletResponse
	if err = json.Unmarshal(res, &temp); err != nil {
		return nil, err
	}
	prices, _ := bk.GetPrice("")
	var responseTemp []model.Wallet
	sum := 0.0
	for k, v := range temp.Result {
		if v > 0 {
			price := prices["THB_"+k]
			var toThb string
			if price.Last == 0 {
				toThb = utils.FormatMoney(v)
				sum += v
			} else {
				toThb = utils.FormatMoney(v * price.Last)
				sum += v * price.Last
			}
			responseTemp = append(responseTemp, model.Wallet{
				Symbol:    k,
				Amount:    v,
				AmountTHB: toThb,
			})

		}
	}

	principle, _ := bk.topUpRepository.SumAmount(ctx)
	profit := sum - principle
	response = &model.WalletInfoResponse{
		Wallet:    responseTemp,
		Principle: utils.FormatMoney(principle),
		Profit:    utils.FormatMoney(profit),
	}

	return response, nil
}

func (bk *BitKubService) DepositHistory() (response *model.BaseResultPage[model.DepositHistoryResponse], err error) {
	path := "/api/v3/fiat/deposit-history"

	// Create a map to hold the request body parameters
	params := make(map[string]string)

	params["p"] = "1"

	res, err := bk.ApiService.PostWithSig(path, nil, params)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(res, &response); err != nil {
		return nil, err
	}
	return response, nil
}

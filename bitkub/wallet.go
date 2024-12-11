package bitkub

import (
	"Mrkonxyz/github.com/model"
)

func (bk *Bitkub) GetWallet() (response model.Response) {
	path := "/api/v3/market/wallet"

	return bk.ApiService.Post(path)
}

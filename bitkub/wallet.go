package bitkub

func (bk *Bitkub) GetWallet() (response []byte, err error) {
	path := "/api/v3/market/wallet"

	return bk.ApiService.PostWithSig(path, nil)
}

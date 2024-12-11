package bitkub

import "Mrkonxyz/github.com/api"

type Bitkub struct {
	ApiService *api.ApiService
}

func NewBitkubService(apiService *api.ApiService) *Bitkub {
	return &Bitkub{ApiService: apiService}
}

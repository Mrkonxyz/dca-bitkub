package bitkub

import (
	"Mrkonxyz/github.com/api"

	"go.mongodb.org/mongo-driver/mongo"
)

type Bitkub struct {
	ApiService *api.ApiService
	MongoDB    *mongo.Client
}

func NewBitkubService(apiService *api.ApiService, mongoDB *mongo.Client) *Bitkub {
	return &Bitkub{ApiService: apiService, MongoDB: mongoDB}
}

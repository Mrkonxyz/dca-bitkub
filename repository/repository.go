package repository

import (
	"Mrkonxyz/github.com/config"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Ctx      context.Context
	User     *mongo.Collection
	DcaOrder *mongo.Collection
}

func NewRepository(client *mongo.Client, cfg *config.Config, ctx context.Context) *Repository {
	return &Repository{
		Ctx:      ctx,
		User:     client.Database("dca-bitkub").Collection("users"),
		DcaOrder: client.Database("dca-bitkub").Collection("dca_orders"),
	}
}

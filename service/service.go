package service

import (
	"Mrkonxyz/github.com/config"
	"Mrkonxyz/github.com/repository"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(client *mongo.Client, cfg *config.Config, ctx context.Context) *Service {
	repo := repository.NewRepository(client, cfg, ctx)
	return &Service{
		Repo: repo,
	}
}

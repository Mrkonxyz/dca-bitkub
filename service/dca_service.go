package service

import (
	"Mrkonxyz/github.com/model"
	"Mrkonxyz/github.com/repository"
	"context"
)

type DcaService struct {
	repo *repository.DcaRepository
}

func NewDcaService(repo *repository.DcaRepository) *DcaService {
	return &DcaService{repo}
}

func (s *DcaService) CreateDca(ctx context.Context, d model.Dca) error {
	return s.repo.Create(ctx, d)
}

func (s *DcaService) GetDca(ctx context.Context) ([]model.Dca, error) {
	return s.repo.FindAll(ctx)
}

func (s *DcaService) RemoveDca(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *DcaService) UpdateDca(ctx context.Context, d model.Dca) error {
	return s.repo.Update(ctx, d)
}

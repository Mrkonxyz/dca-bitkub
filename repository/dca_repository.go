package repository

import (
	"context"

	"Mrkonxyz/github.com/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DcaRepository struct {
	col *mongo.Collection
}

func NewDcaRepository(db *mongo.Database) *DcaRepository {
	return &DcaRepository{
		col: db.Collection("dca_orders"),
	}
}

func (r *DcaRepository) Create(ctx context.Context, d model.Dca) error {
	_, err := r.col.InsertOne(ctx, d)
	return err
}

func (r *DcaRepository) FindAll(ctx context.Context) ([]model.Dca, error) {
	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dcas []model.Dca
	if err = cursor.All(ctx, &dcas); err != nil {
		return nil, err
	}
	return dcas, nil
}

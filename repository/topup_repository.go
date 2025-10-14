package repository

import (
	"Mrkonxyz/github.com/model/entity"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TopUpRepository struct {
	col *mongo.Collection
}

func NewTopUpRepository(db *mongo.Database) *TopUpRepository {
	return &TopUpRepository{
		col: db.Collection("topUp"),
	}
}

func (r *TopUpRepository) SaveAll(ctx context.Context, docs []entity.TopUpHistory) error {

	documents := make([]interface{}, len(docs))
	for i, doc := range docs {
		documents[i] = doc
	}

	opts := options.InsertMany().SetOrdered(false)

	_, err := r.col.InsertMany(ctx, documents, opts)

	if err != nil {
		return fmt.Errorf("TopUpRepository.saveAll: failed to insert documents into collection: %w", err)
	}

	return nil
}

func (r *TopUpRepository) SumAmount(ctx context.Context) (float64, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: nil},
			{Key: "totalAmount", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		}}},
	}

	cursor, err := r.col.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, fmt.Errorf("TopUpRepository.sumAmount: failed to execute aggregation: %w", err)
	}
	defer cursor.Close(ctx)

	var results []struct {
		TotalAmount float64 `bson:"totalAmount"`
	}

	if err := cursor.All(ctx, &results); err != nil {
		return 0, fmt.Errorf("TopUpRepository.sumAmount: failed to decode aggregation results: %w", err)
	}

	if len(results) == 0 {
		return 0, nil // No documents found, return 0
	}

	return results[0].TotalAmount, nil
}

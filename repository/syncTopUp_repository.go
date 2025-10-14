package repository

import (
	"Mrkonxyz/github.com/model/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SyncTopUpRepository struct {
	Col *mongo.Collection
}

func NewSyncTopUpRepository(db *mongo.Database) *SyncTopUpRepository {
	return &SyncTopUpRepository{
		Col: db.Collection("lastSyncTopUp"),
	}
}

func (r *SyncTopUpRepository) FindLastSync(ctx context.Context) (entity.SyncTopUpOffset, error) {
	var result entity.SyncTopUpOffset
	opts := options.FindOne().SetSort(bson.D{{Key: "lastSyncDate", Value: -1}})
	filter := bson.D{}
	err := r.Col.FindOne(ctx, filter, opts).Decode(&result)
	return result, err
}

func (r *SyncTopUpRepository) CreateLastSync(ctx context.Context, sync entity.SyncTopUpOffset) (entity.SyncTopUpOffset, error) {
	res, err := r.Col.InsertOne(ctx, sync)
	if err != nil {
		return entity.SyncTopUpOffset{}, err
	}
	id := res.InsertedID.(primitive.ObjectID)
	sync.ID = id
	return sync, nil
}

func (r *SyncTopUpRepository) UpdateLastSync(ctx context.Context, sync entity.SyncTopUpOffset) error {
	_, err := r.Col.UpdateOne(ctx, bson.M{"_id": sync.ID}, bson.M{"$set": sync})
	return err
}

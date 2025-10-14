package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SyncTopUpOffset struct {
	ID           primitive.ObjectID `bson:"_id,omitempty,unique"`
	LastSyncDate time.Time          `bson:"lastSyncDate"`
}

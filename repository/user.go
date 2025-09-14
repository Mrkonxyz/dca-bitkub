package repository

import (
	"Mrkonxyz/github.com/model"
	"log"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) CreateUser(user *model.User) error {
	user.ID = uuid.NewString()
	_, err := r.User.InsertOne(r.Ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	filter := bson.M{"username": username}
	log.Println("Querying user by username:", filter)
	err := r.User.FindOne(r.Ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

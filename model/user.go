package model

type User struct {
	ID       string `json:"id" bson:"_id,unique"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username,unique"`
	Password string `json:"password" bson:"password"`
}

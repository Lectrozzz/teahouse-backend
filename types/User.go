package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role string

const (
	Admin Role = "Admin"
	Client Role = "Client"
)

type User struct{
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
	Email string `bson:"email,omitempty"`
	Password string `bson:"password,omitempty"`
	Role Role `bson:"role,omitempty"`
	Point int `bson:"point"`
}
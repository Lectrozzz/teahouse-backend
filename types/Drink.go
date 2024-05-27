package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Drink struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name,omitempty"`
	TeaCategory string             `bson:"teacategory"`
	Series      string             `bson:"series"`
	Price       int                `bson:"price"`
}
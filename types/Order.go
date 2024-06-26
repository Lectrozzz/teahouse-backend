package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	Pending    Status = "Pending"
	Processing Status = "Processing"
	Ready  Status = "Ready"
	Rejected   Status = "Rejected"
)

type Item struct {
	DrinkID    primitive.ObjectID `bson:"drinkid"`
	SweetLevel int                `bson:"sweetlevel"`
	Amount     int                `bson:"amount"`
}

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"userid"`
	ItemList []Item             `bson:"itemlist"`
	Date     primitive.DateTime `bson:"date"`
	Status   Status             `bson:"status"`
	Price    int                `bson:"price"`
	Voucher  string             `bson:"voucher"`
}
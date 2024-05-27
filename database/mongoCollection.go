package database

import (
	"github.com/Lectrozzz/teahouse-backend/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var Database *mongo.Database
var DrinksCollection *mongo.Collection
var UsersCollection *mongo.Collection
var OrdersCollection *mongo.Collection

func LoadCollection(){
	Database = config.MongoClient.Database("lectroz")

	DrinksCollection = Database.Collection("Drinks")
	UsersCollection = Database.Collection("Users")
	OrdersCollection = Database.Collection("Orders")
}
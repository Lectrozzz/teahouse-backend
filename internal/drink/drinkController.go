package drink

import (
	"context"

	"github.com/Lectrozzz/teahouse-backend/database"
	"github.com/Lectrozzz/teahouse-backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetAllDrinksHandler(c *fiber.Ctx) error {
	cursor, err := database.DrinksCollection.Find(context.TODO(), bson.D{})
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}
	defer cursor.Close(context.TODO())

	var results []bson.M
	for cursor.Next(context.TODO()){
		var temp bson.M
		if err := cursor.Decode(&temp); err!=nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
		}
		results = append(results, temp)
	}   
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true,"data":results})
}

func GetSingleDrinkHandler(c *fiber.Ctx) error{
	var result *bson.M

	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	filter := bson.D{{Key: "_id", Value: objectID}}
	if err := database.DrinksCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "body": result})
}

func CreateDrinkHandler(c *fiber.Ctx) error {
	newDrink := new(types.Drink)
	if err:=c.BodyParser(&newDrink); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}

	result,err := database.DrinksCollection.InsertOne(context.TODO(),newDrink)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true,"message":result,"body":newDrink})
}

func UpdateDrinkHandler(c *fiber.Ctx) error {
	//TODO
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func DeleteDrinkHandler(c *fiber.Ctx) error{
	var id = c.Params("id")
	filter := bson.D{{Key: "_id", Value: id}}

	result := database.DrinksCollection.FindOneAndDelete(context.TODO(), filter)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true,"body":result})
}
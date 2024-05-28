package order

import (
	"context"
	"time"

	"github.com/Lectrozzz/teahouse-backend/database"
	"github.com/Lectrozzz/teahouse-backend/internal/utils"
	"github.com/Lectrozzz/teahouse-backend/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllOrdersHandler(c *fiber.Ctx) error {
	cursor, err := database.OrdersCollection.Find(context.TODO(), bson.D{})
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data":results})
}

func GetSingleOrderHandler(c *fiber.Ctx)error {
	id := c.Params("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	
	var result *bson.M
	filter := bson.D{{Key: "_id", Value: objectID}}
	if err := database.OrdersCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": result})
}

func GetMyOrdersHandler(c *fiber.Ctx)error {
	token := utils.GetUserToken(c)
	result, userID := utils.GetUserIDFromToken(token)
	if !result{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": "Invalid token"})
	}
	objectUserId, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.D{{Key: "userid", Value: objectUserId}}
	cursor, err := database.OrdersCollection.Find(context.TODO(), filter)
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data":results})
}

func CreateOrderHandler(c *fiber.Ctx)error {
	token := utils.GetUserToken(c)
	_, userID := utils.GetUserIDFromToken(token)
	objectUserId, err := primitive.ObjectIDFromHex(userID)

	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	newOrder := new(types.Order)
	if err:=c.BodyParser(&newOrder); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}

	newOrder.Date = primitive.NewDateTimeFromTime(time.Now())
	newOrder.UserID = objectUserId

	result, err := database.OrdersCollection.InsertOne(context.TODO(), newOrder)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": result, "body": newOrder})
}

func UpdateOrderStatusHandler(c *fiber.Ctx)error {
	id := c.Params("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	
	filter:= bson.D{{Key: "_id", Value: objectID}}
	if err:= database.OrdersCollection.FindOne(context.TODO(), filter).Err(); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	var updatedStatus *types.Status
	if err:=c.BodyParser(&updatedStatus); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}
	update:= bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: updatedStatus}}}}
	if _, err:= database.OrdersCollection.UpdateOne(context.TODO(), filter, update); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
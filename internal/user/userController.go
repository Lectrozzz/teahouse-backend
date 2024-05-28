package user

import (
	"context"

	"github.com/Lectrozzz/teahouse-backend/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
	// TODO
	return c.SendStatus(200)
}

func CollectPointsHandler(c *fiber.Ctx) error {
	type CollectPointsRequest struct {
		OrderID 	string 	`bson:"orderid"`
		PhoneNumber string 	`bson:"phonenumber"`
		Price 		int 	`bson:"price"`
	}

	var body *CollectPointsRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}

	// Find user id from phone number
	var user *bson.M
	userFilter := bson.D{{Key: "phonenumber", Value: body.PhoneNumber}}
	if err := database.UsersCollection.FindOne(context.TODO(), userFilter).Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	// Find order by order id
	var order *bson.M
	objectOrderID, _ := primitive.ObjectIDFromHex(body.OrderID)
	orderFilter := bson.D{{Key: "_id", Value: objectOrderID}}
	if err := database.OrdersCollection.FindOne(context.TODO(), orderFilter).Decode(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	// Calculate points (25 thb per 1 point)
	points := body.Price / 25

	// Update user points
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "points", Value: points}}}}
	if _, err := database.UsersCollection.UpdateOne(context.TODO(), userFilter, update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "data": points})
}
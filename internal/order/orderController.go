package order

import (
	"context"

	"github.com/Lectrozzz/teahouse-backend/database"
	"github.com/Lectrozzz/teahouse-backend/types"
	"github.com/gofiber/fiber/v2"
)

func GetAllOrdersHandler(c *fiber.Ctx) error {
	//TODO
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func GetMyOrderHandler(c *fiber.Ctx)error {
	//TODO
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func GetSingleOrderHandler(c *fiber.Ctx)error {
	//TODO
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func CreateOrderHandler(c *fiber.Ctx)error {
	//TODO
	newOrder := new(types.Order)
	if err:=c.BodyParser(&newOrder); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}

	result, err := database.OrdersCollection.InsertOne(context.TODO(), newOrder)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": result, "body": newOrder})
}

func UpdateOrderStatusHandler(c *fiber.Ctx)error {
	//TODO
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
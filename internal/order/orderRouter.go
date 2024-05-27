package order

import (
	"github.com/Lectrozzz/teahouse-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func OrderRouter(router fiber.Router){
	jwt:= middleware.IsLoggedIn()

	router.Get("", jwt, GetAllOrdersHandler)
	router.Get("me", jwt, GetMyOrderHandler)
	router.Get("/:id", jwt, GetSingleOrderHandler)
	router.Post("", jwt, CreateOrderHandler)
	router.Patch("/:id", jwt, UpdateOrderStatusHandler)
}
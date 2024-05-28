package server

import (
	"github.com/Lectrozzz/teahouse-backend/internal/auth"
	"github.com/Lectrozzz/teahouse-backend/internal/drink"
	"github.com/Lectrozzz/teahouse-backend/internal/order"
	"github.com/Lectrozzz/teahouse-backend/internal/user"
	"github.com/gofiber/fiber/v2"
)

func InitServer(){
	app := fiber.New()

	// Authentication
	app.Route("/auth", auth.AuthRouter)

	// User Profile
	app.Route("/user", user.UserRouter)

	// Drink Menu
	app.Route("/drink", drink.DrinkRouter)

	// Order
	app.Route("/order", order.OrderRouter)

	app.Listen(":8080")
}
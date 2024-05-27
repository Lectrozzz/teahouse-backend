package server

import (
	"github.com/Lectrozzz/teahouse-backend/internal/auth"
	"github.com/Lectrozzz/teahouse-backend/internal/drink"
	"github.com/Lectrozzz/teahouse-backend/internal/user"
	"github.com/gofiber/fiber/v2"
)

func InitServer(){
	app := fiber.New()

	// Authentication
	app.Route("/auth", auth.AuthRouter)

	// User Profile
	app.Route("/user", user.UserRouter)

	// Menu
	menuGroup := app.Group("/menu")
	menuGroup.Route("/drink", drink.DrinkRouter)

	// Order
	// app.Route("/order", router.UserRouter)

	app.Listen(":8080")
}
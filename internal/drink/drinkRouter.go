package drink

import (
	"github.com/Lectrozzz/teahouse-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func DrinkRouter(router fiber.Router){
	jwt:= middleware.IsLoggedIn()

	router.Get("", GetAllDrinksHandler)
	router.Get("/:id", GetSingleDrinkHandler)
	router.Post("", jwt, CreateDrinkHandler)
	router.Patch("/:id", jwt, UpdateDrinkHandler)
	router.Delete("/:id", jwt, DeleteDrinkHandler)
}
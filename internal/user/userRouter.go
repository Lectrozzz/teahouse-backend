package user

import (
	"github.com/Lectrozzz/teahouse-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router){
	jwt := middleware.IsLoggedIn()
	router.Get("", jwt, GetAllUsersHandler)
	router.Post("/collect-points", CollectPointsHandler)
}
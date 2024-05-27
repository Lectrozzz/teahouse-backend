package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(router fiber.Router){
	router.Post("/login", LoginHandler)
	router.Post("/register", RegisterHandler)
}
package user

import (
	"github.com/Lectrozzz/teahouse-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router){
	jwt := middleware.IsLoggedIn()
	router.Get("", jwt, GetAllUsersHandler)
	// router.Get("/:id", GetSingleUserHandler)
	// router.Get("/me", GetMeHandler)
	// router.Patch("/updateAccount", UpdateAccountHandler)
	// router.Delete("/deleteAccount", DeleteAccountHandler)
}
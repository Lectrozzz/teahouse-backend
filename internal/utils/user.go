package utils

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserToken(c *fiber.Ctx) string {
	token := strings.Split(c.Get("Authorization"), "Bearer ")[1]
	return token
}

func GetUserIDFromToken(tokenStr string) (bool, string) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println("Error parsing token:",err)
		return false, ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Error getting claims")
		return false, ""
	}

	return true, claims["id"].(string)
}
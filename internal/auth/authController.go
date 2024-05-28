package auth

import (
	"context"
	"os"
	"time"

	"github.com/Lectrozzz/teahouse-backend/database"
	"github.com/Lectrozzz/teahouse-backend/internal/utils"
	"github.com/Lectrozzz/teahouse-backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)



func RegisterHandler(c *fiber.Ctx) error {
	newUser := new(types.User)
	if err:=c.BodyParser(&newUser); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}

	filter := bson.D{{Key: "phone-number", Value:newUser.PhoneNumber}}
	if err := database.UsersCollection.FindOne(context.TODO(), filter).Decode(&newUser); err == nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":"Phone number has been registered"})
	}

	newPassword, err := utils.HashPassword(newUser.Password); 
	if err!= nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}

	newUser.Password = newPassword
	result,err := database.UsersCollection.InsertOne(context.TODO(), newUser)
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true,"message":result,"body":newUser})
}

func LoginHandler(c *fiber.Ctx) error {
	currentUser := new(types.User)
	userData := new(types.User)
	if err:=c.BodyParser(&currentUser); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":err.Error()})
	}

	filter := bson.D{{Key: "email", Value:currentUser.Email}}
	if err := database.UsersCollection.FindOne(context.TODO(), filter).Decode(&userData); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false,"error":"Invalid email"})
	}
	if err:=utils.ComparePassword(userData.Password, currentUser.Password); err!=nil{
		return fiber.ErrUnauthorized
	}

	claims := jwt.MapClaims{
		"id": userData.ID.Hex(),
		"email": userData.Email,
		"role": userData.Role,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:  "token",
		Value: t,
	}
	c.Cookie(&cookie)
	
	return c.Status(200).JSON(fiber.Map{"success": true,"token":t})
}
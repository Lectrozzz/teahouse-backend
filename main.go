package main

import (
	"log"

	"github.com/Lectrozzz/teahouse-backend/config"
	"github.com/Lectrozzz/teahouse-backend/database"
	"github.com/Lectrozzz/teahouse-backend/server"
	"github.com/joho/godotenv"
)

func main(){
	//Set up .env
	envError := godotenv.Load("./config/.env")
	if envError != nil {
		log.Fatal("Error loading .env file", envError)
	}
	//Set up MongoDB
	config.ConnectMongoDB()
	defer config.CloseDB()

	//Load collection from MongoDB
	database.LoadCollection()
	
	//Set up server
	server.InitServer()
}
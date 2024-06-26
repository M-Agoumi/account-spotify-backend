package main

import (
	"context"
	"fmt"
	"github.com/M-Agoumi/account-spotify-backend/application"
	"github.com/M-Agoumi/account-spotify-backend/config"
	"github.com/M-Agoumi/account-spotify-backend/model"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	errDot := godotenv.Load()
	if errDot != nil {
		log.Println("Error loading .env file")
	}

	app := application.New()

	config.ConnectDatabase() // Initialize the database connection
	errMigration := config.DB.AutoMigrate(&model.User{})

	if errMigration != nil {
		return
	} // Automatically migrate the schema

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

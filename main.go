package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"magoumi/spotify-account/application"
	"magoumi/spotify-account/config"
	"magoumi/spotify-account/model"
)

func main() {
	app := application.New()

	errDot := godotenv.Load()
	if errDot != nil {
		log.Println("Error loading .env file")
	}

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

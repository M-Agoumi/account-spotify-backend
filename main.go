package main

import (
	"fmt"
	"magoumi/spotify-account/application"
)

func main() {
	app := application.New()

	err := app.Start()
	if err != nil {
		fmt.Println("Error starting server: %w", err)
	}
}

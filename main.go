package main

import (
	"context"
	"fmt"
	"magoumi/spotify-account/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

package main

import (
	"context"
	"fmt"

	"go-jwt-rest-mongodb/database"
	"go-jwt-rest-mongodb/handler"
	"go-jwt-rest-mongodb/repository"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load environment: %s", err.Error())
	}

	client := database.ConnectMongoClient()
	Repository := repository.InitRepository(client)
	handler := handler.Handler{Repository}

	fmt.Println("Starting server on http://localhost:8080")
	handler.HandleRequest()

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

package main

import (
	"log"
	"net/http"

	"github.com/go-mongo-todos/db"
	"github.com/go-mongo-todos/handlers"
	"github.com/go-mongo-todos/services"
)

type Application struct {
	Models services.Models
}

func main() {
	// 1. Connect to the database
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Fatal("Could not connect to the database")
	}

	// 2. Initialize the service with the database client
	todoService := services.New(mongoClient)

	// 3. Initialize the handler with the service
	todoHandler := handlers.NewTodoHandler(todoService)

	// 4. Create the router and pass the handler to it
	router := handlers.CreateRouter(todoHandler)

	log.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", router)
}

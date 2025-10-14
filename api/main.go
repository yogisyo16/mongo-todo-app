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

	// 2. Initialize the services with the database client
	todoService := services.New(mongoClient)
	userService := services.User{} // Uses the same client set in services.New()

	// 3. Initialize the handlers with their respective services
	todoHandler := handlers.NewTodoHandler(todoService)
	userHandler := handlers.NewUserHandler(userService)

	// 4. Create the router and pass both handlers to it
	router := handlers.CreateRouter(todoHandler, userHandler)

	log.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", router)
}

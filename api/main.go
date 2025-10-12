package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-mongo-todos/db"
	"github.com/go-mongo-todos/handlers"
	"github.com/go-mongo-todos/services"
)

type Application struct {
	Models services.Models
}

func main() {
	mongoClient, err := db.ConnectToMongo()

	if err != nil {
		log.Panic()
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	services.New(mongoClient)

	log.Println("Services run in port", 8080)
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
}

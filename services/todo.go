package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"_task,omitempty"`
	Completed bool      `json:"completed,omitempty" bson:"_completed,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"_created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty" bson:"_update_at,omitempty"`
}

var client *mongo.Client

func New(mongo *mongo.Client) Todo {
	client = mongo
	return Todo{}
}

func returnCollectionPointer(collection string) *mongo.Collection {
	return client.Database("todos_db").Collection(collection)
}

func (t *Todo) InsertTodo(entry Todo) error {
	collection := returnCollectionPointer("todos")
	_, err := collection.InsertOne(context.TODO(), Todo{
		Task:      entry.Task,
		Completed: entry.Completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	return nil
}

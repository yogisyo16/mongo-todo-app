package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"_task,omitempty"`
	Completed bool      `json:"completed" bson:"_completed"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"_created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty" bson:"_update_at,omitempty"`
}

var client *mongo.Client

func New(mongo *mongo.Client) Todo {
	client = mongo
	return Todo{}
}

func returnTodosCollection(collection string) *mongo.Collection {
	return client.Database("todos_db").Collection(collection)
}

// GetAllTodos
func (t *Todo) GetAllTodos() ([]Todo, error) {
	collection := returnTodosCollection("todos")
	var todos []Todo
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	return todos, nil
}

// GetTodoById
func (t *Todo) GetTodoById(id string) (Todo, error) {
	collection := returnTodosCollection("todos")
	var todo Todo

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Todo{}, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": mongoID}).Decode(&todo)
	if err != nil {
		log.Println(err)
		return Todo{}, err
	}

	return todo, nil
}

// InsertTodo
func (t *Todo) InsertTodo(entry Todo) error {
	collection := returnTodosCollection("todos")
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

// UpdatedTodo

func (t *Todo) UpdatedTodo(id string, entry Todo) (*mongo.UpdateResult, error) {
	collection := returnTodosCollection("todos")
	mongoID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	update := bson.D{
		{"$set", bson.D{
			{"task", entry.Task},
			{"completed", entry.Completed},
			{"updated_at", time.Now()},
		}},
	}

	res, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": mongoID},
		update,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

// DeleteTodo

func (t *Todo) DeleteTodo(id string) error {
	collection := returnTodosCollection("todos")
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = collection.DeleteOne(
		context.Background(),
		bson.M{"_id": mongoID},
	)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

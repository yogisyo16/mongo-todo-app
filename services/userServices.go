package services

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name,omitempty" bson:"_name,omitempty"`
	Email     string    `json:"email,omitempty" bson:"_email,omitempty"`
	Password  string    `json:"password,omitempty" bson:"_password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"_created_at,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty" bson:"_update_at,omitempty"`
}

type UserService interface {
	GetAllUsers() ([]User, error)
	InsertUser(entry User) error
}

func retunrUserCollection(collection string) *mongo.Collection {
	return client.Database("todos_db").Collection(collection)
}

func (u *User) GetAllUsers() ([]User, error) {
	return nil, nil
}

func (u *User) InsertUser(entry User) error {
	return nil
}

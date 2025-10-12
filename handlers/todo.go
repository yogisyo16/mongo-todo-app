package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-mongo-todos/services"
)

var todo services.Todo

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Msg:  "Health Check",
		Code: 200,
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func CreateTodos(w http.ResponseWriter, r *http.Request) {

}

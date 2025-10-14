package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-mongo-todos/services"
)

type TodoHandler struct {
	Service services.Todo
}

func NewTodoHandler(service services.Todo) *TodoHandler {
	return &TodoHandler{
		Service: service,
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Msg:  "Health Check",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func (h *TodoHandler) getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.GetAllTodos()
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) getTodoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	todo, err := h.Service.GetTodoById(id)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo services.Todo

	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
	}

	err = h.Service.InsertTodo(newTodo)
	if err != nil {
		errorRes := Response{
			Msg:  "Error",
			Code: 304,
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	res := Response{
		Msg:  "Succesfully Created Todo",
		Code: 201,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func (h *TodoHandler) updateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updateTodo services.Todo

	err := json.NewDecoder(r.Body).Decode(&updateTodo)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = h.Service.UpdatedTodo(id, updateTodo)
	if err != nil {
		errorRes := Response{
			Msg:  err.Error(),
			Code: 500,
		}
		jsonStr, err := json.Marshal(errorRes)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(errorRes.Code)
		w.Write(jsonStr)
		return
	}

	res := Response{
		Msg:  "Succesfully Updated Todo",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func (h *TodoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.Service.DeleteTodo(id)
	if err != nil {
		errorRes := Response{
			Msg:  "Error",
			Code: 304,
		}
		json.NewEncoder(w).Encode(errorRes)
		w.WriteHeader(errorRes.Code)
		return
	}

	res := Response{
		Msg:  "Succesfully Deleted Todo",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

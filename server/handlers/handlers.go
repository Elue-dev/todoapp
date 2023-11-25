package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/elue-dev/todoapi/models"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatalf("failed to decode json body %v", err)
	}

	// result := controllers.CreateTodo()

}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {

}

func GetSingleTodo(w http.ResponseWriter, r *http.Request) {

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

}
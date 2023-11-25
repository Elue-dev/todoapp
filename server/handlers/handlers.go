package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/elue-dev/todoapi/controllers"
	"github.com/elue-dev/todoapi/models"
	"github.com/gorilla/mux"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatalf("failed to decode json body %v", err)
	}

	result := controllers.CreateTodo(todo)
	
	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data: result,
	})

}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	todos, err := controllers.GetTodos()
	if err != nil {
		log.Fatalf("Could not get all todos %v", err)
	}

	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data: todos,
	})
}

func GetSingleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := controllers.GetTodo(params["id"])

	if err != nil {
		log.Fatalf("failed to get todo %v", err)
	}
	
	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Data: result,
	})
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

}
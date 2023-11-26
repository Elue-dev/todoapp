package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elue-dev/todoapi/controllers"
	"github.com/elue-dev/todoapi/helpers"
	"github.com/elue-dev/todoapi/models"
	"github.com/gorilla/mux"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	var todo models.Todo
	
	err := json.NewDecoder(r.Body).Decode(&todo)

	isValidated := helpers.ValidateRequestBody(todo.Title, todo.Description)

	if !isValidated {
		json.NewEncoder(w).Encode(models.ErrResponse{
			Success: false,
			Error: "Please provide todo title and description",
		})
		return
	}

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
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Fatalf("failed to decode json body %v", err)
	}

	updatedRows, err := controllers.UpdateTodo(params["id"], todo)
	fmt.Println("total rows affected", updatedRows)

	if err != nil {
		log.Fatalf("could not update stock %v", err)
	}

	json.NewEncoder(w).Encode(models.MsgResponse{
		Success: true,
		Message: "Todo updated succesfully",
	})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

}
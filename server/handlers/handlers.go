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

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		isValidated := helpers.ValidateRequestBody(todo.Title, todo.Description)
		if !isValidated {
			json.NewEncoder(w).Encode(models.ErrResponse{
				Success: false,
				Error: "Please provide todo title and description",
			})
		} else {
			json.NewEncoder(w).Encode(models.ErrResponse{
				Success: false,
				Error: "Something went wrong, please try again",
			})
			log.Fatalf("failed to decode json body %v", err)
		}
		return 
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
	todoId := mux.Vars(r)["id"]
	
	result, err := controllers.GetTodo(todoId)

	if result.ID == nil {
		json.NewEncoder(w).Encode(models.ErrResponse{
			Success: false,
			Error: "Todo with the id of " + todoId + " does not exist",
		})
		return
	}

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
	todoId := mux.Vars(r)["id"]

	var todo models.Todo

	result, err := controllers.GetTodo(todoId)

	if result.ID == nil {
		json.NewEncoder(w).Encode(models.ErrResponse{
			Success: false,
			Error: "Todo with the id of " + todoId + " does not exist",
		})
		return
	}

	if err != nil {
		log.Fatalf("failed to get todo %v", err)
	}

	title := helpers.UpdateFieldBasedOfValuePresence(todo.Title, result.Title)
	description := helpers.UpdateFieldBasedOfValuePresence(todo.Description, result.Description)

	todo.Title = title
	todo.Description = description

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		isValidated := helpers.ValidateReuestBodyForUpdate(todo.Title, todo.Description)
		if !isValidated {
			json.NewEncoder(w).Encode(models.ErrResponse{
				Success: false,
				Error: "Please provide at least one field to update",
			})
		} else {
			json.NewEncoder(w).Encode(models.ErrResponse{
				Success: false,
				Error: "Something went wrong, please try again",
			})
			log.Fatalf("failed to decode json body %v", err)
		}
		return 
	}

	updatedRows, err := controllers.UpdateTodo(todoId, todo)
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
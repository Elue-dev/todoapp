package router

import (
	"github.com/elue-dev/todoapi/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/todos", handlers.GetAllTodos).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todos/{id}", handlers.GetSingleTodo).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/todos", handlers.AddTodo).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/todos/{id}", handlers.UpdateTodo).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/todos/{id}", handlers.DeleteTodo).Methods("DELETE", "OPTIONS")

	return router
}
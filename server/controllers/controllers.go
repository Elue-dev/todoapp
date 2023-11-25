package controllers

import (
	"log"

	"github.com/elue-dev/todoapi/connections"
	"github.com/elue-dev/todoapi/models"
)

func CreateTodo(t models.Todo) models.Todo {
	db := connections.CeateConnection()
	defer db.Close()

	sqlQuery := `INSERT INTO todos (title, description) VALUES ($1, $2) RETURNING *`
	var todo models.Todo

	err := db.QueryRow(sqlQuery, t.Title, t.Description).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		log.Fatalf("Could not execute SQL query %v", err)
	}
	
	return todo
}

func GetTodos() ([]models.Todo, error) {
	db := connections.CeateConnection()
	defer db.Close()

	sqlQuery := `SELECT * FROM todos`
	rows, err := db.Query(sqlQuery)

	if err != nil {
		log.Fatalf("Could not execute SQL query %v", err)
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, todo.CreatedAt, todo.UpdatedAt)
		if err != nil {
			log.Fatalf("Could not scan rows %v", err)
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func GetTodo(s models.Todo) {

}

func UpdateTodo(s models.Todo)  {

}

func DeleteTodo(s models.Todo)  {

}
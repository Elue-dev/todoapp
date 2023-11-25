package controllers

import (
	"database/sql"
	"fmt"
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
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			log.Fatalf("Could not scan rows %v", err)
		}
		todos = append(todos, todo)
	}

	return todos, err
}

func GetTodo(todoId string) (models.Todo, error) {
	db := connections.CeateConnection()
	defer db.Close()

	var todo models.Todo

	sqlQuery := `SELECT * FROM todos WHERE id = $1`
	rows := db.QueryRow(sqlQuery, todoId)
	
	err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)

	switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned.")
			return todo, nil
		case nil:
			return todo, nil
		default:
			log.Fatalf("Unable to scan rows %v", err)
	}

	return todo, err
}

func UpdateTodo(todoId string, t models.Todo) (int64, error)  {
	db := connections.CeateConnection()
	defer db.Close()


	sqlQuery := `UPDATE todos SET title = $2, description = $3, iscompleted = $4 WHERE id = $1`
	res, err := db.Exec(sqlQuery, todoId, t.Title, t.Description, t.IsCompleted)

	if err != nil {
		log.Fatalf("Could not execute query %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Could not get affected rows %v", err)
	}

	fmt.Printf("Total rows affected %v", rowsAffected)

	return rowsAffected, err
}

func DeleteTodo(s models.Todo)  {

}
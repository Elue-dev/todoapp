package controllers

import (
	"github.com/elue-dev/todoapi/connections"
	"github.com/elue-dev/todoapi/models"
)

func CreateTodo(todo models.Todo) models.Todo {
	db := connections.CeateConnection()
	defer db.Close()


}

func GetTodos(s models.Todo) {

}

func GetTodo(s models.Todo) {

}

func UpdateTodo(s models.Todo)  {

}

func DeleteTodo(s models.Todo)  {

}
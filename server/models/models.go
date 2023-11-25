package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
}

type CustomTodo struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
}

type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}
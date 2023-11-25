package models

type Todo struct {
	ID *string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

type MsgResponse struct {
	Success bool `json:"success"`
	Message string `json:"string"`
}
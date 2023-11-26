package models

type Todo struct {
	ID *string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"is_completed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	ID *string `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Avatar string `json:"avatar"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data"`
}

type MsgResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type ErrResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}
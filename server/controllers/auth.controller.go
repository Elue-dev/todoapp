package controllers

import (
	"errors"
	"fmt"

	"github.com/elue-dev/todoapi/connections"
	"github.com/elue-dev/todoapi/models"
)

func RegisterUser(u models.User) (models.User, error) {
	db := connections.CeateConnection()
	defer db.Close()

	// var existingUser models.User
    // err := db.QueryRow("SELECT * FROM users WHERE email = $1", u.Email).
    //     Scan(&existingUser.ID, &existingUser.Email)

	// if err != nil {
	// 	return existingUser, errors.New("email already exists")
	//   } else if err != sql.ErrNoRows {
	// 	return existingUser, err
	//   }
	
	sqlQuery := `INSERT INTO users (username, email, password, avatar) VALUES ($1, $2, $3, $4) RETURNING *`
	var user models.User

	err := db.QueryRow(sqlQuery, u.Username, u.Email, u.Password, u.Avatar).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Avatar, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		fmt.Println("Failed to execute sql insert query", err)
		 return user, errors.New("something went wrong. please try again")
		// log.Fatalf("Failed to execute sql insert query %v", err)
	}

	return user, nil
}
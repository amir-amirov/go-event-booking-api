package models

import (
	"fmt"

	"github.com/amir-amirov/go-event-booking-api/db"
	"github.com/amir-amirov/go-event-booking-api/utils"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) Save() error {

	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id
	`
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %s", err.Error())
	}

	user.Password = hashedPassword

	err = db.DB.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to save user: %s", err.Error())
	}

	fmt.Println("User saved successfully")

	return nil
}

func (user *User) ValidateCredentials() error {
	query := `
		SELECT id, password
		FROM users
		WHERE email = $1
		`
	var retrievedPassword string

	err := db.DB.QueryRow(query, user.Email).Scan(&user.ID, &retrievedPassword)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %s", err.Error())
	}

	if !utils.CheckPasswordHash(retrievedPassword, user.Password) {
		return fmt.Errorf("invalid credentials")
	}

	return nil
}

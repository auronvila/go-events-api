package models

import (
	"errors"
	"github.com/golang-events-planning-backend/db"
	"github.com/golang-events-planning-backend/utils"
	"log"
	"time"
)

type User struct {
	Id       string `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type UserResponse struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func (user User) Save() error {
	query := `
INSERT INTO users (id,email,password) VALUES (?,?,?)
`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	_, err = statement.Exec(user.Id, user.Email, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

func GetAlUsers() ([]User, error) {
	var users []User
	query := `
SELECT * FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (user *User) ValidateCredentials() error {
	// Start timing the database query
	start := time.Now()

	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPass string
	err := row.Scan(&user.Id, &retrievedPass)
	if err != nil {
		log.Printf("Database query and scan took: %v", time.Since(start))
		return errors.New("invalid credentials")
	}

	log.Printf("Database query and scan took: %v", time.Since(start))

	// Start timing the password comparison
	start = time.Now()

	passwordIsValid := utils.ComparePassword(retrievedPass, user.Password)
	if !passwordIsValid {
		log.Printf("Password comparison took: %v", time.Since(start))
		return errors.New("invalid credentials")
	}

	log.Printf("Password comparison took: %v", time.Since(start))

	return nil
}

func (user User) CreateUserResponse(token string) UserResponse {
	return UserResponse{
		Id:          user.Id,
		Email:       user.Email,
		AccessToken: token,
	}
}

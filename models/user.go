package models

import (
	"database/sql"
	"errors"
	"fmt"
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

type UserSignInResponse struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type UserResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
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

func GetAlUsers() ([]UserResponse, error) {
	var users []UserResponse
	query := `
SELECT id,email FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserResponse
		err = rows.Scan(&user.Id, &user.Email)
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

func (user UserResponse) GetSingleUser() (UserResponse, error) {
	query := `
    SELECT id, email 
    FROM users 
    WHERE id = ?
    `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return UserResponse{}, fmt.Errorf("error preparing query: %v", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Id).Scan(&user.Id, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserResponse{}, fmt.Errorf("no user found with id: %v", user.Id)
		}
		return UserResponse{}, fmt.Errorf("error executing query: %v", err)
	}

	return UserResponse{
		Id:    user.Id,
		Email: user.Email,
	}, nil
}

func (user User) CreateUserResponse(token string) UserSignInResponse {
	return UserSignInResponse{
		Id:          user.Id,
		Email:       user.Email,
		AccessToken: token,
	}
}

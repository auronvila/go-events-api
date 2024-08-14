package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/models"
	"github.com/golang-events-planning-backend/utils"
	"log"
	"net/http"
	"time"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "user with this email already exists"})
		return
	}

	fmt.Println(user)
	context.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func GetaAllUsers(context *gin.Context) {
	users, err := models.GetAlUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error in fetching the users from the database"})
	}

	context.JSON(http.StatusOK, users)
}

func Login(context *gin.Context) {
	start := time.Now()

	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	log.Printf("JSON binding took: %v", time.Since(start))

	start = time.Now()
	if err := user.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}
	log.Printf("Credential validation took: %v", time.Since(start))

	start = time.Now()
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error in creating the jwt token"})
		return
	}
	log.Printf("Token generation took: %v", time.Since(start))

	context.JSON(http.StatusOK, user.CreateUserResponse(token))
}

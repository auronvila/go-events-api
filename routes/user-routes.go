package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/routes/services"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/users/signup", services.SignUp)
	server.GET("/users", services.GetaAllUsers)
	server.POST("/users/login", services.Login)
}

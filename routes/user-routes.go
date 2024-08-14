package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/middlewares"
	"github.com/golang-events-planning-backend/routes/services"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/users/signup", services.SignUp)
	server.POST("/users/login", services.Login)
	server.GET("/users", middlewares.Authenticate, services.GetaAllUsers)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/middlewares"
	"github.com/golang-events-planning-backend/routes/services"
)

func RegisterCommentRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/comments/:id", services.CreateComment)
	authenticated.GET("/comments", services.GetComments)
}

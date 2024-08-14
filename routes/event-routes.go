package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/routes/services"
)

func RegisterEventRoutes(server *gin.Engine) {
	server.GET("/events", services.GetEvents)
	server.GET("/event/:id", services.GetEventById)
	server.POST("/events", services.CreateEvent)
	server.PUT("/event/:id", services.UpdateEvent)
	server.DELETE("/event/:id", services.DeleteEvent)
}

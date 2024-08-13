package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEventById)
	server.POST("/events", createEvent)
	server.PUT("/event/:id", updateEvent)
}

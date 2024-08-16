package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/middlewares"
	"github.com/golang-events-planning-backend/routes/services"
)

func RegisterEventRoutes(server *gin.Engine) {
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", services.CreateEvent)
	authenticated.POST("/event/:id/register", services.RegisterUserToEvent)
	authenticated.DELETE("/event/:id/register", services.CancelRegistration)
	authenticated.PUT("/event/:id", services.UpdateEvent)
	authenticated.DELETE("/event/:id", services.DeleteEvent)
	authenticated.GET("/events", services.GetEvents)
	authenticated.GET("/event/:id", services.GetEventById)

	authenticated.GET("/events/userAssignedEvents", services.GetUserAssignedEvents)
	authenticated.GET("/event/:id/specificEventAssignedUser", services.GetSpecificEventUserAssignee)
}

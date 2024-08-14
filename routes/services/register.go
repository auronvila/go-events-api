package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/models"
	"github.com/golang-events-planning-backend/utils"
	"net/http"
)

func RegisterUserToEvent(context *gin.Context) {
	userId := context.GetString("userId")
	eventId := context.Param("id")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	id := utils.GenerateUUID()
	err = event.Register(userId, id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "successfully registered the user to the event"})
}

func CancelRegistration(context *gin.Context) {
	userId := context.GetString("userId")
	eventId := context.Param("id")

	var event models.Event
	event.Id = eventId
	err := event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errMsg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"errMsg": "successfully canceled from event"})
}

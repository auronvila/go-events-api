package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/models"
	"net/http"
	"strconv"
)

func RegisterUserToEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errMsg": "could not parse the eventId"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "successfully registered the user to the event"})
}

func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errMsg": "could not parse the eventId"})
		return
	}

	var event models.Event
	event.Id = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errMsg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"errMsg": "successfully canceled from event"})
}

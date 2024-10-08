package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/models"
	"github.com/golang-events-planning-backend/utils"
	"net/http"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the events from the database"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEventById(context *gin.Context) {
	eventId := context.Param("id")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
	return
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	event.Id = utils.GenerateUUID()
	userId := context.GetString("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func UpdateEvent(context *gin.Context) {
	userId := context.GetString("userId")
	eventId := context.Param("id")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in getting the event"})
		return
	}

	if event.UserId != userId {
		context.JSON(http.StatusBadRequest, gin.H{"message": "only the owner of the event can update the event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	updatedEvent.Id = eventId

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func DeleteEvent(context *gin.Context) {
	userId := context.GetString("userId")
	eventId := context.Param("id")

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event with the given id was not found"})
		return
	}
	if event.UserId != userId {
		context.JSON(http.StatusBadRequest, gin.H{"message": "only the owner of the event can delete the event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}

func GetUserAssignedEvents(context *gin.Context) {
	res, err := models.GetUserAssignedRegistrations()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, res)
}

func GetSpecificEventUserAssignee(context *gin.Context) {
	eventId := context.Param("id")

	fmt.Println(eventId)
	res, err := models.GetSpecificEventUserAssignee(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, res)
}

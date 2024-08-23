package services

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-events-planning-backend/models"
	"github.com/golang-events-planning-backend/utils"
	"net/http"
)

func CreateComment(context *gin.Context) {
	var comment models.Comment
	err := context.ShouldBindJSON(&comment)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	eventId := context.Param("id")
	userId := context.GetString("userId")
	comment.Id = utils.GenerateUUID()
	comment.UserId = userId
	comment.EventId = eventId

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errMsg": err.Error()})
		return
	}

	err = comment.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"errMsg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": comment})
}

func GetComments(context *gin.Context) {
	data, err := models.GetAllComments()
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, data)
}

func GetCommentById(context *gin.Context) {
	eventId := context.Param("id")
	comments, err := models.GetAllCommentsByEventId(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"errMsg": err.Error()})
		return
	}

	context.JSON(http.StatusOK, comments)
}

package routes

import (
	"net/http"
	"strconv"

	"github.com/amir-amirov/go-event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body" + err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "data": event})
}

func getEvents(context *gin.Context) {
	var events []models.Event

	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve events"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": events})

}

func deleteEvent(context *gin.Context) {
	eventId := context.Param("id")
	id, err := strconv.Atoi(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	err = models.Delete(int64(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to delete event" + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}

func updateEvent(context *gin.Context) {
	eventId := context.Param("id")

	id, err := strconv.Atoi(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid body" + err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	err = event.Update(int64(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to delete event" + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event update successfully", "data": event})
}

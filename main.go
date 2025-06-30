package main

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events")
	// Make sure that the server listens to the 8080 port
	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request Data"})
		return
	}
	// Dummy values for now
	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"messagge": "Event created", "event": event})
}

// Package routes: establishes which endpoint to serve and how
package routes

import (
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", RegisterForEvent)
	authenticated.DELETE("/events/:id/register", CancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}

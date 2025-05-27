package routes

import (
	"github.com/amir-amirov/go-event-booking-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.POST("/register", createUser)
	server.POST("/login", loginUser)

	server.GET("events", getEvents)

	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("", createEvent)
	authenticated.DELETE("/:id", deleteEvent)
	authenticated.PUT("/:id", updateEvent)

}

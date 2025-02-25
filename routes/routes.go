package routes

import (
	"example.com/api-rest/middlewares"
	"github.com/gin-gonic/gin"
)

// Here is the archive for the routes and endpoints
func RegisterRoutes(server *gin.Engine) {
	//Route for getting all events, no authentication needed
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register",cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
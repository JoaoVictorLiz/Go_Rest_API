package main

import (
	"example.com/api-rest/db"
	"example.com/api-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() //Start the DB connection
	server := gin.Default()
	
	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080

}

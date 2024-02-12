package main

import (
	"example.com/beers/db"
	"example.com/beers/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/beers", getBeers)
	server.GET("/beers/:id", getBeer)
	server.POST("/beers", createBeer)
	server.PUT("/beers/:id", updateBeer)
	server.DELETE("/beers/:id", deleteBeer)
}

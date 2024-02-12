package routes

import (
	"net/http"
	"strconv"

	"example.com/beers/models"
	"github.com/gin-gonic/gin"
)

func createBeer(context *gin.Context) {

	var beer models.Beer
	err := context.ShouldBindJSON(&beer)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requst data"})
		return
	}

	err = beer.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create beer.Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "beer created!", "beer": beer})
}

func getBeers(context *gin.Context) {

	beers, err := models.GetAllBeers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch beers.Try again later", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, beers)
}

func getBeer(context *gin.Context) {
	id := context.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse beer id"})
		return
	}

	beer, err := models.GetBeerById(parsedId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch beer"})
		return
	}

	context.JSON(http.StatusOK, beer)
}

func deleteBeer(context *gin.Context) {
	id := context.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse beer id"})
		return
	}
	beer, err := models.GetBeerById(parsedId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch beer"})
		return
	}

	err = beer.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the beer"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Beer deleted successfully"})
}

func updateBeer(context *gin.Context) {
	id := context.Param("id")
	parsedId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse beer id"})
		return
	}
	_, err = models.GetBeerById(parsedId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch beer"})
		return
	}
	var updatedBeer models.Beer
	err = context.ShouldBindJSON(&updatedBeer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	updatedBeer.ID = parsedId

	err = updatedBeer.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update beer"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Beer updated successfully"})
}

package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PokerSquareIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.PokerSquarePaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.PokerSquare{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Score ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

func PokerSquareById(c *gin.Context) {
	// get id
	id := c.Param("id")

	pokerSquare := models.PokerSquare{}
	initializers.DB.Preload("User").First(&pokerSquare, id)

	// response
	c.JSON(http.StatusOK, pokerSquare.Json())
}

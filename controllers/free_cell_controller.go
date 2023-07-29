package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FreeCellIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.FreeCellPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Find(&models.FreeCell{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Moves ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response)
}

func FreeCellById(c *gin.Context) {
	// get id
	id := c.Param("id")

	freeCell := models.FreeCell{}
	initializers.DB.Preload("User").First(&freeCell, id)

	// response
	c.JSON(http.StatusOK, freeCell)
}

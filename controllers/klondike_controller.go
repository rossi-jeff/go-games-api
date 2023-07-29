package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func KlondikeIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.KlondikePaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.Klondike{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Moves ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response)
}

func KlondikeById(c *gin.Context) {
	// get id
	id := c.Param("id")

	klondike := models.Klondike{}
	initializers.DB.Preload("User").First(&klondike, id)

	// response
	c.JSON(http.StatusOK, klondike)
}

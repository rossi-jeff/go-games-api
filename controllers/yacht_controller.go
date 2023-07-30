package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func YachtIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.YachtPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("NumTurns = 12").Find(&models.Yacht{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("NumTurns = 12").Order("Total DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

func YachtById(c *gin.Context) {
	// get id
	id := c.Param("id")

	yacht := models.Yacht{}
	initializers.DB.Preload("User").First(&yacht, id)

	// response
	c.JSON(http.StatusOK, yacht.Json())
}

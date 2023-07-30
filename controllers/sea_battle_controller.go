package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func SeaBattleIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.SeaBattlePaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.SeaBattle{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

func SeaBattleById(c *gin.Context) {
	// get id
	id := c.Param("id")

	seaBattle := models.SeaBattle{}
	initializers.DB.Preload("Ships.Points").Preload("Ships.Hits").Preload(clause.Associations).First(&seaBattle, id)

	// response
	c.JSON(http.StatusOK, seaBattle.Json())
}

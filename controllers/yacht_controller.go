package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Yacht
// @Description  paginated list of yacht
// @Tags         Yacht
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.YachtPaginated
// @Router       /api/yacht [get]
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

// @Summary      Yacht by Id
// @Description  get a yacht
// @Tags         Yacht
// @Accept       json
// @Produce      json
// @Param Id path int true "Yacht ID"
// @Success      200  {object} models.Yacht
// @Router       /api/yacht/{Id} [get]
func YachtById(c *gin.Context) {
	// get id
	id := c.Param("id")

	yacht := models.Yacht{}
	initializers.DB.Preload("User").Preload("Turns").First(&yacht, id)

	// response
	c.JSON(http.StatusOK, yacht.Json())
}

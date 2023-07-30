package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Hang Man
// @Description  paginated list of hang man
// @Tags         Hang Man
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.HangManPaginated
// @Router       /api/hang_man [get]
func HangManIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.HangManPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.HangMan{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Preload("Word").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Hang Man by Id
// @Description  get a hang man
// @Tags         Hang Man
// @Accept       json
// @Produce      json
// @Param Id path int true "Hang Man ID"
// @Success      200  {object} models.HangMan
// @Router       /api/hang_man/{Id} [get]
func HangManById(c *gin.Context) {
	// get id
	id := c.Param("id")

	hangMan := models.HangMan{}
	initializers.DB.Preload("User").Preload("Word").First(&hangMan, id)

	// response
	c.JSON(http.StatusOK, hangMan.Json())
}

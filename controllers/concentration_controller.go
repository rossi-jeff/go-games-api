package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Concentration
// @Description  paginated list of concentation
// @Tags         Concentration
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.ConcentrationPaginated
// @Router       /api/concentration [get]
func ConcentrationIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.ConcentrationPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Find(&models.Concentration{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Moves ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Concentration by Id
// @Description  get a concentation
// @Tags         Concentration
// @Accept       json
// @Produce      json
// @Param Id path int true "Concentration ID"
// @Success      200  {object} models.Concentration
// @Router       /api/concentration/{Id} [get]
func ConcentrationById(c *gin.Context) {
	// get id
	id := c.Param("id")

	concentration := models.Concentration{}
	initializers.DB.Preload("User").First(&concentration, id)

	// response
	c.JSON(http.StatusOK, concentration.Json())
}

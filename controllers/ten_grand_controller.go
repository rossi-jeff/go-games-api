package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// @Summary      Ten Grand
// @Description  paginated list of ten grand
// @Tags         Ten Grand
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.TenGrandPaginated
// @Router       /api/ten_grand [get]
func TenGrandIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.TenGrandPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.TenGrand{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Ten Grand by Id
// @Description  get a ten grand
// @Tags         Ten Grand
// @Accept       json
// @Produce      json
// @Param Id path int true "Ten Grand ID"
// @Success      200  {object} models.TenGrand
// @Router       /api/ten_grand/{Id} [get]
func TenGrandById(c *gin.Context) {
	// get id
	id := c.Param("id")

	tenGrand := models.TenGrand{}
	initializers.DB.Preload("Turns.Scores").Preload(clause.Associations).First(&tenGrand, id)

	// response
	c.JSON(http.StatusOK, tenGrand.Json())
}
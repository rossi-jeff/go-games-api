package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Spider
// @Description  paginated list of spider
// @Tags         Spider
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.SpiderPaginated
// @Router       /api/spider [get]
func SpiderIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.SpiderPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.Spider{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Moves ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Spider by Id
// @Description  get a spider
// @Tags         Spider
// @Accept       json
// @Produce      json
// @Param Id path int true "Spider ID"
// @Success      200  {object} models.Spider
// @Router       /api/spider/{Id} [get]
func SpiderById(c *gin.Context) {
	// get id
	id := c.Param("id")

	spider := models.Spider{}
	initializers.DB.Preload("User").First(&spider, id)

	// response
	c.JSON(http.StatusOK, spider.Json())
}

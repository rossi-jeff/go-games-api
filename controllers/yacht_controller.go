package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary      Yacht
// @Description  paginated list of yacht
// @Tags         Yacht
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.YachtPaginatedJson
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
// @Success      200  {object} models.YachtJson
// @Router       /api/yacht/{Id} [get]
func YachtById(c *gin.Context) {
	// get id
	id := c.Param("id")

	yacht := models.Yacht{}
	initializers.DB.Preload("User").Preload("Turns").First(&yacht, id)

	// response
	c.JSON(http.StatusOK, yacht.Json())
}

func YachtCreate(c *gin.Context) {
	now := time.Now().Format(time.RFC3339)
	yacht := models.Yacht{}
	yacht.CreatedAt = now
	yacht.UpdatedAt = now
	initializers.DB.Save(&yacht)

	c.JSON(http.StatusCreated, yacht.Json())
}

func YachtRoll(c *gin.Context) {
	params := payloads.YachtRollPayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	yacht := models.Yacht{}
	initializers.DB.First(&yacht, id)
}

func YachtScore(c *gin.Context) {
	params := payloads.YachtScorePayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	yacht := models.Yacht{}
	initializers.DB.First(&yacht, id)
}

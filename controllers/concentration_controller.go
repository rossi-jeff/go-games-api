package controllers

import (
	"go-games-api/enum"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary      Concentration
// @Description  paginated list of concentation
// @Tags         Concentration
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.ConcentrationPaginatedJson
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
// @Success      200  {object} models.ConcentrationJson
// @Router       /api/concentration/{Id} [get]
func ConcentrationById(c *gin.Context) {
	// get id
	id := c.Param("id")

	concentration := models.Concentration{}
	initializers.DB.Preload("User").First(&concentration, id)

	// response
	c.JSON(http.StatusOK, concentration.Json())
}

// @Summary      Create Concentration
// @Description  returns a new  concentation
// @Tags         Concentration
// @Accept       json
// @Produce      json
// @Success      201  {object} models.ConcentrationJson
// @Router       /api/concentration [post]
func ConcentrationCreate(c *gin.Context) {
	now := time.Now().Format(time.RFC3339)
	concentration := models.Concentration{}
	concentration.Status = enum.Playing
	concentration.CreatedAt = now
	concentration.UpdatedAt = now
	initializers.DB.Save(&concentration)

	c.JSON(http.StatusCreated, concentration.Json())
}

// @Summary      Update Concentration
// @Description  update a concentation
// @Tags         Concentration
// @Accept       json
// @Produce      json
// @Param Id path int true "Concentration ID"
// @Param	data	body	payloads.ConcentrationUpdatePayload		true	"Concentration Updates"
// @Success      200  {object} models.ConcentrationJson
// @Router       /api/concentration/{Id} [patch]
func ConcentrationUpdate(c *gin.Context) {
	params := payloads.ConcentrationUpdatePayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	concentration := models.Concentration{}
	initializers.DB.Select("id,created_at,user_id").Find(&concentration, id)

	concentration.Moves = params.Moves
	concentration.Elapsed = params.Elapsed
	concentration.Matched = params.Matched
	concentration.Status = enum.GameStatus(enum.GameStatusArrayIndex(string(params.Status)))
	concentration.UpdatedAt = time.Now().Format(time.RFC3339)

	initializers.DB.Preload("User").Save(&concentration)

	c.JSON(http.StatusOK, concentration.Json())
}

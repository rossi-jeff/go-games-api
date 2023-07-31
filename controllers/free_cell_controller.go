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

// @Summary      Free Cell
// @Description  paginated list of free cell
// @Tags         Free Cell
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.FreeCellPaginated
// @Router       /api/free_cell [get]
func FreeCellIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.FreeCellPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Find(&models.FreeCell{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Moves ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Free Cell by Id
// @Description  get a free cell
// @Tags         Free Cell
// @Accept       json
// @Produce      json
// @Param Id path int true "Free Cell ID"
// @Success      200  {object} models.FreeCell
// @Router       /api/free_cell/{Id} [get]
func FreeCellById(c *gin.Context) {
	// get id
	id := c.Param("id")

	freeCell := models.FreeCell{}
	initializers.DB.Preload("User").First(&freeCell, id)

	// response
	c.JSON(http.StatusOK, freeCell.Json())
}

// @Summary      Create Free Cell
// @Description  create a free cell
// @Tags         Free Cell
// @Accept       json
// @Produce      json
// @Success      201  {object} models.FreeCellPaginated
// @Router       /api/free_cell [post]
func FreeCellCreate(c *gin.Context) {
	now := time.Now().Format(time.RFC3339)
	freeCell := models.FreeCell{}
	freeCell.Status = enum.Playing
	freeCell.CreatedAt = now
	freeCell.UpdatedAt = now
	initializers.DB.Save(&freeCell)
	c.JSON(http.StatusOK, freeCell.Json())
}

// @Summary      Update Free Cell
// @Description  update a free cell
// @Tags         Free Cell
// @Accept       json
// @Produce      json
// @Param Id path int true "Free Cell ID"
// @Param	data	body	payloads.FreeCellUpdatePayload		true	"Free Cell Updates"
// @Success      200  {object} models.FreeCell
// @Router       /api/free_cell/{Id} [patch]
func FreeCellUpdate(c *gin.Context) {
	params := payloads.FreeCellUpdatePayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	freeCell := models.FreeCell{}
	initializers.DB.Select("id,created_at,user_id").First(&freeCell, id)

	freeCell.Elapsed = params.Elapsed
	freeCell.Moves = params.Moves
	freeCell.Status = enum.GameStatus(enum.GameStatusArrayIndex(string(params.Status)))
	freeCell.UpdatedAt = time.Now().Format(time.RFC3339)

	initializers.DB.Save(&freeCell)

	c.JSON(http.StatusOK, freeCell.Json())
}

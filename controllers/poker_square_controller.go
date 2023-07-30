package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary      Poker Square
// @Description  paginated list of poker square
// @Tags         Poker Square
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.PokerSquarePaginated
// @Router       /api/poker_square [get]
func PokerSquareIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.PokerSquarePaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.PokerSquare{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Score ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Poker Square by Id
// @Description  get a poker square
// @Tags         Poker Square
// @Accept       json
// @Produce      json
// @Param Id path int true "Poker Square ID"
// @Success      200  {object} models.PokerSquare
// @Router       /api/poker_square/{Id} [get]
func PokerSquareById(c *gin.Context) {
	// get id
	id := c.Param("id")

	pokerSquare := models.PokerSquare{}
	initializers.DB.Preload("User").First(&pokerSquare, id)

	// response
	c.JSON(http.StatusOK, pokerSquare.Json())
}

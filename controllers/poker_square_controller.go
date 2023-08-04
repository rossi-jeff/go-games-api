package controllers

import (
	"database/sql"
	"go-games-api/enum"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary      Poker Square
// @Description  paginated list of poker square
// @Tags         Poker Square
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.PokerSquarePaginatedJson
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
// @Success      200  {object} models.PokerSquareJson
// @Router       /api/poker_square/{Id} [get]
func PokerSquareById(c *gin.Context) {
	// get id
	id := c.Param("id")

	pokerSquare := models.PokerSquare{}
	initializers.DB.Preload("User").First(&pokerSquare, id)

	// response
	c.JSON(http.StatusOK, pokerSquare.Json())
}

// @Summary      Create Poker Square
// @Description  create a  poker square
// @Tags         Poker Square
// @Accept       json
// @Produce      json
// @Success      200  {object} models.PokerSquareJson
// @Router       /api/poker_square [post]
func PokerSquareCreate(c *gin.Context) {
	now := time.Now().Format(time.RFC3339)
	userId := utilities.UserIdFromAuthHeader(c)

	pokerSquare := models.PokerSquare{}
	if userId > 0 {
		pokerSquare.UserId = sql.NullInt64{Int64: int64(userId), Valid: true}
	}
	pokerSquare.Status = enum.Playing
	pokerSquare.CreatedAt = now
	pokerSquare.UpdatedAt = now
	initializers.DB.Save(&pokerSquare)

	c.JSON(http.StatusCreated, pokerSquare.Json())
}

// @Summary      Update PokerSquare
// @Description  get a PokerSquare
// @Tags         Poker Square
// @Accept       json
// @Produce      json
// @Param Id path int true "PokerSquare ID"
// @Param	data	body	payloads.PokerSquareUpdatePayload		true	"Poker Square Updates"
// @Success      200  {object} models.PokerSquareJson
// @Router       /api/poker_square/{Id} [patch]
func PokerSquareUpdate(c *gin.Context) {
	params := payloads.PokerSquareUpdatePayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	pokerSquare := models.PokerSquare{}
	initializers.DB.Select("id,created_at,user_id").First(&pokerSquare, id)
	pokerSquare.UpdatedAt = now
	pokerSquare.Status = enum.GameStatus(enum.GameStatusArrayIndex(string(params.Status)))
	pokerSquare.Score = params.Score
	initializers.DB.Save(&pokerSquare)

	c.JSON(http.StatusOK, pokerSquare.Json())
}

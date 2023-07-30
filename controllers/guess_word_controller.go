package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// @Summary      Guess Word
// @Description  paginated list of guess word
// @Tags         Guess Word
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.GuessWordPaginated
// @Router       /api/guess_word [get]
func GuessWordIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.GuessWordPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.GuessWord{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Preload("Word").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Guess Word by Id
// @Description  get a code guess word
// @Tags         Guess Word
// @Accept       json
// @Produce      json
// @Param Id path int true "Guess Word ID"
// @Success      200  {object} models.GuessWord
// @Router       /api/guess_word/{Id} [get]
func GuessWordById(c *gin.Context) {
	// get id
	id := c.Param("id")

	guessWord := models.GuessWord{}
	initializers.DB.Preload("Guesses.Ratings").Preload(clause.Associations).First(&guessWord, id)

	// response
	c.JSON(http.StatusOK, guessWord.Json())
}

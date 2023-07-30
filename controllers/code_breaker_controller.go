package controllers

import (
	"go-games-api/enum"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// @Summary      Code Breakers
// @Description  paginated list of code breakers
// @Tags         Code Breaker
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.CodeBreakerPaginated
// @Router       /api/code_breaker [get]
func CodeBreakerIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.CodeBreakerPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.CodeBreaker{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Code Breaker by Id
// @Description  get a code breaker
// @Tags         Code Breaker
// @Accept       json
// @Produce      json
// @Param Id path int true "Code Breaker ID"
// @Success      200  {object} models.CodeBreaker
// @Router       /api/code_breaker/{Id} [get]
func CodeBreakerById(c *gin.Context) {
	// get id
	id := c.Param("id")

	// find code breaker
	codeBreaker := models.CodeBreaker{}
	initializers.DB.Preload("Guesses.Colors").Preload("Guesses.Keys").Preload(clause.Associations).First(&codeBreaker, id)

	// response
	c.JSON(http.StatusOK, codeBreaker.Json())
}

// @Summary      Create Code Breaker
// @Description  Create a code breaker and the secret code
// @Tags         Code Breaker
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.CodeBreakerCreatePayload		true	"Create Code Breaker Options"
// @Success      201  {object} models.CodeBreaker
// @Router       /api/code_breaker [post]
func CodeBreakerCreate(c *gin.Context) {
	params := payloads.CodeBreakerCreatePayload{}

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	codeBreaker := models.CodeBreaker{}

	now := time.Now().Format(time.RFC3339)
	codeBreaker.Columns = params.Columns
	codeBreaker.Available = strings.Join(params.Colors, ",")
	codeBreaker.Colors = len(params.Colors)
	codeBreaker.CreatedAt = now
	codeBreaker.UpdatedAt = now
	codeBreaker.Status = enum.Playing

	initializers.DB.Save(&codeBreaker)

	if params.Columns > 0 {
		for i := 0; i < params.Columns; i++ {
			idx := rand.Intn(len(params.Colors))
			color := params.Colors[idx]

			code := models.CodeBreakerCode{}
			code.CodeBreakerId = int64(codeBreaker.Id)
			code.Color = enum.Color(enum.ColorArrayIndex(color))
			code.CreatedAt = now
			code.UpdatedAt = now
			initializers.DB.Save(&code)
		}
	}

	// look up new code breaker for codes
	found := models.CodeBreaker{}
	initializers.DB.Preload("Codes").First(&found, codeBreaker.Id)

	// response
	c.JSON(http.StatusOK, found.Json())
}

// @Summary      Code Breaker Guess
// @Description  attempt to guess the code breaker code
// @Tags         Code Breaker
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.CodeBreakerGuessPayload		true	"Code Breaker Guess"
// @Param Id path int true "Code Breaker ID"
// @Success      201  {object} models.CodeBreakerGuess
// @Router       /api/code_breaker/{Id}/guess [post]
func CodeBreakerGuess(c *gin.Context) {
	params := payloads.CodeBreakerGuessPayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	codeBreaker := models.CodeBreaker{}
	initializers.DB.Preload("Codes").Preload("Guesses").Find(&codeBreaker, id)

	now := time.Now().Format(time.RFC3339)

	guess := models.CodeBreakerGuess{}
	guess.CodeBreakerId = int64(codeBreaker.Id)
	guess.CreatedAt = now
	guess.UpdatedAt = now

	initializers.DB.Save(&guess)

	CalculateGuessRatings(c, codeBreaker, guess, GetCodeColors(codeBreaker), params.Colors)
}

func GetCodeColors(codeBreaker models.CodeBreaker) []string {
	var colors []string
	for i := 0; i < len(codeBreaker.Codes); i++ {
		color := codeBreaker.Codes[i].Color.String()
		colors = append(colors, color)
	}
	return colors
}

func CalculateGuessRatings(c *gin.Context, codeBreaker models.CodeBreaker, guess models.CodeBreakerGuess, code []string, colors []string) {
	if len(code) != len(colors) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect number of colors"})
		return
	}

	now := time.Now().Format(time.RFC3339)
	black := 0
	white := 0

	for i := 0; i < len(colors); i++ {
		guessColor := models.CodeBreakerGuessColor{}
		guessColor.CodeBreakerGuessId = int64(guess.Id)
		guessColor.CreatedAt = now
		guessColor.UpdatedAt = now
		guessColor.Color = enum.Color(enum.ColorArrayIndex(colors[i]))
		initializers.DB.Save(&guessColor)
	}

	for i := len(code) - 1; i >= 0; i-- {
		if code[i] == colors[i] {
			black++
			code = DeleteStringSliceIndex(code, i)
			colors = DeleteStringSliceIndex(colors, i)
		}
	}

	for i := 0; i < len(colors); i++ {
		idx := SliceIndexOf(colors[i], code)
		if idx != -1 {
			white++
			code = DeleteStringSliceIndex(code, idx)
		}
	}

	for i := 0; i < black; i++ {
		guessKey := models.CodeBreakerGuessKey{}
		guessKey.CodeBreakerGuessId = int64(guess.Id)
		guessKey.CreatedAt = now
		guessKey.UpdatedAt = now
		guessKey.Key = enum.BLACK
		initializers.DB.Save(&guessKey)
	}
	for i := 0; i < white; i++ {
		guessKey := models.CodeBreakerGuessKey{}
		guessKey.CodeBreakerGuessId = int64(guess.Id)
		guessKey.CreatedAt = now
		guessKey.UpdatedAt = now
		guessKey.Key = enum.WHITE
		initializers.DB.Save(&guessKey)
	}

	UpdateCodeBreakerStatus(codeBreaker, black)

	found := models.CodeBreakerGuess{}
	initializers.DB.Preload("Colors").Preload("Keys").First(&found, guess.Id)

	c.JSON(http.StatusOK, found.Json())
}

func UpdateCodeBreakerStatus(codeBreaker models.CodeBreaker, black int) {
	update := models.CodeBreaker{
		Status: enum.Playing,
	}

	if black == codeBreaker.Columns {
		update.Status = enum.Won
	} else if len(codeBreaker.Guesses)+1 > codeBreaker.Columns*2 {
		update.Status = enum.Lost
	}

	if update.Status != enum.Playing {
		initializers.DB.Model(&codeBreaker).Updates(models.CodeBreaker{Status: update.Status})
	}
}

func DeleteStringSliceIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func SliceIndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

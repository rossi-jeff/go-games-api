package controllers

import (
	"database/sql"
	"go-games-api/enum"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// @Summary      Hang Man
// @Description  paginated list of hang man
// @Tags         Hang Man
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.HangManPaginatedJson
// @Router       /api/hang_man [get]
func HangManIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.HangManPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.HangMan{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Preload("Word").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Hang Man by Id
// @Description  get a hang man
// @Tags         Hang Man
// @Accept       json
// @Produce      json
// @Param Id path int true "Hang Man ID"
// @Success      200  {object} models.HangManJson
// @Router       /api/hang_man/{Id} [get]
func HangManById(c *gin.Context) {
	// get id
	id := c.Param("id")

	hangMan := models.HangMan{}
	initializers.DB.Preload("User").Preload("Word").First(&hangMan, id)

	// response
	c.JSON(http.StatusOK, hangMan.Json())
}

// @Summary      Create Hang Man
// @Description  create a hang man
// @Tags         Hang Man
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.HangManCreatePayload		true	"Hang Man Options"
// @Success      200  {object} models.HangManJson
// @Router       /api/hang_man [post]
func HangManCreate(c *gin.Context) {
	params := payloads.HangManCreatePayload{}
	userId := utilities.UserIdFromAuthHeader(c)

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	word := models.Word{}
	initializers.DB.First(&word, params.WordId)

	now := time.Now().Format(time.RFC3339)
	hangMan := models.HangMan{}
	if userId > 0 {
		hangMan.UserId = sql.NullInt64{Int64: int64(userId), Valid: true}
	}
	hangMan.WordId = sql.NullInt64{Int64: int64(word.Id), Valid: true}
	hangMan.CreatedAt = now
	hangMan.UpdatedAt = now
	hangMan.Status = enum.Playing

	initializers.DB.Save(&hangMan)

	hangMan.Word = word

	c.JSON(http.StatusCreated, hangMan.Json())
}

// @Summary      Hang Man Guess
// @Description  guess a letter in hang man
// @Tags         Hang Man
// @Accept       json
// @Produce      json
// @Param Id path int true "Hang Man ID"
// @Param	data	body	payloads.HangManGuessPayload		true	"Hang Man Options"
// @Success      200  {object} payloads.HangManGuessResponse
// @Router       /api/hang_man/{Id}/guess [post]
func HangManGuess(c *gin.Context) {
	params := payloads.HangManGuessPayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hangMan := models.HangMan{}
	initializers.DB.First(&hangMan, id)

	word := strings.Split(params.Word, "")
	correct := strings.Split(hangMan.Correct, ",")
	wrong := strings.Split(hangMan.Wrong, ",")

	Found := false
	if utilities.StringSliceIndexOf(params.Letter, word) != -1 {
		Found = true
		correct = append(correct, params.Letter)
	} else {
		wrong = append(wrong, params.Letter)
	}

	hangMan.Correct = strings.Join(utilities.StringSliceUnique(correct), ",")
	hangMan.Wrong = strings.Join(utilities.StringSliceUnique(wrong), ",")
	hangMan.Status = hangManStatus(word, correct, wrong)
	if hangMan.Status != enum.Playing {
		hangMan.Score = hangManScore(word, correct, wrong)
	}
	initializers.DB.Save(&hangMan)

	response := payloads.HangManGuessResponse{
		Found:  Found,
		Letter: params.Letter,
	}

	c.JSON(http.StatusOK, response)
}

func hangManStatus(word []string, correct []string, wrong []string) enum.GameStatus {
	unique := utilities.StringSliceUnique(word)
	var missed []string
	for i := 0; i < len(unique); i++ {
		letter := unique[i]
		if utilities.StringSliceIndexOf(letter, correct) == -1 {
			missed = append(missed, letter)
		}
	}
	if len(missed) == 0 {
		return enum.Won
	}
	if len(wrong) >= 6 {
		return enum.Lost
	}
	return enum.Playing
}

func hangManScore(word []string, correct []string, wrong []string) int {
	perLetter := 10
	perCorrect := 5
	score := len(utilities.StringSliceUnique(word)) * perLetter
	score = score + (len(correct) * perCorrect)
	score = score - (len(wrong) * perLetter)
	return score
}

func HangManInProgress(c *gin.Context) {
	userId := utilities.UserIdFromAuthHeader(c)
	hangMenJson := []models.HangManJson{}
	if userId > 0 {
		hangMen := []models.HangMan{}
		initializers.DB.Where("user_id = ? AND Status = 1", userId).Preload(clause.Associations).Find(&hangMen)
		for i := 0; i < len(hangMen); i++ {
			hangMan := hangMen[i].Json()
			hangMenJson = append(hangMenJson, hangMan)
		}
	}
	c.JSON(http.StatusOK, hangMenJson)
}

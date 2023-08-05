package controllers

import (
	"database/sql"
	"go-games-api/enum"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"math"
	"net/http"
	"strings"
	"time"

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
// @Success      200  {object} models.GuessWordPaginatedJson
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
// @Success      200  {object} models.GuessWordJson
// @Router       /api/guess_word/{Id} [get]
func GuessWordById(c *gin.Context) {
	// get id
	id := c.Param("id")

	guessWord := models.GuessWord{}
	initializers.DB.Preload("Guesses.Ratings").Preload(clause.Associations).First(&guessWord, id)

	// response
	c.JSON(http.StatusOK, guessWord.Json())
}

// @Summary      Create Guess Word
// @Description  create a  guess word
// @Tags         Guess Word
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.GuesssWordCreatePayload		true	"Guess Word Options"
// @Success      200  {object} models.GuessWordJson
// @Router       /api/guess_word [post]
func GuessWordCreate(c *gin.Context) {
	params := payloads.GuesssWordCreatePayload{}
	userId := utilities.UserIdFromAuthHeader(c)

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	word := models.Word{}
	initializers.DB.First(&word, params.WordId)

	now := time.Now().Format(time.RFC3339)
	guessWord := models.GuessWord{}
	guessWord.WordId = sql.NullInt64{Int64: int64(word.Id), Valid: true}
	if userId > 0 {
		guessWord.UserId = sql.NullInt64{Int64: int64(userId), Valid: true}
	}
	guessWord.CreatedAt = now
	guessWord.UpdatedAt = now
	guessWord.Status = enum.Playing

	initializers.DB.Save(&guessWord)

	guessWord.Word = word

	// response
	c.JSON(http.StatusOK, guessWord.Json())
}

// @Summary      Guess Word Guess
// @Description  try to guess the word
// @Tags         Guess Word
// @Accept       json
// @Produce      json
// @Param Id path int true "Guess Word ID"
// @Param	data	body	payloads.GuessWordGuessPayload		true	"Guess Options"
// @Success      200  {object} models.GuessWordJson
// @Router       /api/guess_word/{Id}/guess [post]
func GuessWordGuess(c *gin.Context) {
	params := payloads.GuessWordGuessPayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	guessWord := models.GuessWord{}
	initializers.DB.First(&guessWord, id)

	now := time.Now().Format(time.RFC3339)
	guess := models.GuessWordGuess{}
	guess.GuessWordId = int64(guessWord.Id)
	guess.CreatedAt = now
	guess.UpdatedAt = now
	guess.Guess = params.Guess

	initializers.DB.Save(&guess)

	green := calculateGuessRatings(params.Word, params.Guess, int64(guess.Id))

	var count int64
	initializers.DB.Where("guess_word_id = ?", id).Find(&models.GuessWordGuess{}).Count(&count)

	updateGuessWordStatus(green, len(params.Word), int(count), int64(guessWord.Id))

	found := models.GuessWordGuess{}
	initializers.DB.Preload("Ratings").First(&found, guess.Id)

	c.JSON(http.StatusOK, found.Json())
}

// @Summary      Guess Word Hints
// @Description  list of possible guesses
// @Tags         Guess Word
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.GuessWordHintsPayload		true	"Guess Word Hint Options"
// @Success      200  {array} string
// @Router       /api/guess_word/hint [post]
func GuessWordHints(c *gin.Context) {
	params := payloads.GuessWordHintsPayload{}

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rows, err := initializers.DB.Model(&models.Word{}).Select("Word").Where("Length = ?", params.Length).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var hints []string

	for rows.Next() {
		var Word string
		rows.Scan(&Word)

		matchGreen := matchGreen(Word, params.Green)
		matchGray := matchGray(Word, params.Green, params.Gray)
		matchBrown := matchBrown(Word, params.Brown)
		includeAllBrown := includeAllBrown(Word, params.Brown)

		if matchGreen && !matchGray && !matchBrown && includeAllBrown {
			hints = append(hints, Word)
		}

	}

	c.JSON(http.StatusOK, hints)
}

func GuessWordInProgress(c *gin.Context) {
	userId := utilities.UserIdFromAuthHeader(c)
	guessWordsJson := []models.GuessWordJson{}
	if userId > 0 {
		guessWords := []models.GuessWord{}
		initializers.DB.Where("user_id = ? AND Status = 1", userId).Preload(clause.Associations).Find(&guessWords)
		for i := 0; i < len(guessWords); i++ {
			guessWord := guessWords[i].Json()
			guessWordsJson = append(guessWordsJson, guessWord)
		}
	}
	c.JSON(http.StatusOK, guessWordsJson)
}

func matchGreen(Word string, green []string) bool {
	word := strings.Split(Word, "")
	if len(green) == 0 {
		return true
	}
	for i := 0; i < len(word); i++ {
		if word[i] != green[i] && green[i] != "" {
			return false
		}
	}
	return true
}

func matchBrown(Word string, brown [][]string) bool {
	word := strings.Split(Word, "")
	for i := 0; i < len(word); i++ {
		if len(brown[i]) > 0 && utilities.StringSliceIndexOf(word[i], brown[i]) != -1 {
			return true
		}
	}
	return false
}

func matchGray(Word string, green []string, gray []string) bool {
	word := strings.Split(Word, "")
	for i := 0; i < len(word); i++ {
		if utilities.StringSliceIndexOf(word[i], green) == -1 && utilities.StringSliceIndexOf(word[i], gray) != -1 {
			return true
		}
	}
	return false
}

func includeAllBrown(Word string, brown [][]string) bool {
	word := strings.Split(Word, "")
	var allBrown []string
	for _, row := range brown {
		allBrown = append(allBrown, row...)
	}
	if len(allBrown) == 0 {
		return true
	}
	for i := 0; i < len(allBrown); i++ {
		if utilities.StringSliceIndexOf(allBrown[i], word) == -1 {
			return false
		}
	}
	return true
}

func calculateGuessRatings(Word string, Guess string, guessId int64) int {
	var green []int
	var brown []int
	word := strings.Split(Word, "")
	guess := strings.Split(Guess, "")

	l := len(word)

	for i := 0; i < l; i++ {
		letter := guess[i]
		if letter == word[i] {
			green = append(green, i)
			word[i] = ""
		}
	}
	for i := l - 1; i > -0; i-- {
		if utilities.IntSliceIndexOf(i, green) == -1 {
			letter := guess[i]
			idx := utilities.StringSliceIndexOf(letter, word)
			if idx != -1 {
				brown = append(brown, i)
				word[idx] = ""
			}
		}
	}
	for i := 0; i < l; i++ {
		rating := enum.Gray
		if utilities.IntSliceIndexOf(i, green) != -1 {
			rating = enum.GREEN
		} else if utilities.IntSliceIndexOf(i, brown) != -1 {
			rating = enum.BROWN
		}

		now := time.Now().Format(time.RFC3339)
		guessRating := models.GuessWordGuessRating{}
		guessRating.GuessWordGuessId = guessId
		guessRating.Rating = rating
		guessRating.CreatedAt = now
		guessRating.UpdatedAt = now

		initializers.DB.Save(&guessRating)
	}

	return len(green)
}

func updateGuessWordStatus(green int, wordLength int, guesses int, id int64) {
	status := enum.Playing
	if green == wordLength {
		status = enum.Won
	} else if guesses > int(math.Ceil((float64(wordLength)*3)/2)) {
		status = enum.Lost
	}
	if status != enum.Playing {
		guessWord := models.GuessWord{}
		initializers.DB.First(&guessWord, id)
		initializers.DB.Model(&guessWord).Update("Status", status)
	}
}

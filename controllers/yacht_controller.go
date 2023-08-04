package controllers

import (
	"database/sql"
	"go-games-api/enum"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"go-games-api/utilities"
	"math/rand"
	"net/http"
	"sort"
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

// @Summary      Create Yacht
// @Description  create a yacht
// @Tags         Yacht
// @Accept       json
// @Produce      json
// @Success      200  {object} models.YachtJson
// @Router       /api/yacht [post]
func YachtCreate(c *gin.Context) {
	now := time.Now().Format(time.RFC3339)
	userId := utilities.UserIdFromAuthHeader(c)

	yacht := models.Yacht{}
	if userId > 0 {
		yacht.UserId = sql.NullInt64{Int64: int64(userId), Valid: true}
	}
	yacht.CreatedAt = now
	yacht.UpdatedAt = now
	initializers.DB.Save(&yacht)

	c.JSON(http.StatusCreated, yacht.Json())
}

// @Summary      Yacht roll
// @Description  roll yacht dice
// @Tags         Yacht
// @Accept       json
// @Produce      json
// @Param Id path int true "Yacht ID"
// @Param	data	body	payloads.YachtRollPayload		true	"Yacht Options"
// @Success      200  {object} payloads.YachtRollResponse
// @Router       /api/yacht/{Id}/roll [post]
func YachtRoll(c *gin.Context) {
	params := payloads.YachtRollPayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dice []int
	for i := 0; i < len(params.Keep); i++ {
		face := params.Keep[i]
		dice = append(dice, face)
	}
	for len(dice) < 5 {
		roll := rand.Intn(6) + 1
		dice = append(dice, roll)
	}

	yacht := models.Yacht{}
	initializers.DB.First(&yacht, id)

	now := time.Now().Format(time.RFC3339)
	turn := models.YachtTurn{}
	initializers.DB.Where("yacht_id = ? AND Category IS NULL", id).Order("id DESC").First(&turn)
	if turn.Id > 0 {
		if turn.RollTwo != "" {
			turn.RollThree = utilities.IntSliceJoin(dice, ",")
		} else {
			turn.RollTwo = utilities.IntSliceJoin(dice, ",")
		}
	} else {
		turn.RollOne = utilities.IntSliceJoin(dice, ",")
		turn.YachtId = int64(yacht.Id)
		turn.CreatedAt = now
	}
	turn.UpdatedAt = now

	skip := utilities.YachtCatagorySkip(int64(yacht.Id))
	options := utilities.YachtScoreOptions(dice, skip)

	sort.Slice(options, func(i, j int) bool {
		return options[j].Score < options[i].Score
	})

	initializers.DB.Save(&turn)

	response := payloads.YachtRollResponse{}
	response.Turn = turn.Json()
	response.Options = options

	c.JSON(http.StatusOK, response)
}

// @Summary      Yacht roll
// @Description  roll yacht dice
// @Tags         Yacht
// @Accept       json
// @Produce      json
// @Param Id path int true "Yacht ID"
// @Param	data	body	payloads.YachtScorePayload		true	"Yacht Options"
// @Success      200  {object} models.YachtTurnJson
// @Router       /api/yacht/{Id}/score [post]
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

	now := time.Now().Format(time.RFC3339)
	turn := models.YachtTurn{}
	initializers.DB.First(&turn, params.TurnId)
	var dice []int
	category := string(params.Category)
	if turn.RollThree != "" {
		dice = utilities.StringToIntSlice(turn.RollThree, ",")
	} else if turn.RollTwo != "" {
		dice = utilities.StringToIntSlice(turn.RollTwo, ",")
	} else {
		dice = utilities.StringToIntSlice(turn.RollOne, ",")
	}
	turn.Category = sql.NullInt32{Int32: int32(enum.YachtCategoryArrayIndex(category)), Valid: true}
	turn.Score = utilities.ScoreYachtCategory(category, dice)
	turn.UpdatedAt = now

	initializers.DB.Save(&turn)

	updateYachtTotal(int64(yacht.Id))

	c.JSON(http.StatusOK, turn.Json())
}

func updateYachtTotal(id int64) {
	turns := []models.YachtTurn{}
	initializers.DB.Where("yacht_id = ? and Category IS NOT NULL", id).Select("Score").Find(&turns)

	total := 0
	for i := 0; i < len(turns); i++ {
		total = total + turns[i].Score
	}

	yacht := models.Yacht{}
	initializers.DB.First(&yacht, id)

	yacht.NumTurns = len(turns)
	yacht.Total = total
	yacht.UpdatedAt = time.Now().Format(time.RFC3339)

	initializers.DB.Save(&yacht)
}

func YachtInProgress(c *gin.Context) {
	userId := utilities.UserIdFromAuthHeader(c)
	yachtsJson := []models.YachtJson{}
	if userId > 0 {
		yachts := []models.Yacht{}
		initializers.DB.Where("user_id = ? AND NumTurns < 12", userId).Find(&yachts)
		for i := 0; i < len(yachts); i++ {
			yacht := yachts[i].Json()
			yachtsJson = append(yachtsJson, yacht)
		}
	}
	c.JSON(http.StatusOK, yachtsJson)
}

package controllers

import (
	"database/sql"
	"fmt"
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
	"gorm.io/gorm/clause"
)

// @Summary      Ten Grand
// @Description  paginated list of ten grand
// @Tags         Ten Grand
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.TenGrandPaginatedJson
// @Router       /api/ten_grand [get]
func TenGrandIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.TenGrandPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.TenGrand{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Ten Grand by Id
// @Description  get a ten grand
// @Tags         Ten Grand
// @Accept       json
// @Produce      json
// @Param Id path int true "Ten Grand ID"
// @Success      200  {object} models.TenGrandJson
// @Router       /api/ten_grand/{Id} [get]
func TenGrandById(c *gin.Context) {
	// get id
	id := c.Param("id")

	tenGrand := models.TenGrand{}
	initializers.DB.Preload("Turns.Scores").Preload(clause.Associations).First(&tenGrand, id)

	// response
	c.JSON(http.StatusOK, tenGrand.Json())
}

// @Summary      Create Ten Grand
// @Description  create a ten grand
// @Tags         Ten Grand
// @Accept       json
// @Produce      json
// @Success      201  {object} models.TenGrandJson
// @Router       /api/ten_grand [post]
func TenGrandCreate(c *gin.Context) {
	now := time.Now().Format(time.RFC3339)
	userId := utilities.UserIdFromAuthHeader(c)

	tenGrand := models.TenGrand{}
	if userId > 0 {
		tenGrand.UserId = sql.NullInt64{Int64: int64(userId), Valid: true}
	}
	tenGrand.Status = enum.Playing
	tenGrand.CreatedAt = now
	tenGrand.UpdatedAt = now
	initializers.DB.Save(&tenGrand)

	c.JSON(http.StatusCreated, tenGrand.Json())
}

// @Summary      Ten Grand Die Roll
// @Description  get an array of dice
// @Tags         Ten Grand
// @Accept       json
// @Produce      json
// @Param Id path int true "Ten Grand ID"
// @Param	data	body	payloads.TenGrandRollPayload		true	"Ten Grand Options"
// @Success      200  {object} []int
// @Router       /api/ten_grand/{Id}/roll [post]
func TenGrandRoll(c *gin.Context) {
	params := payloads.TenGrandRollPayload{}

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pretty, err := utilities.PrettyStruct(params)
	if err == nil {
		fmt.Println(pretty)
	}

	var dice []int
	for i := 0; i < params.Quantity; i++ {
		roll := rand.Intn(6) + 1
		dice = append(dice, roll)
	}

	c.JSON(http.StatusOK, dice)
}

// @Summary      Ten Grand Die Scores
// @Description  get an array of score options
// @Tags         Ten Grand
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.TenGrandOptionsPayload		true	"Ten Grand Options"
// @Success      200  {object} []payloads.TenGrandOption
// @Router       /api/ten_grand/options [post]
func TenGrandOptions(c *gin.Context) {
	params := payloads.TenGrandOptionsPayload{}

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pretty, err := utilities.PrettyStruct(params)
	if err == nil {
		fmt.Println(pretty)
	}
	dice, options := utilities.TenGrandDiceScoreOptions(params.Dice)
	type OptionsResponse struct {
		Dice    []int
		Options []payloads.TenGrandOption
	}

	c.JSON(http.StatusOK, OptionsResponse{Dice: dice, Options: options})
}

// @Summary      Ten Grand Turn Score
// @Description  get an array of dice
// @Tags         Ten Grand
// @Accept       json
// @Produce      json
// @Param Id path int true "Ten Grand ID"
// @Param	data	body	payloads.TenGrandScoreOptionsPayload		true	"Ten Grand Options"
// @Success      200  {object} models.TenGrandTurnJson
// @Router       /api/ten_grand/{Id}/score [post]
func TenGrandScore(c *gin.Context) {
	params := payloads.TenGrandScoreOptionsPayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dice := params.Dice
	options := params.Options
	sort.Slice(options, func(i, j int) bool {
		return enum.TenGrandDiceRequired[string(options[j].Category)] < enum.TenGrandDiceRequired[string(options[i].Category)]
	})

	tenGrand := models.TenGrand{}
	initializers.DB.First(&tenGrand, id)

	now := time.Now().Format(time.RFC3339)
	turn := models.TenGrandTurn{}
	if params.TurnId > 0 {
		initializers.DB.First(&turn, params.TurnId)
	} else {
		turn.CreatedAt = now
		turn.UpdatedAt = now
		turn.TenGrandId = int64(tenGrand.Id)
		initializers.DB.Save(&turn)
	}
	for i := 0; i < len(options); i++ {
		opt := options[i]
		score, used := utilities.CategoryScoreAndDice(opt.Category, dice)
		dice = utilities.RemoveUsedDice(dice, used)

		tenGrandScore := models.TenGrandScore{}
		tenGrandScore.TenGrandTurnId = int64(turn.Id)
		tenGrandScore.Score = score
		tenGrandScore.Dice = utilities.IntSliceJoin(used, ",")
		tenGrandScore.Category = enum.TenGrandCategory(enum.TenGrandCategoryArrayIndex(string(opt.Category)))
		tenGrandScore.CreatedAt = now
		tenGrandScore.UpdatedAt = now
		initializers.DB.Save(&tenGrandScore)
	}

	updateTenGrandTurnScore(int64(turn.Id))

	found := models.TenGrandTurn{}
	initializers.DB.Preload("Scores").First(&found, turn.Id)

	updateTenGrandScore(int64(tenGrand.Id))

	c.JSON(http.StatusOK, found.Json())
}

func updateTenGrandTurnScore(id int64) {
	score := 0
	crapOut := false
	scores := []models.TenGrandScore{}
	initializers.DB.Where("ten_grand_turn_id = ?", id).Find(&scores)
	for s := 0; s < len(scores); s++ {
		if scores[s].Score > 0 {
			score = score + scores[s].Score
		}
		if scores[s].Category == enum.CrapOut {
			crapOut = true
		}
	}
	if crapOut {
		score = 0
	}
	turn := models.TenGrandTurn{}
	initializers.DB.First(&turn, id)
	initializers.DB.Model(&turn).Update("Score", score)
}

func updateTenGrandScore(id int64) {
	turns := []models.TenGrandTurn{}
	score := 0
	status := enum.Playing
	initializers.DB.Where("ten_grand_id = ?", id).Find(&turns)
	for t := 0; t < len(turns); t++ {
		if turns[t].Score > 0 {
			score = score + turns[t].Score
		}
	}
	if score >= 10000 {
		status = enum.Won
	}
	tenGrand := models.TenGrand{}
	initializers.DB.First(&tenGrand, id)
	initializers.DB.Model(&tenGrand).Update("Score", score).Update("Status", status)
}

func TenGrandInProgress(c *gin.Context) {
	userId := utilities.UserIdFromAuthHeader(c)
	tenGrandsJson := []models.TenGrandJson{}
	if userId > 0 {
		tenGrands := []models.TenGrand{}
		initializers.DB.Where("user_id = ? AND Status = 1", userId).Preload(clause.Associations).Find(&tenGrands)
		for i := 0; i < len(tenGrands); i++ {
			tenGrand := tenGrands[i].Json()
			tenGrandsJson = append(tenGrandsJson, tenGrand)
		}
	}
	c.JSON(http.StatusOK, tenGrandsJson)
}

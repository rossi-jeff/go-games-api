package controllers

import (
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
	tenGrand := models.TenGrand{}
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

	c.JSON(http.StatusOK, utilities.TenGrandDiceScoreOptions(params.Dice))
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

	now := time.Now().Format(time.RFC3339)
	turn := models.TenGrandTurn{}
	if params.TurnId > 0 {
		initializers.DB.First(&turn, params.TurnId)
	} else {
		tenGrand := models.TenGrand{}
		initializers.DB.First(&tenGrand, id)

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

	found := models.TenGrandTurn{}
	initializers.DB.Preload("Scores").First(&found, turn.Id)

	c.JSON(http.StatusOK, found.Json())
}

func TenGrandInProgress(c *gin.Context) {
	userId := utilities.UserIdFromAuthHeader(c)
	tenGrandsJson := []models.TenGrandJson{}
	if userId > 0 {
		tenGrands := []models.TenGrand{}
		initializers.DB.Where("user_id = ? AND Status = 1", userId).Find(&tenGrands)
		for i := 0; i < len(tenGrands); i++ {
			tenGrand := tenGrands[i].Json()
			tenGrandsJson = append(tenGrandsJson, tenGrand)
		}
	}
	c.JSON(http.StatusOK, tenGrandsJson)
}

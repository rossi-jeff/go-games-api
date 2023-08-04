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

// @Summary      Klondike
// @Description  paginated list of klondike
// @Tags         Klondike
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.KlondikePaginatedJson
// @Router       /api/klondike [get]
func KlondikeIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.KlondikePaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.Klondike{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Moves ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Klondike by Id
// @Description  get a klondike
// @Tags         Klondike
// @Accept       json
// @Produce      json
// @Param Id path int true "Klondike ID"
// @Success      200  {object} models.KlondikeJson
// @Router       /api/klondike/{Id} [get]
func KlondikeById(c *gin.Context) {
	// get id
	id := c.Param("id")

	klondike := models.Klondike{}
	initializers.DB.Preload("User").First(&klondike, id)

	// response
	c.JSON(http.StatusOK, klondike.Json())
}

// @Summary      Create Klondike
// @Description  paginated list of klondike
// @Tags         Klondike
// @Accept       json
// @Produce      json
// @Success      201  {object} models.KlondikeJson
// @Router       /api/klondike [post]
func KlondikeCreate(c *gin.Context) {
	now := time.Now().Format(time.RFC3339)
	userId := utilities.UserIdFromAuthHeader(c)

	klondike := models.Klondike{}
	if userId > 0 {
		klondike.UserId = sql.NullInt64{Int64: int64(userId), Valid: true}
	}
	klondike.Status = enum.Playing
	klondike.CreatedAt = now
	klondike.UpdatedAt = now
	initializers.DB.Save(&klondike)

	c.JSON(http.StatusCreated, klondike.Json())
}

// @Summary      Update Klondike
// @Description  get a klondike
// @Tags         Klondike
// @Accept       json
// @Produce      json
// @Param Id path int true "Klondike ID"
// @Param	data	body	payloads.KlondikeUpdatePayload		true	"Klondike Updates"
// @Success      200  {object} models.KlondikeJson
// @Router       /api/klondike/{Id} [patch]
func KlondikeUpdate(c *gin.Context) {
	params := payloads.KlondikeUpdatePayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	klondike := models.Klondike{}
	initializers.DB.Select("id,created_at,user_id").First(&klondike, id)
	klondike.Elapsed = params.Elapsed
	klondike.Moves = params.Moves
	klondike.Status = enum.GameStatus(enum.GameStatusArrayIndex(string(params.Status)))
	klondike.UpdatedAt = time.Now().Format(time.RFC3339)

	initializers.DB.Save(&klondike)

	c.JSON(http.StatusOK, klondike.Json())
}

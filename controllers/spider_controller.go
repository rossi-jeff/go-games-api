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

// @Summary      Spider
// @Description  paginated list of spider
// @Tags         Spider
// @Accept       json
// @Produce      json
// @Param	Limit	query	int	false	"Limit"
// @Param	Offset	query	int	false	"Offset"
// @Success      200  {object} models.SpiderPaginatedJson
// @Router       /api/spider [get]
func SpiderIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.SpiderPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Where("Status <> 1").Find(&models.Spider{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Where("Status <> 1").Order("Status DESC, Moves ASC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response.Json())
}

// @Summary      Spider by Id
// @Description  get a spider
// @Tags         Spider
// @Accept       json
// @Produce      json
// @Param Id path int true "Spider ID"
// @Success      200  {object} models.SpiderJson
// @Router       /api/spider/{Id} [get]
func SpiderById(c *gin.Context) {
	// get id
	id := c.Param("id")

	spider := models.Spider{}
	initializers.DB.Preload("User").First(&spider, id)

	// response
	c.JSON(http.StatusOK, spider.Json())
}

// @Summary      Create Spider
// @Description  create a spider
// @Tags         Spider
// @Accept       json
// @Produce      json
// @Param	data	body	payloads.SpiderCreatePayload		true	"Spider Options"
// @Success      200  {object} models.SpiderJson
// @Router       /api/spider [post]
func SpiderCreate(c *gin.Context) {
	params := payloads.SpiderCreatePayload{}
	userId := utilities.UserIdFromAuthHeader(c)

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	spider := models.Spider{}
	if userId > 0 {
		spider.UserId = sql.NullInt64{Int64: int64(userId), Valid: true}
	}
	spider.Suits = enum.Suit(params.Suits)
	spider.Status = enum.Playing
	spider.CreatedAt = now
	spider.UpdatedAt = now
	initializers.DB.Save(&spider)

	found := models.Spider{}
	initializers.DB.Preload("User").First(&found, spider.Id)

	// response
	c.JSON(http.StatusOK, found.Json())
}

// @Summary      Update Spider
// @Description  update a spider
// @Tags         Spider
// @Accept       json
// @Produce      json
// @Param Id path int true "Spider ID"
// @Param	data	body	payloads.SpiderUpdatePayload		true	"Spider Options"
// @Success      200  {object} models.SpiderJson
// @Router       /api/spider/{Id} [patch]
func SpiderUpdate(c *gin.Context) {
	params := payloads.SpiderUpdatePayload{}
	id := c.Param("id")

	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().Format(time.RFC3339)
	spider := models.Spider{}
	initializers.DB.Select("id,created_at,Suits").First(&spider, id)
	spider.UpdatedAt = now
	spider.Moves = params.Moves
	spider.Elapsed = params.Elapsed
	spider.Status = enum.GameStatus(enum.GameStatusArrayIndex(string(params.Status)))

	initializers.DB.Save(&spider)

	c.JSON(http.StatusOK, spider.Json())
}

package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

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

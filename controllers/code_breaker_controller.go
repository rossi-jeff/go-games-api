package controllers

import (
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CodeBreakerIndex(c *gin.Context) {
	params := utilities.ParseIndexParams(c)
	response := models.CodeBreakerPaginated{}
	response.Limit = params.Limit
	response.Offset = params.Offset
	var count int64
	initializers.DB.Find(&models.CodeBreaker{}).Count(&count)
	response.Count = int(count)
	initializers.DB.Order("Score DESC").Offset(params.Offset).Limit(params.Limit).Preload("User").Find(&response.Items)

	// response
	c.JSON(http.StatusOK, response)
}

func CodeBreakerById(c *gin.Context) {
	// get id
	id := c.Param("id")

	// find code breaker
	codeBreaker := models.CodeBreaker{}
	initializers.DB.Preload("Guesses.Colors").Preload("Guesses.Keys").Preload(clause.Associations).First(&codeBreaker, id)

	// response
	c.JSON(http.StatusOK, codeBreaker)
}

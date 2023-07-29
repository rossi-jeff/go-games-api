package controllers

import (
	"fmt"
	"go-games-api/initializers"
	"go-games-api/models"
	"go-games-api/payloads"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WordById(c *gin.Context) {
	// get id
	id := c.Param("id")

	// find word
	word := models.Word{}
	initializers.DB.First(&word, id)

	// response
	c.JSON(http.StatusOK, word)
}

func WordRandom(c *gin.Context) {
	// get parameters
	params := payloads.RandomWordPayload{}
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// select random word
	var where string
	word := models.Word{}
	if params.Length > 0 {
		where = fmt.Sprintf("Length = %d", params.Length)
	} else if params.Min > 0 && params.Max > 0 {
		where = fmt.Sprintf("Length BETWEEN %d AND %d", params.Min, params.Max)
	} else if params.Min > 0 {
		where = fmt.Sprintf("Length >= %d", params.Min)
	} else if params.Max > 0 {
		where = fmt.Sprintf("Length <= %d", params.Max)
	}
	initializers.DB.Where(where).Order("RAND()").Limit(1).Find(&word)

	// response
	c.JSON(http.StatusOK, word)
}

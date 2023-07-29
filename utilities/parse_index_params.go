package utilities

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var defaultOffset = 0
var defaultLimit = 10

type IndexParams struct {
	Limit  int
	Offset int
}

func ParseIndexParams(c *gin.Context) IndexParams {
	params := IndexParams{
		Limit:  defaultLimit,
		Offset: defaultOffset,
	}
	limit := c.Query("Limit")
	offset := c.Query("Offset")

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err == nil {
			params.Limit = l
		}
	}
	if offset != "" {
		o, err := strconv.Atoi(offset)
		if err == nil {
			params.Offset = o
		}
	}

	return params
}

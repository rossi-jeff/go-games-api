package models

import "go-games-api/enum"

type YachtTurn struct {
	BaseModel
	RollOne   string
	RollTwo   string
	RollThree string
	Category  enum.YachtCategory
	Score     int
	YachtId   int64 `json:"yacht_id"`
}

type YachtTurnJson struct {
	BaseModel
	RollOne   string
	RollTwo   string
	RollThree string
	Score     int
	YachtId   int64 `json:"yacht_id"`
	Category  enum.YachtCategoryString
}

func (y YachtTurn) Json() YachtTurnJson {
	return YachtTurnJson{
		BaseModel: y.BaseModel,
		RollOne:   y.RollOne,
		RollTwo:   y.RollTwo,
		RollThree: y.RollThree,
		Score:     y.Score,
		YachtId:   y.YachtId,
		Category:  enum.YachtCategoryString(y.Category.String()),
	}
}

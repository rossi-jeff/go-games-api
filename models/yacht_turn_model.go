package models

import "go-games-api/enum"

type YachtTurn struct {
	BaseModel
	RollOne, RollTwo, RollThree string
	Category                    enum.YachtCategory
	Score                       int
	YachtId                     int64 `json:"yacht_id"`
}

type YachtTurnJson struct {
	BaseModel
	RollOne, RollTwo, RollThree string
	Score                       int
	YachtId                     int64 `json:"yacht_id"`
	Category                    string
}

func (y YachtTurn) Json() YachtTurnJson {
	return YachtTurnJson{
		BaseModel: y.BaseModel,
		RollOne:   y.RollOne,
		RollTwo:   y.RollTwo,
		RollThree: y.RollThree,
		Score:     y.Score,
		YachtId:   y.YachtId,
		Category:  y.Category.String(),
	}
}

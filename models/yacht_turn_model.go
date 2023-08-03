package models

import "go-games-api/enum"

type YachtTurn struct {
	BaseModel
	RollOne   string             `gorm:"column:RollOne"`
	RollTwo   string             `gorm:"column:RollTwo"`
	RollThree string             `gorm:"column:RollThree"`
	Category  enum.YachtCategory `json:"Category,omitempty" gorm:"default:null"`
	Score     int
	YachtId   int64 `json:"yacht_id"`
}

type YachtTurnJson struct {
	BaseModel
	RollOne   string `gorm:"column:RollOne"`
	RollTwo   string `gorm:"column:RollTwo"`
	RollThree string `gorm:"column:RollThree"`
	Score     int
	YachtId   int64                    `json:"yacht_id"`
	Category  enum.YachtCategoryString `json:"Category,omitempty" gorm:"default:null"`
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

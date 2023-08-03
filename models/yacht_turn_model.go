package models

import "go-games-api/enum"

type YachtTurn struct {
	BaseModel
	RollOne   string    `gorm:"column:RollOne"`
	RollTwo   string    `gorm:"column:RollTwo"`
	RollThree string    `gorm:"column:RollThree"`
	Category  NullInt32 `json:"Category,omitempty" gorm:"default:null"`
	Score     int
	YachtId   int64 `json:"yacht_id"`
}

type YachtTurnJson struct {
	BaseModel
	RollOne   string `gorm:"column:RollOne"`
	RollTwo   string `gorm:"column:RollTwo"`
	RollThree string `gorm:"column:RollThree"`
	Score     int
	YachtId   int64  `json:"yacht_id"`
	Category  string `json:"Category,omitempty" gorm:"default:null"`
}

func (y YachtTurn) Json() YachtTurnJson {
	return YachtTurnJson{
		BaseModel: y.BaseModel,
		RollOne:   y.RollOne,
		RollTwo:   y.RollTwo,
		RollThree: y.RollThree,
		Score:     y.Score,
		YachtId:   y.YachtId,
		Category:  YachtCategoryJson(y.Category),
	}
}

func YachtCategoryJson(category NullInt32) string {
	if category.Valid {
		return enum.YachtCategoryArray[category.Int32]
	} else {
		return ""
	}
}

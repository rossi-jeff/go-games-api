package models

import "go-games-api/enum"

type TenGrandScore struct {
	BaseModel
	Dice           string
	Category       enum.TenGrandCategory
	Score          int
	TenGrandTurnId int64 `json:"ten_grand_turn_id"`
}

type TenGrandScoreJson struct {
	BaseModel
	Dice           string
	Score          int
	TenGrandTurnId int64 `json:"ten_grand_turn_id"`
	Category       enum.TenGrandCategoryString
}

func (t TenGrandScore) Json() TenGrandScoreJson {
	return TenGrandScoreJson{
		BaseModel:      t.BaseModel,
		Dice:           t.Dice,
		Score:          t.Score,
		TenGrandTurnId: t.TenGrandTurnId,
		Category:       enum.TenGrandCategoryString(t.Category.String()),
	}
}

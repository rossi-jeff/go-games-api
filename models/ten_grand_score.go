package models

import "go-games-api/enum"

type TenGrandScore struct {
	BaseModel
	Dice           string
	Category       enum.TenGrandCategory
	Score          int
	TenGrandTurnId int64 `json:"ten_grand_turn_id"`
}

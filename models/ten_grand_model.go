package models

import "go-games-api/enum"

type TenGrand struct {
	BaseModel
	Status enum.GameStatus
	Score  int
	UserId int64          `json:"user_id"`
	User   User           `json:"user,omitempty"`
	Turns  []TenGrandTurn `json:"turns,omitempty"`
}

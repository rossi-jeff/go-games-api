package models

import "go-games-api/enum"

type PokerSquare struct {
	BaseModel
	Status enum.GameStatus
	Score  int
	UserId int64 `json:"user_id"`
	User   User  `json:"user,omitempty"`
}

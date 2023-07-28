package models

import "go-games-api/enum"

type SeaBattle struct {
	BaseModel
	Status      enum.GameStatus
	Score, Axis int
	UserId      int64           `json:"user_id"`
	User        User            `json:"user,omitempty"`
	Ships       []SeaBattleShip `json:"ships"`
	Turns       []SeaBattleTurn `json:"turns"`
}

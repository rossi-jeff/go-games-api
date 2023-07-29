package models

import "go-games-api/enum"

type Klondike struct {
	BaseModel
	Status         enum.GameStatus
	Moves, Elapsed int
	UserId         int64 `json:"user_id"`
	User           User  `json:"user,omitempty"`
}

type KlondikePaginated struct {
	Items []Klondike
	Paginated
}

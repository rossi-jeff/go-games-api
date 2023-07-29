package models

import "go-games-api/enum"

type FreeCell struct {
	BaseModel
	Status         enum.GameStatus
	Moves, Elapsed int
	UserId         int64 `json:"user_id"`
	User           User  `json:"user,omitempty"`
}

type FreeCellPaginated struct {
	Items []FreeCell
	Paginated
}

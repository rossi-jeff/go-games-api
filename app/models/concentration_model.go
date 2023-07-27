package models

import "go-games-api/app/enum"

type Concentration struct {
	BaseModel
	Status                  enum.GameStatus
	Moves, Matched, Elapsed int
	UserId                  int64 `json:"user_id"`
	User                    User  `json:"user,omitempty"`
}

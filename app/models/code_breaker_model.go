package models

import "go-games-api/app/enum"

type CodeBreaker struct {
	BaseModel
	Status                 enum.GameStatus
	Columns, Colors, Score int
	Available              string
	UserId                 int64 `json:"user_id"`
	User                   User  `json:"user,omitempty"`
}

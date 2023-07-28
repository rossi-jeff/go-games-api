package models

import "go-games-api/app/enum"

type YachtTurn struct {
	BaseModel
	RollOne, RollTwo, RollThree string
	Category                    enum.YachtCategory
	Score                       int
	YachtId                     int64 `json:"yacht_id"`
}

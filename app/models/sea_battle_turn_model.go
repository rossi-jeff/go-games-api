package models

import "go-games-api/app/enum"

type SeaBattleTurn struct {
	BaseModel
	ShipType    enum.ShipType
	Navy        enum.Navy
	Target      enum.Target
	Horizontal  string
	Vertical    int
	SeaBattleId int64 `json:"sea_battle_id"`
}

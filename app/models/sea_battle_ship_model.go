package models

import "go-games-api/app/enum"

type SeaBattleShip struct {
	BaseModel
	Type        enum.ShipType
	Navy        enum.Navy
	Size        int
	Sunk        bool
	SeaBattleId int64 `json:"sea_battle_id"`
}

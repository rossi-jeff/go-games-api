package models

type SeaBattleShipHit struct {
	BaseModel
	Horizontal      string
	Vertical        int
	SeaBattleShipId int64 `json:"sea_battle_ship_id"`
}

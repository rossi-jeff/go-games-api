package models

type SeaBattleShipGridPoint struct {
	BaseModel
	Horizontal      string
	Vertical        int
	SeaBattleShipId int64 `json:"sea_battle_ship_id"`
}

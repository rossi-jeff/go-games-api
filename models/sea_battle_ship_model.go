package models

import "go-games-api/enum"

type SeaBattleShip struct {
	BaseModel
	Type        enum.ShipType
	Navy        enum.Navy
	Size        int
	Sunk        bool
	SeaBattleId int64                    `json:"sea_battle_id"`
	Points      []SeaBattleShipGridPoint `json:"points"`
	Hits        []SeaBattleShipHit       `json:"hits"`
}

type SeaBattleShipJson struct {
	BaseModel
	Size        int
	Sunk        bool
	SeaBattleId int64                    `json:"sea_battle_id"`
	Points      []SeaBattleShipGridPoint `json:"points"`
	Hits        []SeaBattleShipHit       `json:"hits"`
	Type        enum.ShipTypeString
	Navy        enum.NavyString
}

func (s SeaBattleShip) Json() SeaBattleShipJson {
	return SeaBattleShipJson{
		BaseModel:   s.BaseModel,
		Size:        s.Size,
		Sunk:        s.Sunk,
		SeaBattleId: s.SeaBattleId,
		Points:      s.Points,
		Hits:        s.Hits,
		Type:        enum.ShipTypeString(s.Type.String()),
		Navy:        enum.NavyString(s.Navy.String()),
	}
}

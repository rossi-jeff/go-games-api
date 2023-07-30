package models

import "go-games-api/enum"

type SeaBattleTurn struct {
	BaseModel
	ShipType    enum.ShipType
	Navy        enum.Navy
	Target      enum.Target
	Horizontal  string
	Vertical    int
	SeaBattleId int64 `json:"sea_battle_id"`
}

type SeaBattleTurnJson struct {
	BaseModel
	Horizontal  string
	Vertical    int
	SeaBattleId int64 `json:"sea_battle_id"`
	ShipType    string
	Navy        string
	Target      string
}

func (s SeaBattleTurn) Json() SeaBattleTurnJson {
	return SeaBattleTurnJson{
		BaseModel:   s.BaseModel,
		Horizontal:  s.Horizontal,
		Vertical:    s.Vertical,
		SeaBattleId: s.SeaBattleId,
		ShipType:    s.ShipType.String(),
		Navy:        s.Navy.String(),
		Target:      s.Target.String(),
	}
}

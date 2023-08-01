package models

import "go-games-api/enum"

type SeaBattleTurn struct {
	BaseModel
	ShipType    enum.ShipType `gorm:"column:ShipType"`
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
	SeaBattleId int64               `json:"sea_battle_id"`
	ShipType    enum.ShipTypeString `gorm:"column:ShipType"`
	Navy        enum.NavyString
	Target      enum.TargetString
}

func (s SeaBattleTurn) Json() SeaBattleTurnJson {
	return SeaBattleTurnJson{
		BaseModel:   s.BaseModel,
		Horizontal:  s.Horizontal,
		Vertical:    s.Vertical,
		SeaBattleId: s.SeaBattleId,
		ShipType:    enum.ShipTypeString(s.ShipType.String()),
		Navy:        enum.NavyString(s.Navy.String()),
		Target:      enum.TargetString(s.Target.String()),
	}
}

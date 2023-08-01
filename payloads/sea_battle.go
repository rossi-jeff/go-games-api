package payloads

import "go-games-api/enum"

type SeaBattleCreatePayload struct {
	Axis int
}

type SeaBattleShipPayload struct {
	Navy     enum.NavyString
	ShipType enum.ShipTypeString
	Size     int
	Points   []SeaBattlePoint `json:"Points,omitempty"`
}

type SeaBattleFirePayload struct {
	Navy       enum.NavyString
	Horizontal string `json:"Horizontal,omitempty"`
	Vertical   int    `json:"Vertical,omitempty"`
}

type SeaBattlePoint struct {
	Horizontal string
	Vertical   int
}

var HorizontalAxisMax = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var VerticalAxisMax = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 224, 25, 26}
var Directions = []string{"N", "S", "E", "W"}

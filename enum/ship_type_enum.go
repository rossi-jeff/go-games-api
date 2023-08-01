package enum

// ShipType enum info
// @Description database contains integer but values are "BattleShip", "Carrier", "Cruiser", "PatrolBoat", "SubMarine"
type ShipType int

const (
	BattleShip ShipType = iota
	Carrier
	Cruiser
	PatrolBoat
	SubMarine
)

type ShipTypeString string

const (
	ST0 ShipTypeString = "BattleShip"
	ST1 ShipTypeString = "Carrier"
	ST2 ShipTypeString = "Cruiser"
	ST3 ShipTypeString = "PatrolBoat"
	ST4 ShipTypeString = "SubMarine"
)

var ShipTypeArray = [...]string{"BattleShip", "Carrier", "Cruiser", "PatrolBoat", "SubMarine"}

func (s ShipType) String() string {
	if s < 0 || s >= 5 {
		return ""
	}
	return ShipTypeArray[s]
}

func (s ShipType) EnumIndex() int {
	return int(s)
}

func ShipTypeArrayIndex(t string) int {
	for i := 0; i < len(ShipTypeArray); i++ {
		if ShipTypeArray[i] == t {
			return i
		}
	}
	return -1
}

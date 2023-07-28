package enum

type ShipType int

const (
	BattleShip ShipType = iota
	Carrier
	Cruiser
	PatrolBoat
	SubMarine
)

var ShipTypeArray = [...]string{"BattleShip", "Carrier", "Cruiser", "PatrolBoat", "SubMarine"}

func (s ShipType) String() string {
	return ShipTypeArray[s]
}

func (s ShipType) EnumIndex() int {
	return int(s)
}

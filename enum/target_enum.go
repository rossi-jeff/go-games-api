package enum

// Target enum info
// @Description database contains integer but values are "Miss", "Hit", "Sunk"
type Target int

const (
	Miss Target = iota
	Hit
	Sunk
)

type TargetString string

const (
	T0 TargetString = "Miss"
	T1 TargetString = "Hit"
	T2 TargetString = "Sunk"
)

var TargetArray = [3]string{"Miss", "Hit", "Sunk"}

func (t Target) String() string {
	if t < 0 || t >= 3 {
		return ""
	}
	return TargetArray[t]
}

func (t Target) EnumIndex() int {
	return int(t)
}

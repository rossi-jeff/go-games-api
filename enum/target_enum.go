package enum

type Target int

const (
	Miss Target = iota
	Hit
	Sunk
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

package enum

// Navy enum info
// @Description database contains integer but values are "Player", "Opponent"
type Navy int

const (
	Player Navy = iota
	Opponent
)

var NavyArray = [2]string{"Player", "Opponent"}

func (n Navy) String() string {
	if n < 0 || n >= 2 {
		return ""
	}
	return NavyArray[n]
}

func (n Navy) EnumIndex() int {
	return int(n)
}

package enum

type Navy int

const (
	Player Navy = iota
	Opponent
)

var NavyArray = [2]string{"Player", "Opponent"}

func (n Navy) String() string {
	return NavyArray[n]
}

func (n Navy) EnumIndex() int {
	return int(n)
}

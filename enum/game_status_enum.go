package enum

// GameStatus enum info
// @Description  database contains integer but values are "Lost", "Playing", "Won"
type GameStatus int

const (
	Lost GameStatus = iota
	Playing
	Won
)

var GameStatusArray = [3]string{"Lost", "Playing", "Won"}

func (g GameStatus) String() string {
	if g < 0 || g >= 3 {
		return ""
	}
	return GameStatusArray[g]
}

func (g GameStatus) EnumIndex() int {
	return int(g)
}

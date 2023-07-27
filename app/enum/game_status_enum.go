package enum

type GameStatus int

const (
	Lost GameStatus = iota
	Playing
	Won
)

var GameStatusArray = [3]string{"Lost", "Playing", "Won"}

func (g GameStatus) String() string {
	return GameStatusArray[g]
}

func (g GameStatus) EnumIndex() int {
	return int(g)
}

package enum

type Suit int

const (
	ONE  Suit = 1
	TWO  Suit = 2
	FOUR Suit = 4
)

func (s Suit) EnumIndex() int {
	return int(s)
}

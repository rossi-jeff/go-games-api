package enum

// YachtCategory enum info
// @Description database contains integer but values are "BigStraight", "Choice", "Fives", "FourOfKind", "Fours", "FullHouse",
// @description "LittleStraight", "Ones", "Sixes", "Threes", "Twos", "Yacht"
type YachtCategory int

const (
	BigStraight YachtCategory = iota
	Choice
	FIVES
	FourOfKind
	Fours
	FULLHOUSE
	LittleStraight
	ONES
	Sixes
	Threes
	Twos
	Yacht
)

var YachtCategoryArray = []string{"BigStraight", "Choice", "Fives", "FourOfKind", "Fours", "FullHouse", "LittleStraight", "Ones", "Sixes", "Threes", "Twos", "Yacht"}

func (y YachtCategory) String() string {
	return YachtCategoryArray[y]
}

func (y YachtCategory) EnumIndex() int {
	return int(y)
}

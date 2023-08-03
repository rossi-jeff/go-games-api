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

type YachtCategoryString string

const (
	YC0 YachtCategoryString = "BigStraight"
	YC1 YachtCategoryString = "Choice"
	YC2 YachtCategoryString = "Fives"
	YC3 YachtCategoryString = "FourOfKind"
	YC4 YachtCategoryString = "Fours"
	YC5 YachtCategoryString = "FullHouse"
	YC6 YachtCategoryString = "LittleStraight"
	YC7 YachtCategoryString = "Ones"
	YC8 YachtCategoryString = "Sixes"
	YC9 YachtCategoryString = "Threes"
	YCA YachtCategoryString = "Twos"
	YCB YachtCategoryString = "Yacht"
)

var YachtCategoryArray = []string{"BigStraight", "Choice", "Fives", "FourOfKind", "Fours", "FullHouse", "LittleStraight", "Ones", "Sixes", "Threes", "Twos", "Yacht"}

func (y YachtCategory) String() string {
	return YachtCategoryArray[y]
}

func (y YachtCategory) EnumIndex() int {
	return int(y)
}

func YachtCategoryArrayIndex(cat string) int {
	for i := 0; i < len(YachtCategoryArray); i++ {
		if YachtCategoryArray[i] == cat {
			return i
		}
	}
	return -1
}

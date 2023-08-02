package enum

// TenGrandCategory enum info
// @Description  database contains integer but values are "CrapOut", "Ones", "Fives", "ThreePairs", "Straight", "FullHouse",
// @Description "DoubleThreeKind", "ThreeKind", "FourKind", "FiveKind", "SixKind"
type TenGrandCategory int

const (
	CrapOut TenGrandCategory = iota
	Ones
	Fives
	ThreePairs
	Straight
	FullHouse
	DoubleThreeKind
	ThreeKind
	FourKind
	FiveKind
	SixKind
)

type TenGrandCategoryString string

const (
	TG0 TenGrandCategoryString = "CrapOut"
	TG1 TenGrandCategoryString = "Ones"
	TG2 TenGrandCategoryString = "Fives"
	TG3 TenGrandCategoryString = "ThreePairs"
	TG4 TenGrandCategoryString = "Straight"
	TG5 TenGrandCategoryString = "FullHouse"
	TG6 TenGrandCategoryString = "DoubleThreeKind"
	TG7 TenGrandCategoryString = "ThreeKind"
	TG8 TenGrandCategoryString = "FourKind"
	TG9 TenGrandCategoryString = "FiveKind"
	TGA TenGrandCategoryString = "SixKind"
)

var TenGrandCategoryArray = [...]string{"CrapOut", "Ones", "Fives", "ThreePairs", "Straight", "FullHouse", "DoubleThreeKind", "ThreeKind", "FourKind", "FiveKind", "SixKind"}

func (t TenGrandCategory) String() string {
	return TenGrandCategoryArray[t]
}

func (t TenGrandCategory) EnumIndex() int {
	return int(t)
}

var TenGrandDiceRequired = map[string]int{
	"CrapOut":         1,
	"Ones":            1,
	"Fives":           1,
	"ThreePairs":      6,
	"Straight":        6,
	"FullHouse":       5,
	"DoubleThreeKind": 6,
	"ThreeKind":       3,
	"FourKind":        4,
	"FiveKind":        5,
	"SixKind":         6,
}

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

var TenGrandCategoryArray = [...]string{"CrapOut", "Ones", "Fives", "ThreePairs", "Straight", "FullHouse", "DoubleThreeKind", "ThreeKind", "FourKind", "FiveKind", "SixKind"}

func (t TenGrandCategory) String() string {
	return TenGrandCategoryArray[t]
}

func (t TenGrandCategory) EnumIndex() int {
	return int(t)
}

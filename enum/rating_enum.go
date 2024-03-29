package enum

// Rating enum info
// @Description database contains integer but values are "Gray", "Brown", "Green"
type Rating int

const (
	Gray Rating = iota
	BROWN
	GREEN
)

type RatingString string

const (
	RS0 RatingString = "Gray"
	RS1 RatingString = "Brown"
	RS2 RatingString = "Green"
)

var RatingArray = [3]string{"Gray", "Brown", "Green"}

func (r Rating) String() string {
	if r < 0 || r >= 3 {
		return ""
	}
	return RatingArray[r]
}

func (r Rating) EnumIndex() int {
	return int(r)
}

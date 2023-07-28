package enum

type Rating int

const (
	Gray Rating = iota
	BROWN
	GREEN
)

var RatingArray = [3]string{"Gray", "Brown", "Green"}

func (r Rating) String() string {
	return RatingArray[r]
}

func (r Rating) EnumIndex() int {
	return int(r)
}

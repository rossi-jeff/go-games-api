package enum

type Color int

const (
	Black Color = iota
	Blue
	Brown
	Green
	Orange
	Purple
	Red
	White
	Yellow
)

var ColorArray = [...]string{"Black", "Blue", "Brown", "Green", "Orange", "Purple", "Red", "White", "Yellow"}

func (c Color) String() string {
	if c < 0 || c >= 9 {
		return ""
	}
	return ColorArray[c]
}

func (c Color) EnumIndex() int {
	return int(c)
}
